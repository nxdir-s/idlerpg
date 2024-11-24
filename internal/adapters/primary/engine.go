package primary

import "github.com/nxdir-s/idlerpg/internal/ports"

type EngineAdapter struct {
	players ports.PlayersPort
}

// NewEngineAdapter creates an EngineAdapter
func NewEngineAdapter(domain ports.PlayersPort) (*EngineAdapter, error) {
	return &EngineAdapter{
		players: domain,
	}, nil
}
