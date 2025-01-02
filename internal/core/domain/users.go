package domain

import (
	"context"

	"github.com/nxdir-s/idlerpg/internal/ports"
)

type Users struct {
	service ports.UserServicePort
}

func NewUsers(ctx context.Context, service ports.UserServicePort) (*Users, error) {
	return &Users{
		service: service,
	}, nil
}
