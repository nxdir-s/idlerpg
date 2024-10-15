package secondary

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ConnectionError struct {
	err error
}

func (e *ConnectionError) Error() string {
	return "unable to create connection pool: " + e.err.Error()
}

type PgxPool interface {
	Ping(ctx context.Context) error
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
}

func NewPgxPool(ctx context.Context, dbUrl string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		return nil, &ConnectionError{err}
	}

	return pool, nil
}

type PostgresAdapter struct {
	conn   PgxPool
	logger *slog.Logger
}

func NewPostgresAdapter(ctx context.Context, pool PgxPool, logger *slog.Logger) (*PostgresAdapter, error) {
	return &PostgresAdapter{
		conn:   pool,
		logger: logger,
	}, nil
}

func (a *PostgresAdapter) Ping(ctx context.Context) error {
	return a.conn.Ping(ctx)
}
