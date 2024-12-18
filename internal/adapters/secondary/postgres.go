package secondary

import (
	"context"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nxdir-s/idlerpg/internal/ports"
)

type ErrConnect struct {
	err error
}

func (e *ErrConnect) Error() string {
	return "error creating connection pool: " + e.err.Error()
}

type PgxPool interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
}

// NewPgxPool creates a pgxpool.Pool
func NewPgxPool(ctx context.Context, dbUrl string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		return nil, &ErrConnect{err}
	}

	return pool, nil
}

type PgxTx interface {
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type PostgresOpt func(a *PostgresAdapter)

// WithPgxTx sets the transaction adapter
func WithPgxTx(tx PgxTx) PostgresOpt {
	return func(a *PostgresAdapter) {
		a.tx = tx
	}
}

type PostgresAdapter struct {
	conn   PgxPool
	tx     PgxTx
	logger *slog.Logger
}

// NewPostgresAdapter creates a new PostgresAdapter
func NewPostgresAdapter(ctx context.Context, pool PgxPool, logger *slog.Logger, opts ...PostgresOpt) (*PostgresAdapter, error) {
	adapter := &PostgresAdapter{
		conn:   pool,
		logger: logger,
	}

	for _, opt := range opts {
		opt(adapter)
	}

	return adapter, nil
}

// NewTransactionAdapter creates a postgres adapter for executing transactions
func (a *PostgresAdapter) NewTransactionAdapter(ctx context.Context) (ports.DatabaseTxPort, error) {
	tx, err := a.conn.Begin(ctx)
	if err != nil {
		return nil, err
	}

	txAdapter, err := NewPostgresAdapter(ctx, tx, a.logger, WithPgxTx(tx))
	if err != nil {
		return nil, err
	}

	return txAdapter, nil
}

// Commit commits the transaction after checking if the context has been canceled
func (a *PostgresAdapter) Commit(ctx context.Context) error {
	select {
	case <-ctx.Done():
		ctx, cancel := context.WithTimeout(ctx, time.Second*10)
		defer cancel()

		return a.tx.Rollback(ctx)
	default:
		return a.tx.Commit(ctx)
	}
}

// Rollback initiates a transaction rollback
func (a *PostgresAdapter) Rollback(ctx context.Context) error {
	return a.tx.Rollback(ctx)
}
