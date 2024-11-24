package secondary

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/IBM/sarama"
	"github.com/nxdir-s/idlerpg/protobuf"
	"google.golang.org/protobuf/proto"
)

const (
	ProducerMaxRetry   int    = 10
	DefaultStrategy    string = "roundrobin"
	DefaultKeepRunning bool   = true

	PlayerEventsTopic string = "player-events"
)

type ConfigError struct {
	err error
}

func (e *ConfigError) Error() string {
	return "error creating sarama config: " + e.err.Error()
}

type StrategyError struct {
	strategy string
}

func (e *StrategyError) Error() string {
	return "unrecognized consumer group partition assignor: " + e.strategy
}

type ConsumerError struct {
	err error
}

func (e *ConsumerError) Error() string {
	return "error starting sarama consumer group: " + e.err.Error()
}

type CloseError struct {
	clientType string
	err        error
}

func (e *CloseError) Error() string {
	return "error closing " + e.clientType + " client: " + e.err.Error()
}

type ProducerError struct {
	err error
}

func (e *ProducerError) Error() string {
	return "error starting sarama producer: " + e.err.Error()
}

type ErrNilProducer struct{}

func (e *ErrNilProducer) Error() string {
	return "error, nil producer in SaramaAdapter"
}

type ErrNilConsumer struct{}

func (e *ErrNilConsumer) Error() string {
	return "error, nil consumer in SaramaAdapter"
}

type SaramaHandler struct {
	Ready chan bool
}

type ConsumeHandler interface {
	Setup(session sarama.ConsumerGroupSession) error
	Cleanup(session sarama.ConsumerGroupSession) error
	ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error
	AwaitSetup()
	Reset()
}

// NewSyncProducerCfg creates a producer config
func NewSyncProducerCfg() *sarama.Config {
	config := sarama.NewConfig()

	config.Version = sarama.DefaultVersion
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = ProducerMaxRetry
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	return config
}

// NewConsumerCfg creates a consumer config
func NewConsumerCfg(strategy string) (*sarama.Config, error) {
	config := sarama.NewConfig()

	config.Version = sarama.DefaultVersion
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	switch strategy {
	case "sticky":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategySticky()}
	case "roundrobin":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}
	case "range":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRange()}
	default:
		return nil, &StrategyError{}
	}

	return config, nil
}

type SaramaAdapterOpt func(a *SaramaAdapter) error

func WithConsumer() SaramaAdapterOpt {
	return func(a *SaramaAdapter) error {
		cfg, err := NewConsumerCfg(DefaultStrategy)
		if err != nil {
			return &ConfigError{err}
		}

		consumer, err := sarama.NewConsumerGroup(a.brokers, "", cfg)
		if err != nil {
			return &ConsumerError{err}
		}
		a.consumer = consumer

		return nil
	}
}

func WithProducer() SaramaAdapterOpt {
	return func(a *SaramaAdapter) error {
		producer, err := sarama.NewSyncProducer(a.brokers, NewSyncProducerCfg())
		if err != nil {
			return &ProducerError{err}
		}
		a.producer = producer

		return nil
	}
}

type SaramaAdapter struct {
	producer sarama.SyncProducer
	consumer sarama.ConsumerGroup

	brokers []string

	sigUsr1     chan os.Signal
	sigTerm     chan os.Signal
	keepRunning bool

	paused bool
}

// NewSaramaAdapter creates a SaramaAdapter set up for producing and consuming kafka messages
func NewSaramaAdapter(brokers []string, opts ...SaramaAdapterOpt) (*SaramaAdapter, error) {
	sigusr1 := make(chan os.Signal, 1)
	signal.Notify(sigusr1, syscall.SIGUSR1)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	adapter := &SaramaAdapter{
		sigUsr1:     sigusr1,
		sigTerm:     sigterm,
		keepRunning: DefaultKeepRunning,
	}

	for _, opt := range opts {
		if err := opt(adapter); err != nil {
			return nil, err
		}
	}

	return adapter, nil
}

// SendPlayerEvent marshals and sends player events to kafka
func (a *SaramaAdapter) SendPlayerEvent(ctx context.Context, event *protobuf.PlayerEvent) error {
	if a.producer == nil {
		return &ErrNilProducer{}
	}

	data, err := proto.Marshal(event)
	if err != nil {
		return err
	}

	_, _, err = a.producer.SendMessage(&sarama.ProducerMessage{
		Topic: PlayerEventsTopic,
		Value: sarama.ByteEncoder(data),
	})
	if err != nil {
		return err
	}

	return nil
}

func (a *SaramaAdapter) CloseProducer() error {
	if a.producer == nil {
		return &ErrNilProducer{}
	}

	return a.producer.Close()
}

func (a *SaramaAdapter) Consume(ctx context.Context, handler ConsumeHandler, topics []string) error {
	if a.consumer == nil {
		return &ErrNilConsumer{}
	}

	ctx, cancel := context.WithCancel(ctx)

	var wg sync.WaitGroup
	wg.Add(1)

	var consumeErr error
	go func() {
		defer wg.Done()
		for {
			if err := a.consumer.Consume(ctx, topics, handler); err != nil {
				if errors.Is(err, sarama.ErrClosedConsumerGroup) {
					return
				}

				consumeErr = err
				fmt.Fprintf(os.Stdout, "error in consumer, exiting...\n")

				return
			}

			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}

			handler.Reset()
		}
	}()

	handler.AwaitSetup()

	fmt.Fprint(os.Stdout, "sarama consumer up and running...\n")

	for a.keepRunning {
		select {
		case <-ctx.Done():
			fmt.Fprint(os.Stdout, "terminating: context cancelled...\n")
			a.keepRunning = false
		case <-a.sigTerm:
			fmt.Fprint(os.Stdout, "terminating via signal...\n")
			a.keepRunning = false
		case <-a.sigUsr1:
			a.toggleConsumption()
		}
	}

	cancel()
	wg.Wait()

	if consumeErr != nil {
		return consumeErr
	}

	if err := a.consumer.Close(); err != nil {
		return &CloseError{"consumer", err}
	}

	return nil
}

func (a *SaramaAdapter) toggleConsumption() {
	switch a.paused {
	case true:
		a.consumer.ResumeAll()
		fmt.Fprint(os.Stdout, "resuming consumption...\n")
	case false:
		a.consumer.PauseAll()
		fmt.Fprint(os.Stdout, "pausing consumption...\n")
	}

	a.paused = !a.paused
}
