package secondary

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/IBM/sarama"
	"github.com/nxdir-s/IdleRpg/internal/core/valobj"
)

const (
	ProducerMaxRetry   int    = 10
	DefaultStrategy    string = "roundrobin"
	DefaultKeepRunning bool   = true
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
	return "failed to start sarama consumer group: " + e.err.Error()
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
	return "failed to start sarama producer: " + e.err.Error()
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

func NewSyncProducerCfg() *sarama.Config {
	config := sarama.NewConfig()

	config.Version = sarama.DefaultVersion
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = ProducerMaxRetry
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	return config
}

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

type SaramaAdapter struct {
	producer sarama.SyncProducer
	consumer sarama.ConsumerGroup

	sigUsr1     chan os.Signal
	sigTerm     chan os.Signal
	keepRunning bool

	pauseConsumption bool
}

func NewSaramaAdapter(brokers []string) (*SaramaAdapter, error) {
	// On the broker side, you can change the following settings to get
	// stronger consistency guarantees:
	// - For your broker, set `unclean.leader.election.enable` to false
	// - For the topic, you could increase `min.insync.replicas`.

	producer, err := sarama.NewSyncProducer(brokers, NewSyncProducerCfg())
	if err != nil {
		return nil, &ProducerError{err}
	}

	cfg, err := NewConsumerCfg(DefaultStrategy)
	if err != nil {
		return nil, &ConfigError{err}
	}

	consumer, err := sarama.NewConsumerGroup(brokers, "", cfg)
	if err != nil {
		return nil, &ConsumerError{}
	}

	sigusr1 := make(chan os.Signal, 1)
	signal.Notify(sigusr1, syscall.SIGUSR1)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	return &SaramaAdapter{
		producer:    producer,
		consumer:    consumer,
		sigUsr1:     sigusr1,
		sigTerm:     sigterm,
		keepRunning: DefaultKeepRunning,
	}, nil
}

func (a *SaramaAdapter) Consume(ctx context.Context, handler ConsumeHandler, topics []string) error {
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
				fmt.Fprintf(os.Stderr, "error in consumer, exiting...\n")
				return
			}

			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}

			// handler.Ready = make(chan bool)
			handler.Reset()
		}
	}()

	// <-handler.Ready // wait til the consumer has been set up
	handler.AwaitSetup()

	fmt.Fprintf(os.Stdout, "Sarama consumer up and running!...\n")

	for a.keepRunning {
		select {
		case <-ctx.Done():
			fmt.Fprintf(os.Stdout, "terminating: context cancelled\n")
			a.keepRunning = false
		case <-a.sigTerm:
			fmt.Fprintf(os.Stdout, "terminating via signal\n")
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
	if a.pauseConsumption {
		a.consumer.ResumeAll()
		fmt.Fprintf(os.Stdout, "Resuming consumption...\n")
	}

	if !a.pauseConsumption {
		a.consumer.PauseAll()
		fmt.Fprintf(os.Stdout, "Pausing consumption...\n")
	}

	a.pauseConsumption = !a.pauseConsumption
}

// SendPlayerEvent marshals and sends player events to kafka
func (a *SaramaAdapter) SendPlayerEvent(ctx context.Context, event *valobj.PlayerEvent) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	_, _, err = a.producer.SendMessage(&sarama.ProducerMessage{
		Topic: "player-events",
		Value: sarama.StringEncoder(data),
	})
	if err != nil {
		return err
	}

	return nil
}

func (a *SaramaAdapter) CloseProducer() error {
	return a.producer.Close()
}
