package domain

import (
	"context"

	"github.com/nxdir-s/idlerpg/internal/core/valobj"
	"github.com/nxdir-s/idlerpg/internal/ports"
)

type ErrUserExists struct {
	email string
}

func (e *ErrUserExists) Error() string {
	return "user already exists for " + e.email
}

type Users struct {
	service ports.UserServicePort
	jwt     ports.JWTPort
}

func NewUsers(ctx context.Context, service ports.UserServicePort, jwt ports.JWTPort) (*Users, error) {
	return &Users{
		service: service,
		jwt:     jwt,
	}, nil
}

func (d *Users) Login(ctx context.Context, email string) (*valobj.Token, error) {
	id, err := d.service.GetUserID(ctx, email)
	if err != nil {
		return nil, err
	}

	return d.jwt.IssueToken(ctx, id, email)
}

func (d *Users) Register(ctx context.Context, email string) (*valobj.Token, error) {
	emailExists, err := d.service.EmailExists(ctx, email)
	if err != nil {
		return nil, err
	}

	if emailExists {
		return nil, &ErrUserExists{email}
	}

	id, err := d.service.CreateUser(ctx, email)
	if err != nil {
		return nil, err
	}

	return d.jwt.IssueToken(ctx, id, email)
}
