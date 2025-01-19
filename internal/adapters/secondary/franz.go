package secondary

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/nxdir-s/idlerpg/protobuf"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sasl/scram"
	"google.golang.org/protobuf/proto"
)

type FranzAdapterOpt func(a *FranzAdapter) error

func WithFranzConsumer(brokers []string, username string, pass string) FranzAdapterOpt {
	return func(a *FranzAdapter) error {
		client, err := kgo.NewClient(
			kgo.SeedBrokers(brokers...),
			kgo.DialTLSConfig(&tls.Config{}),
			kgo.SASL(scram.Auth{User: username, Pass: pass}.AsSha256Mechanism()),
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
			kgo.ConsumerGroup(uuid.New().String()),
			kgo.ConsumeTopics(a.topic),
			kgo.ConsumeResetOffset(kgo.NewOffset().AtStart()),
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
}

func NewFranzAdapter(ctx context.Context, opts ...FranzAdapterOpt) (*FranzAdapter, error) {
	adapter := &FranzAdapter{
		topic: "player-events",
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

type Message struct {
	User    string `json:"user"`
	Message string `json:"message"`
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

				var msg Message
				if err := json.Unmarshal(record.Value, &msg); err != nil {
					fmt.Printf("error decoding message: %v\n", err)
					continue
				}

				fmt.Printf("%s: %s\n", msg.User, msg.Message)
			}
		}
	}
}
