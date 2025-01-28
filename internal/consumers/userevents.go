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
	kafka    ports.KafkaPort
	database ports.DatabasePort
	logger   *slog.Logger
}

func NewUserEvents(kafka ports.KafkaPort, db ports.DatabasePort, logger *slog.Logger) *UserEvents {
	return &UserEvents{
		kafka:    kafka,
		database: db,
		logger:   logger,
	}
}

func (c *UserEvents) Start(ctx context.Context) {
	c.logger.Info("starting user.events consumer")

	c.kafka.Consume(ctx, c)
}

func (c *UserEvents) Process(ctx context.Context, record *kgo.Record) error {
	var msg protobuf.UserEvent
	if err := proto.Unmarshal(record.Value, &msg); err != nil {
		return err
	}

	return nil
}

func (c *UserEvents) Close() {
	c.logger.Info("closing user.events consumer")

	if err := c.kafka.Close(); err != nil {
		c.logger.Error("error closing user.events consumer", slog.Any("err", err))
	}
}
