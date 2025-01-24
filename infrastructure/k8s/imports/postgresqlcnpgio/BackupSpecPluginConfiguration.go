package postgresqlcnpgio


// Configuration parameters passed to the plugin managing this backup.
type BackupSpecPluginConfiguration struct {
	// Name is the name of the plugin managing this backup.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Parameters are the configuration parameters passed to the backup plugin for this backup.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
}

