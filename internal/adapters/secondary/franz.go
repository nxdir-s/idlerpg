package secondary

import (
	"context"
	"crypto/tls"
	"log/slog"

	"github.com/nxdir-s/idlerpg/internal/util"
	"github.com/nxdir-s/idlerpg/protobuf"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sasl/scram"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/proto"
)

const (
	MaxPollFetches int = 1000

	UserEventsTopic  string = "user.events"
	UserUpdatesTopic string = "user.updates"
)

type FranzAdapterOpt func(a *FranzAdapter) error

func WithFranzConsumer(topic string, groupname string, brokers []string, username string, pass string) FranzAdapterOpt {
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
		a.consumer = client
		a.groupName = groupname

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
	producer  *kgo.Client
	consumer  *kgo.Client
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

func (a *FranzAdapter) SendUserEvent(ctx context.Context, event *protobuf.UserEvent) error {
	if a.producer == nil {
		return nil
	}

	ctx, span := a.tracer.Start(ctx, "send "+UserEventsTopic,
		trace.WithSpanKind(trace.SpanKindProducer),
		trace.WithAttributes(
			attribute.String("messaging.system", "kafka"),
			attribute.String("messaging.destination.name", UserEventsTopic),
			attribute.String("messaging.operation.name", "send"),
			attribute.String("messaging.operation.type", "send"),
		),
	)
	defer span.End()

	data, err := proto.Marshal(event)
	if err != nil {
		err = &ErrProtoMarshal{err}
		util.RecordError(span, "error encoding UserEvent", err)

		return err
	}

	a.producer.Produce(ctx, &kgo.Record{Topic: UserEventsTopic, Value: data}, nil)

	return nil
}

func (a *FranzAdapter) SendUserUpdate(ctx context.Context, update *protobuf.UserUpdate) error {
	if a.producer == nil {
		return nil
	}

	ctx, span := a.tracer.Start(ctx, "send "+UserUpdatesTopic,
		trace.WithSpanKind(trace.SpanKindProducer),
		trace.WithAttributes(
			attribute.String("messaging.system", "kafka"),
			attribute.String("messaging.destination.name", UserUpdatesTopic),
			attribute.String("messaging.operation.name", "send"),
			attribute.String("messaging.operation.type", "send"),
		),
	)
	defer span.End()

	data, err := proto.Marshal(update)
	if err != nil {
		err = &ErrProtoMarshal{err}
		util.RecordError(span, "error encoding UserUpdate", err)

		return err
	}

	a.producer.Produce(ctx, &kgo.Record{Topic: UserUpdatesTopic, Value: data}, nil)

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
			ctx, span := a.tracer.Start(ctx, "receive "+UserEventsTopic,
				trace.WithSpanKind(trace.SpanKindClient),
				trace.WithAttributes(
					attribute.String("messaging.system", "kafka"),
					attribute.String("messaging.consumer.group.name", a.groupName),
					attribute.String("messaging.operation.name", "receive"),
					attribute.String("messaging.operation.type", "receive"),
				),
			)

			fetches := a.consumer.PollRecords(ctx, MaxPollFetches)
			span.End()

			ctx, span = a.tracer.Start(ctx, "process "+UserEventsTopic,
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

				var msg protobuf.UserEvent
				if err := proto.Unmarshal(record.Value, &msg); err != nil {
					a.logger.Error("error decoding UserEvent", slog.Any("err", err))

					util.RecordError(span, "error decoding UserEvent", err)
					span.End()

					continue
				}

				a.logger.Info("consumed message",
					slog.String("action", msg.Action.String()),
					slog.Int("exp", int(msg.Exp)),
				)
			}

			if err := a.consumer.CommitUncommittedOffsets(ctx); err != nil {
				if err == context.Canceled {
					a.logger.Error("received interrupt", slog.Any("err", err))

					util.RecordError(span, "received interrupt", err)
					span.End()

					return
				}

				a.logger.Error("unable to commit offsets", slog.Any("err", err))
			}

			a.consumer.AllowRebalance()
			span.End()
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
