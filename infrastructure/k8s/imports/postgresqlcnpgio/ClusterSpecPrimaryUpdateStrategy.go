package postgresqlcnpgio


// Deployment strategy to follow to upgrade the primary server during a rolling update procedure, after all replicas have been successfully updated: it can be automated (`unsupervised` - default) or manual (`supervised`).
type ClusterSpecPrimaryUpdateStrategy string

const (
	// unsupervised.
	ClusterSpecPrimaryUpdateStrategy_UNSUPERVISED ClusterSpecPrimaryUpdateStrategy = "UNSUPERVISED"
	// supervised.
	ClusterSpecPrimaryUpdateStrategy_SUPERVISED ClusterSpecPrimaryUpdateStrategy = "SUPERVISED"
)

