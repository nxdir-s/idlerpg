package secondary

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/IBM/sarama"
	"github.com/nxdir-s/idlerpg/protobuf"
	"github.com/xdg-go/scram"
	"google.golang.org/protobuf/proto"
)

const (
	ProducerMaxRetry   int    = 10
	DefaultStrategy    string = "roundrobin"
	DefaultKeepRunning bool   = true

	PlayerEventsTopic string = "player-events"
)

type ErrSaramaCfg struct {
	err error
}

func (e *ErrSaramaCfg) Error() string {
	return "error creating sarama config: " + e.err.Error()
}

type ErrRPCfg struct {
	err error
}

func (e *ErrRPCfg) Error() string {
	return "error creating Redpanda config: " + e.err.Error()
}

type ErrUnknownStrategy struct {
	strategy string
}

func (e *ErrUnknownStrategy) Error() string {
	return "unrecognized consumer group partition assignor: " + e.strategy
}

type ErrConsumerStart struct {
	err error
}

func (e *ErrConsumerStart) Error() string {
	return "error starting sarama consumer group: " + e.err.Error()
}

type ErrProducerStart struct {
	err error
}

func (e *ErrProducerStart) Error() string {
	return "error starting sarama producer: " + e.err.Error()
}

type ErrCloseConsumer struct {
	err error
}

func (e *ErrCloseConsumer) Error() string {
	return "error closing consumer: " + e.err.Error()
}

type ErrCloseProducer struct {
	err error
}

func (e *ErrCloseProducer) Error() string {
	return "error closing producer: " + e.err.Error()
}

type ErrProtoMarshal struct {
	err error
}

func (e *ErrProtoMarshal) Error() string {
	return "failed to marshal protobuf: " + e.err.Error()
}

type ErrSendMesssage struct {
	err error
}

func (e *ErrSendMesssage) Error() string {
	return "failed to send kafka message: " + e.err.Error()
}

type ErrConsumeMessage struct {
	err error
}

func (e *ErrConsumeMessage) Error() string {
	return "failed to consume kafka message: " + e.err.Error()
}

type ErrNilProducer struct{}

func (e *ErrNilProducer) Error() string {
	return "nil producer in SaramaAdapter"
}

type ErrNilConsumer struct{}

func (e *ErrNilConsumer) Error() string {
	return "nil consumer in SaramaAdapter"
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
		return nil, &ErrUnknownStrategy{strategy}
	}

	return config, nil
}

// NewRPProducerCfg creates a config for RedPanda
func NewRPProducerCfg(userName string, pass string) (*sarama.Config, error) {
	config := sarama.NewConfig()

	version, err := sarama.ParseKafkaVersion("2.0.0")
	if err != nil {
		return nil, err
	}

	config.Version = version
	config.ClientID = "sasl_scram_client"
	config.Metadata.Full = true

	config.Producer.Retry.Max = ProducerMaxRetry
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	config.Net.SASL.Enable = true
	config.Net.SASL.User = userName
	config.Net.SASL.Password = pass
	config.Net.SASL.Handshake = true
	config.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient { return &XDGSCRAMClient{HashGeneratorFcn: sha256.New} }
	config.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA256

	return config, nil
}

type SaramaAdapterOpt func(a *SaramaAdapter) error

// WithSaramaConsumer adds a sarama.ConsumerGroup to the adapter
func WithSaramaConsumer() SaramaAdapterOpt {
	return func(a *SaramaAdapter) error {
		cfg, err := NewConsumerCfg(DefaultStrategy)
		if err != nil {
			return &ErrSaramaCfg{err}
		}

		consumer, err := sarama.NewConsumerGroup(a.brokers, "", cfg)
		if err != nil {
			return &ErrConsumerStart{err}
		}
		a.consumer = consumer

		return nil
	}
}

// WithSaramaProducer adds a sarama.SyncProducer to the adapter
func WithSaramaProducer() SaramaAdapterOpt {
	return func(a *SaramaAdapter) error {
		producer, err := sarama.NewSyncProducer(a.brokers, NewSyncProducerCfg())
		if err != nil {
			return &ErrProducerStart{err}
		}
		a.producer = producer

		return nil
	}
}

// WithRedPandaProducer adds a sarama.SyncProducer for RedPanda to the adapter
func WithRedPandaProducer(username string, pass string) SaramaAdapterOpt {
	return func(a *SaramaAdapter) error {
		cfg, err := NewRPProducerCfg(username, pass)
		if err != nil {
			return &ErrRPCfg{err}
		}

		producer, err := sarama.NewSyncProducer(a.brokers, cfg)
		if err != nil {
			return &ErrProducerStart{err}
		}
		a.producer = producer

		return nil
	}
}

type SaramaAdapter struct {
	producer sarama.SyncProducer
	consumer sarama.ConsumerGroup

	brokers []string

	sigUsr1 chan os.Signal
	sigTerm chan os.Signal

	keepRunning bool
	paused      bool
}

// NewSaramaAdapter creates a SaramaAdapter set up for producing and/or consuming kafka messages
func NewSaramaAdapter(ctx context.Context, brokers []string, opts ...SaramaAdapterOpt) (*SaramaAdapter, error) {
	sigusr1 := make(chan os.Signal, 1)
	signal.Notify(sigusr1, syscall.SIGUSR1)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	adapter := &SaramaAdapter{
		sigUsr1:     sigusr1,
		sigTerm:     sigterm,
		brokers:     brokers,
		keepRunning: DefaultKeepRunning,
	}

	for _, opt := range opts {
		if err := opt(adapter); err != nil {
			return nil, err
		}
	}

	return adapter, nil
}

// SendUserEvent marshals and sends player events to kafka
func (a *SaramaAdapter) SendUserEvent(ctx context.Context, event *protobuf.UserEvent) error {
	if a.producer == nil {
		return &ErrNilProducer{}
	}

	data, err := proto.Marshal(event)
	if err != nil {
		return &ErrProtoMarshal{err}
	}

	_, _, err = a.producer.SendMessage(&sarama.ProducerMessage{
		Topic: PlayerEventsTopic,
		Value: sarama.ByteEncoder(data),
	})
	if err != nil {
		return &ErrSendMesssage{err}
	}

	return nil
}

// CloseProducer closes the sarama producer
func (a *SaramaAdapter) CloseProducer() error {
	if a.producer == nil {
		return &ErrNilProducer{}
	}

	if err := a.producer.Close(); err != nil {
		return &ErrCloseProducer{err}
	}

	return nil
}

// Consume starts the sarama consumer
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

			if ctx.Err() != nil {
				return
			}

			handler.Reset()
		}
	}()

	handler.AwaitSetup()

	fmt.Fprint(os.Stdout, "sarama consumer up and running!\n")

	for a.keepRunning {
		select {
		case <-ctx.Done():
			fmt.Fprint(os.Stdout, "consumer terminating: context cancelled...\n")
			a.keepRunning = false
		case <-a.sigTerm:
			fmt.Fprint(os.Stdout, "consumer terminating via signal...\n")
			a.keepRunning = false
		case <-a.sigUsr1:
			a.toggleConsumption()
		}
	}

	cancel()
	wg.Wait()

	if consumeErr != nil {
		return &ErrConsumeMessage{consumeErr}
	}

	if err := a.consumer.Close(); err != nil {
		return &ErrCloseConsumer{err}
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

type XDGSCRAMClient struct {
	*scram.Client
	*scram.ClientConversation
	scram.HashGeneratorFcn
}

func (c *XDGSCRAMClient) Begin(userName, password, authzID string) error {
	client, err := c.HashGeneratorFcn.NewClient(userName, password, authzID)
	if err != nil {
		return err
	}

	c.Client = client
	c.ClientConversation = c.Client.NewConversation()

	return nil
}

func (c *XDGSCRAMClient) Step(challenge string) (string, error) {
	response, err := c.ClientConversation.Step(challenge)
	if err != nil {
		return "", err
	}

	return response, nil
}

func (c *XDGSCRAMClient) Done() bool {
	return c.ClientConversation.Done()
}
