package postgresqlcnpgio


// Method to select synchronous replication standbys from the listed servers, accepting 'any' (quorum-based synchronous replication) or 'first' (priority-based synchronous replication) as values.
type ClusterSpecPostgresqlSynchronousMethod string

const (
	// any.
	ClusterSpecPostgresqlSynchronousMethod_ANY ClusterSpecPostgresqlSynchronousMethod = "ANY"
	// first.
	ClusterSpecPostgresqlSynchronousMethod_FIRST ClusterSpecPostgresqlSynchronousMethod = "FIRST"
)

