package postgresqlcnpgio


// Type of service to forward traffic to.
//
// Default: `rw`.
type PoolerSpecType string

const (
	// rw.
	PoolerSpecType_RW PoolerSpecType = "RW"
	// ro.
	PoolerSpecType_RO PoolerSpecType = "RO"
)

