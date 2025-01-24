package postgresqlcnpgio


// PluginConfiguration specifies a plugin that need to be loaded for this cluster to be reconciled.
type ClusterSpecPlugins struct {
	// Name is the plugin name.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Enabled is true if this plugin will be used.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Parameters is the configuration of the plugin.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
}

