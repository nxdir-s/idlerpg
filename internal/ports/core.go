package ports

import "context"

type PlayersPort interface{}

type PlayerServicePort interface{}

type PlayerTxServicePort interface {
	PlayerServicePort
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}
