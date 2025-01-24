package postgresqlcnpgio


// Bootstrap the cluster taking a physical backup of another compatible PostgreSQL instance.
type ClusterSpecBootstrapPgBasebackup struct {
	// The name of the server of which we need to take a physical backup.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Name of the database used by the application.
	//
	// Default: `app`.
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Name of the owner of the database in the instance to be used by applications.
	//
	// Defaults to the value of the `database` key.
	// Default: the value of the `database` key.
	//
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// Name of the secret containing the initial credentials for the owner of the user database.
	//
	// If empty a new secret will be
	// created from scratch.
	Secret *ClusterSpecBootstrapPgBasebackupSecret `field:"optional" json:"secret" yaml:"secret"`
}

