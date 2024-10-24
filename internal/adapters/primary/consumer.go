package primary

import (
	"context"

	"github.com/nxdir-s/IdleRpg/internal/ports"
)

type ConsumerAdapter struct {
	players ports.PlayersPort
}

func NewConsumerAdapter(ctx context.Context, players ports.PlayersPort) (*ConsumerAdapter, error) {
	return &ConsumerAdapter{
		players: players,
	}, nil
}
