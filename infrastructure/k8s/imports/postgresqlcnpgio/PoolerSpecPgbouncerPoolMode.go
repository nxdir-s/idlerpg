package postgresqlcnpgio


// The pool mode.
//
// Default: `session`.
type PoolerSpecPgbouncerPoolMode string

const (
	// session.
	PoolerSpecPgbouncerPoolMode_SESSION PoolerSpecPgbouncerPoolMode = "SESSION"
	// transaction.
	PoolerSpecPgbouncerPoolMode_TRANSACTION PoolerSpecPgbouncerPoolMode = "TRANSACTION"
)

