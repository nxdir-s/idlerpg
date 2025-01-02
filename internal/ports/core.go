package ports

import (
	"context"
	"time"
)

type UsersPort interface{}

type UserServicePort interface{}

type UserTxServicePort interface {
	UserServicePort
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type JWTPort interface {
	IssueToken(ctx context.Context, userId int, email string, expires time.Time) (string, error)
	IssueRefreshToken(ctx context.Context, userId int, email string, expires time.Time) (string, error)

	ValidAccessToken(ctx context.Context, tokenStr string) error
	ValidRefreshToken(ctx context.Context, tokenStr string) error
}
