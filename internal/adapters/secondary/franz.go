package secondary

import (
	"context"
	"crypto/tls"
	"log/slog"

	"github.com/nxdir-s/idlerpg/protobuf"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sasl/scram"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/proto"
)

const (
	MaxPollFetches int = 1000
)

type FranzAdapterOpt func(a *FranzAdapter) error

func WithFranzConsumer(groupname string, brokers []string, username string, pass string) FranzAdapterOpt {
	return func(a *FranzAdapter) error {
		client, err := kgo.NewClient(
			kgo.SeedBrokers(brokers...),
			kgo.DialTLSConfig(&tls.Config{}),
			kgo.SASL(scram.Auth{User: username, Pass: pass}.AsSha256Mechanism()),
			kgo.ConsumerGroup(groupname),
			kgo.ConsumeTopics(a.topic),
			kgo.ConsumeResetOffset(kgo.NewOffset().AtStart()),
			kgo.DisableAutoCommit(),
			kgo.BlockRebalanceOnPoll(),
		)
		if err != nil {
			return err
		}

		a.consumer = client

		return nil
	}
}

func WithFranzProducer(brokers []string, username string, pass string) FranzAdapterOpt {
	return func(a *FranzAdapter) error {
		client, err := kgo.NewClient(
			kgo.SeedBrokers(brokers...),
			kgo.DialTLSConfig(&tls.Config{}),
			kgo.SASL(scram.Auth{User: username, Pass: pass}.AsSha256Mechanism()),
		)
		if err != nil {
			return err
		}

		a.producer = client

		return nil
	}
}

type FranzAdapter struct {
	producer *kgo.Client
	consumer *kgo.Client
	logger   *slog.Logger
	tracer   trace.Tracer
	topic    string
}

func NewFranzAdapter(topic string, logger *slog.Logger, tracer trace.Tracer, opts ...FranzAdapterOpt) (*FranzAdapter, error) {
	adapter := &FranzAdapter{
		logger: logger,
		topic:  topic,
	}

	for _, opt := range opts {
		if err := opt(adapter); err != nil {
			return nil, err
		}
	}

	return adapter, nil
}

func (a *FranzAdapter) SendUserEvent(ctx context.Context, event *protobuf.UserEvent) error {
	if a.producer == nil {
		return nil
	}

	data, err := proto.Marshal(event)
	if err != nil {
		return &ErrProtoMarshal{err}
	}

	a.producer.Produce(ctx, &kgo.Record{Topic: a.topic, Value: data}, nil)

	return nil
}

func (a *FranzAdapter) CloseProducer() error {
	if a.producer == nil {
		return nil
	}

	a.producer.Close()

	return nil
}

func (a *FranzAdapter) ConsumeUserEvents(ctx context.Context) {
	if a.consumer == nil {
		return
	}

	for {
		select {
		case <-ctx.Done():
			return
		default:
			fetches := a.consumer.PollRecords(ctx, MaxPollFetches)

			if errors := fetches.Errors(); len(errors) > 0 {
				for _, e := range errors {
					if e.Err == context.Canceled {
						a.logger.Info("received interrupt", slog.Any("err", e.Err))
						return
					}

					a.logger.Error("poll error", slog.Any("err", e))
				}
			}

			iter := fetches.RecordIter()
			for !iter.Done() {
				record := iter.Next()

				var msg protobuf.UserEvent
				if err := proto.Unmarshal(record.Value, &msg); err != nil {
					a.logger.Error("error decoding UserEvent", slog.Any("err", err))
					continue
				}

				a.logger.Info("consumed message",
					slog.String("action", msg.Action.String()),
					slog.Int("exp", int(msg.Exp)),
				)
			}

			if err := a.consumer.CommitUncommittedOffsets(ctx); err != nil {
				if err == context.Canceled {
					a.logger.Info("received interrupt", slog.Any("err", err))
					return
				}

				a.logger.Error("unable to commit offsets", slog.Any("err", err))
			}

			a.consumer.AllowRebalance()
		}
	}
}

func (a *FranzAdapter) CloseConsumer() error {
	if a.consumer == nil {
		return nil
	}

	a.consumer.Close()

	return nil
}
