package consumers

import (
	"context"
	"log/slog"

	"github.com/nxdir-s/idlerpg/internal/ports"
)

type UserEvents struct {
	kafka  ports.KafkaPort
	logger *slog.Logger
}

func NewUserEvents(kafka ports.KafkaPort, logger *slog.Logger) *UserEvents {
	return &UserEvents{
		kafka:  kafka,
		logger: logger,
	}
}

func (c *UserEvents) Start(ctx context.Context) {
	c.logger.Info("starting user.events consumer")
	c.kafka.ConsumeUserEvents(ctx)
}

func (c *UserEvents) Close() {
	c.logger.Info("closing user.events consumer")

	if err := c.kafka.CloseConsumer(); err != nil {
		c.logger.Error("error closing user.events consumer", slog.Any("err", err))
	}
}
