package secondary

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nxdir-s/idlerpg/internal/ports"
	"github.com/nxdir-s/idlerpg/internal/util"
	"go.opentelemetry.io/otel/trace"
)

type ErrConnect struct {
	err error
}

func (e *ErrConnect) Error() string {
	return "error creating connection pool: " + e.err.Error()
}

type ErrQueryRow struct {
	err error
}

func (e *ErrQueryRow) Error() string {
	return "error querying row: " + e.err.Error()
}

type ErrExecQuery struct {
	err error
}

func (e *ErrExecQuery) Error() string {
	return "error executing query: " + e.err.Error()
}

type ErrNotFound struct {
	name string
}

func (e *ErrNotFound) Error() string {
	return "error no rows affected, " + e.name + " not found"
}

type ErrCreateUser struct{}

func (e *ErrCreateUser) Error() string {
	return "failed to create new user"
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
	tracer trace.Tracer
}

// NewPostgresAdapter creates a new PostgresAdapter
func NewPostgresAdapter(pool PgxPool, logger *slog.Logger, tracer trace.Tracer, opts ...PostgresOpt) (*PostgresAdapter, error) {
	adapter := &PostgresAdapter{
		conn:   pool,
		logger: logger,
		tracer: tracer,
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

	txAdapter, err := NewPostgresAdapter(tx, a.logger, a.tracer, WithPgxTx(tx))
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

// CreateUser creates a new user and returns the user's id
func (a *PostgresAdapter) CreateUser(ctx context.Context, email string) (int, error) {
	ctx, span := a.tracer.Start(ctx, "create.user",
		trace.WithLinks(trace.LinkFromContext(ctx)),
	)
	defer span.End()

	args := pgx.NamedArgs{
		"email": email,
	}

	var userId int
	err := a.conn.QueryRow(ctx, CreateUserQuery, args).Scan(&userId)
	if err != nil {
		err = &ErrExecQuery{err}
		util.RecordError(span, "create.user failed", err)

		return 0, err
	}

	return userId, nil
}

const CreateUserQuery string = `
    INSERT INTO users (
        email
    ) VALUES (
        @email
    ) RETURNING id
`

// RemoveUser deletes a user using the supplied id
func (a *PostgresAdapter) RemoveUser(ctx context.Context, id int) error {
	ctx, span := a.tracer.Start(ctx, "remove.user",
		trace.WithLinks(trace.LinkFromContext(ctx)),
	)
	defer span.End()

	args := pgx.NamedArgs{
		"id": id,
	}

	resp, err := a.conn.Exec(ctx, RemoveUserQuery, args)
	if err != nil {
		err = &ErrExecQuery{err}
		util.RecordError(span, "remove.user failed", err)

		return err
	}

	if resp.RowsAffected() == 0 {
		return &ErrNotFound{"user"}
	}

	return nil
}

const RemoveUserQuery string = `
    DELETE FROM users
    WHERE id = @id
`

func (a *PostgresAdapter) GetUserID(ctx context.Context, email string) (int, error) {
	ctx, span := a.tracer.Start(ctx, "get.user.id",
		trace.WithLinks(trace.LinkFromContext(ctx)),
	)
	defer span.End()

	args := pgx.NamedArgs{
		"email": email,
	}

	var userId int
	err := a.conn.QueryRow(ctx, UserIdQuery, args).Scan(&userId)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		err = &ErrQueryRow{err}
		util.RecordError(span, "get.user.id failed", err)

		return 0, err
	}

	if userId == 0 || errors.Is(err, pgx.ErrNoRows) {
		return 0, nil
	}

	return userId, nil
}

const UserIdQuery string = `
    SELECT id
    FROM users
    WHERE email = @email
`

func (a *PostgresAdapter) UserExists(ctx context.Context, id int) (bool, error) {
	ctx, span := a.tracer.Start(ctx, "user.exists",
		trace.WithLinks(trace.LinkFromContext(ctx)),
	)
	defer span.End()

	args := pgx.NamedArgs{
		"id": id,
	}

	var userId int
	err := a.conn.QueryRow(ctx, UserExistsQuery, args).Scan(&userId)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		err = &ErrQueryRow{err}
		util.RecordError(span, "user.exists failed", err)

		return false, err
	}

	if userId == 0 || errors.Is(err, pgx.ErrNoRows) {
		return false, nil
	}

	return true, nil
}

const UserExistsQuery string = `
    SELECT id
    FROM users
    WHERE id = @id
`

func (a *PostgresAdapter) EmailExists(ctx context.Context, email string) (bool, error) {
	ctx, span := a.tracer.Start(ctx, "email.exists",
		trace.WithLinks(trace.LinkFromContext(ctx)),
	)
	defer span.End()

	args := pgx.NamedArgs{
		"email": email,
	}

	var userEmail string
	err := a.conn.QueryRow(ctx, EmailExistsQuery, args).Scan(&userEmail)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		err = &ErrQueryRow{err}
		util.RecordError(span, "email.exists failed", err)

		return false, err
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return false, nil
	}

	return true, nil
}

const EmailExistsQuery string = `
    SELECT email
    FROM users
    WHERE email = @email
`

func (a *PostgresAdapter) UpdateRefreshToken(ctx context.Context, id int, token string) error {
	ctx, span := a.tracer.Start(ctx, "update.refresh.token",
		trace.WithLinks(trace.LinkFromContext(ctx)),
	)
	defer span.End()

	args := pgx.NamedArgs{
		"id":    id,
		"token": token,
	}

	resp, err := a.conn.Exec(ctx, UpdateRefreshTokenQuery, args)
	if err != nil {
		err = &ErrExecQuery{err}
		util.RecordError(span, "update.refresh.token failed", err)

		return err
	}

	if resp.RowsAffected() == 0 {
		return &ErrNotFound{"user"}
	}

	return nil
}

const UpdateRefreshTokenQuery string = `
    UPDATE users
    SET refresh_token = @token
    WHERE id = @id
`
