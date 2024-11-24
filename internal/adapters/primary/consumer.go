package primary

import (
	"github.com/nxdir-s/idlerpg/internal/ports"
)

type ConsumerAdapter struct {
	players ports.PlayersPort
}

// NewConsumerAdapter creates a ConsumerAdapter
func NewConsumerAdapter(domain ports.PlayersPort) (*ConsumerAdapter, error) {
	return &ConsumerAdapter{
		players: domain,
	}, nil
}
