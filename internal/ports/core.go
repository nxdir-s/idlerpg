package ports

import "context"

type UsersPort interface{}

type UserServicePort interface{}

type UserTxServicePort interface {
	UserServicePort
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type PlayersPort interface{}

type PlayerServicePort interface{}

type PlayerTxServicePort interface {
	PlayerServicePort
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}
