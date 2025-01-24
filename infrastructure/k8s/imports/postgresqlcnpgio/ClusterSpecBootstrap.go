package postgresqlcnpgio


// Instructions to bootstrap this cluster.
type ClusterSpecBootstrap struct {
	// Bootstrap the cluster via initdb.
	Initdb *ClusterSpecBootstrapInitdb `field:"optional" json:"initdb" yaml:"initdb"`
	// Bootstrap the cluster taking a physical backup of another compatible PostgreSQL instance.
	PgBasebackup *ClusterSpecBootstrapPgBasebackup `field:"optional" json:"pgBasebackup" yaml:"pgBasebackup"`
	// Bootstrap the cluster from a backup.
	Recovery *ClusterSpecBootstrapRecovery `field:"optional" json:"recovery" yaml:"recovery"`
}

