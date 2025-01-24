package postgresqlcnpgio


// The cluster to backup.
type ScheduledBackupSpecCluster struct {
	// Name of the referent.
	Name *string `field:"required" json:"name" yaml:"name"`
}

