package secondary

import (
	"context"
	"crypto/tls"
	"log/slog"

	"github.com/nxdir-s/idlerpg/internal/adapters/secondary/franz"
	"github.com/nxdir-s/idlerpg/internal/util"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sasl/scram"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	MaxPollFetches int = 1000

	UserEventsTopic  string = "user.events"
	UserUpdatesTopic string = "user.updates"
)

type FranzAdapterOpt func(a *FranzAdapter) error

func WithConsumer(topic string, groupname string, brokers []string, username string, pass string) FranzAdapterOpt {
	return func(a *FranzAdapter) error {
		client, err := kgo.NewClient(
			kgo.SeedBrokers(brokers...),
			kgo.DialTLSConfig(&tls.Config{}),
			kgo.SASL(scram.Auth{User: username, Pass: pass}.AsSha256Mechanism()),
			kgo.ConsumerGroup(groupname),
			kgo.ConsumeTopics(topic),
			kgo.ConsumeResetOffset(kgo.NewOffset().AtStart()),
			kgo.DisableAutoCommit(),
			kgo.BlockRebalanceOnPoll(),
		)
		if err != nil {
			return err
		}

		a.topic = topic
		a.client = client
		a.groupName = groupname

		return nil
	}
}

func WithProducer(brokers []string, username string, pass string) FranzAdapterOpt {
	return func(a *FranzAdapter) error {
		client, err := kgo.NewClient(
			kgo.SeedBrokers(brokers...),
			kgo.DialTLSConfig(&tls.Config{}),
			kgo.SASL(scram.Auth{User: username, Pass: pass}.AsSha256Mechanism()),
		)
		if err != nil {
			return err
		}

		a.client = client

		return nil
	}
}

type FranzAdapter struct {
	client    *kgo.Client
	logger    *slog.Logger
	tracer    trace.Tracer
	topic     string
	groupName string
}

func NewFranzAdapter(logger *slog.Logger, tracer trace.Tracer, opts ...FranzAdapterOpt) (*FranzAdapter, error) {
	adapter := &FranzAdapter{
		logger: logger,
		tracer: tracer,
	}

	for _, opt := range opts {
		if err := opt(adapter); err != nil {
			return nil, err
		}
	}

	return adapter, nil
}

func (a *FranzAdapter) Send(ctx context.Context, record protoreflect.ProtoMessage) error {
	if a.client == nil {
		return nil
	}

	ctx, span := a.tracer.Start(ctx, "send "+a.topic,
		trace.WithSpanKind(trace.SpanKindProducer),
		trace.WithAttributes(
			attribute.String("messaging.system", "kafka"),
			attribute.String("messaging.destination.name", a.topic),
			attribute.String("messaging.operation.name", "send"),
			attribute.String("messaging.operation.type", "send"),
		),
	)
	defer span.End()

	data, err := proto.Marshal(record)
	if err != nil {
		err = &ErrProtoMarshal{err}
		util.RecordError(span, "error encoding "+a.topic+" record", err)

		return err
	}

	a.client.Produce(ctx, &kgo.Record{Topic: a.topic, Value: data}, nil)

	return nil
}

func (a *FranzAdapter) Consume(ctx context.Context, consumer franz.Consumer) {
	if a.client == nil {
		return
	}

	for {
		select {
		case <-ctx.Done():
			return
		default:
			ctx, span := a.tracer.Start(ctx, "receive "+a.topic,
				trace.WithSpanKind(trace.SpanKindClient),
				trace.WithAttributes(
					attribute.String("messaging.system", "kafka"),
					attribute.String("messaging.consumer.group.name", a.groupName),
					attribute.String("messaging.operation.name", "receive"),
					attribute.String("messaging.operation.type", "receive"),
				),
			)

			fetches := a.client.PollRecords(ctx, MaxPollFetches)
			span.End()

			ctx, span = a.tracer.Start(ctx, "process "+a.topic,
				trace.WithSpanKind(trace.SpanKindConsumer),
				trace.WithAttributes(
					attribute.String("messaging.system", "kafka"),
					attribute.String("messaging.consumer.group.name", a.groupName),
					attribute.String("messaging.operation.name", "process"),
					attribute.String("messaging.operation.type", "process"),
				),
			)

			if errors := fetches.Errors(); len(errors) > 0 {
				for _, e := range errors {
					if e.Err == context.Canceled {
						a.logger.Error("received interrupt", slog.Any("err", e.Err))

						util.RecordError(span, "received interrupt", e.Err)
						span.End()

						return
					}

					a.logger.Error("poll error", slog.Any("err", e))
				}
			}

			iter := fetches.RecordIter()
			for !iter.Done() {
				record := iter.Next()

				if err := consumer.Process(ctx, record); err != nil {
					a.logger.Error("error processing "+a.topic+" record", slog.Any("err", err))

					util.RecordError(span, "error processing "+a.topic+" record", err)
					span.End()

					return // return or continue?
				}

				a.logger.Info("consumed record", slog.String("topic", a.topic))
			}

			if err := a.client.CommitUncommittedOffsets(ctx); err != nil {
				if err == context.Canceled {
					a.logger.Error("received interrupt", slog.Any("err", err))

					util.RecordError(span, "received interrupt", err)
					span.End()

					return
				}

				a.logger.Error("unable to commit offsets", slog.Any("err", err))
			}

			a.client.AllowRebalance()
			span.End()
		}
	}
}

func (a *FranzAdapter) Close() error {
	if a.client == nil {
		return nil
	}

	a.client.Close()

	return nil
}
