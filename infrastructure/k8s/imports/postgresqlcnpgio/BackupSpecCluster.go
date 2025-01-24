package postgresqlcnpgio


// The cluster to backup.
type BackupSpecCluster struct {
	// Name of the referent.
	Name *string `field:"required" json:"name" yaml:"name"`
}

