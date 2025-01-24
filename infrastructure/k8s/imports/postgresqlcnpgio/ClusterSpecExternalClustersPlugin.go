package postgresqlcnpgio


// The configuration of the plugin that is taking care of WAL archiving and backups for this external cluster.
type ClusterSpecExternalClustersPlugin struct {
	// Name is the plugin name.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Enabled is true if this plugin will be used.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Parameters is the configuration of the plugin.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
}

