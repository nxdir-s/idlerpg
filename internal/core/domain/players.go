package domain

import "github.com/nxdir-s/idlerpg/internal/ports"

type Players struct {
	service ports.PlayerServicePort
}

func NewPlayers(service ports.PlayerServicePort) (*Players, error) {
	return &Players{
		service: service,
	}, nil
}
