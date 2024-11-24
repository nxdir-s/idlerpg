package primary

import (
	"context"

	"github.com/nxdir-s/idlerpg/internal/ports"
)

type ConsumerAdapter struct {
	players ports.PlayersPort
}

// NewConsumerAdapter creates a ConsumerAdapter
func NewConsumerAdapter(ctx context.Context, domain ports.PlayersPort) (*ConsumerAdapter, error) {
	return &ConsumerAdapter{
		players: domain,
	}, nil
}
