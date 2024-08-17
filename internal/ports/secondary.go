package ports

import (
	"context"

	"github.com/nxdir-s/IdleRpg/internal/core/valobj"
)

type KafkaPort interface {
	SendPlayerEvent(ctx context.Context, event *valobj.PlayerEvent) error
	Close() error
}
