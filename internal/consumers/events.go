package consumers

import (
	"context"
	"log/slog"

	"github.com/nxdir-s/idlerpg/internal/ports"
)

type EventsConsumer struct {
	kafka  ports.KafkaPort
	logger *slog.Logger
}

func NewEventsConsumer(kafka ports.KafkaPort, logger *slog.Logger) *EventsConsumer {
	return &EventsConsumer{
		kafka:  kafka,
		logger: logger,
	}
}

func (c *EventsConsumer) Start(ctx context.Context) {
	c.logger.Info("starting events consumer")
	c.kafka.ConsumeUserEvent(ctx)
}

func (c *EventsConsumer) Close() {
	c.logger.Info("closing events consumer")

	if err := c.kafka.CloseConsumer(); err != nil {
		c.logger.Error("error closing events consumer", slog.Any("err", err))
	}
}
