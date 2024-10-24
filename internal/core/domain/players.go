package domain

import "github.com/nxdir-s/IdleRpg/internal/ports"

type Players struct {
	service ports.PlayerServicePort
}

func NewPlayers(service ports.PlayerServicePort) (*Players, error) {
	return &Players{
		service: service,
	}, nil
}
