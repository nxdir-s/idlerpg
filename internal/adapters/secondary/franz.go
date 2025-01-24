package secondary

import (
	"context"
	"crypto/tls"
	"fmt"
	"log/slog"

	"github.com/nxdir-s/idlerpg/protobuf"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sasl/scram"
	"google.golang.org/protobuf/proto"
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
	topic    string
	logger   *slog.Logger
}

func NewFranzAdapter(ctx context.Context, logger *slog.Logger, opts ...FranzAdapterOpt) (*FranzAdapter, error) {
	adapter := &FranzAdapter{
		topic:  "player-events",
		logger: logger,
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

func (a *FranzAdapter) ConsumeUserEvent(ctx context.Context) {
	if a.consumer == nil {
		return
	}

	for {
		select {
		case <-ctx.Done():
			return
		default:
			fetches := a.consumer.PollFetches(ctx)
			iter := fetches.RecordIter()

			for !iter.Done() {
				record := iter.Next()

				var msg protobuf.UserEvent
				if err := proto.Unmarshal(record.Value, &msg); err != nil {
					fmt.Printf("error decoding message: %v\n", err)
					continue
				}

				a.logger.Info("consumed message",
					slog.String("action", msg.Action.String()),
					slog.Int("exp", int(msg.Exp)),
				)
			}
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