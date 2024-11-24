package primary

import (
	"context"

	"github.com/nxdir-s/idlerpg/internal/ports"
)

type EngineAdapter struct {
	players ports.PlayersPort
}

// NewEngineAdapter creates an EngineAdapter
func NewEngineAdapter(ctx context.Context, domain ports.PlayersPort) (*EngineAdapter, error) {
	return &EngineAdapter{
		players: domain,
	}, nil
}
