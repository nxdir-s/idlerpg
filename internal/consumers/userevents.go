package consumers

import (
	"context"
	"log/slog"

	"github.com/nxdir-s/idlerpg/internal/ports"
	"github.com/nxdir-s/idlerpg/protobuf"
	"github.com/twmb/franz-go/pkg/kgo"
	"google.golang.org/protobuf/proto"
)

type UserEvents struct {
	kafka   ports.KafkaPort
	adapter ports.ConsumerPort
	logger  *slog.Logger
}

func NewUserEvents(kafka ports.KafkaPort, adapter ports.ConsumerPort, logger *slog.Logger) *UserEvents {
	return &UserEvents{
		kafka:   kafka,
		adapter: adapter,
		logger:  logger,
	}
}

func (c *UserEvents) Start(ctx context.Context) {
	c.logger.Info("starting user.events consumer")

	c.kafka.Consume(ctx, c)
}

func (c *UserEvents) Process(ctx context.Context, record *kgo.Record) error {
	var event protobuf.UserEvent
	if err := proto.Unmarshal(record.Value, &event); err != nil {
		return err
	}

	return c.adapter.ProcessUserEvent(ctx, &event)
}

func (c *UserEvents) Close() {
	c.logger.Info("closing user.events consumer")

	if err := c.kafka.Close(); err != nil {
		c.logger.Error("error closing user.events consumer", slog.Any("err", err))
	}
}
