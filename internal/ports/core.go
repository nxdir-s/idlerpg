package ports

import (
	"context"
	"time"

	"github.com/nxdir-s/idlerpg/internal/core/valobj"
)

type UsersPort interface{}

type UserServicePort interface {
	GetUserID(ctx context.Context, email string) (int, error)
	EmailExists(ctx context.Context, email string) (bool, error)
	CreateUser(ctx context.Context, email string) (int, error)
}

type UserTxServicePort interface {
	UserServicePort
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type JWTPort interface {
	IssueToken(ctx context.Context, id int, email string) (*valobj.Token, error)
	IssueAccessToken(ctx context.Context, userId int, email string, expires time.Time) (string, error)
	IssueRefreshToken(ctx context.Context, userId int, email string, expires time.Time) (string, error)

	ValidAccessToken(ctx context.Context, tokenStr string) error
	ValidRefreshToken(ctx context.Context, tokenStr string) error
}
