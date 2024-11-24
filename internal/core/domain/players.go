package domain

import (
	"context"

	"github.com/nxdir-s/idlerpg/internal/ports"
)

type Players struct {
	service ports.PlayerServicePort
}

func NewPlayers(ctx context.Context, service ports.PlayerServicePort) (*Players, error) {
	return &Players{
		service: service,
	}, nil
}
