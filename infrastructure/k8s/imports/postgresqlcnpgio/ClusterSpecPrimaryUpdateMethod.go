package postgresqlcnpgio


// Method to follow to upgrade the primary server during a rolling update procedure, after all replicas have been successfully updated: it can be with a switchover (`switchover`) or in-place (`restart` - default).
type ClusterSpecPrimaryUpdateMethod string

const (
	// switchover.
	ClusterSpecPrimaryUpdateMethod_SWITCHOVER ClusterSpecPrimaryUpdateMethod = "SWITCHOVER"
	// restart.
	ClusterSpecPrimaryUpdateMethod_RESTART ClusterSpecPrimaryUpdateMethod = "RESTART"
)

