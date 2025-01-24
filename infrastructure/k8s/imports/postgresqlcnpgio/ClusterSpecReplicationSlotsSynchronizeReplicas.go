package postgresqlcnpgio


// Configures the synchronization of the user defined physical replication slots.
type ClusterSpecReplicationSlotsSynchronizeReplicas struct {
	// When set to true, every replication slot that is on the primary is synchronized on each standby.
	Enabled *bool `field:"required" json:"enabled" yaml:"enabled"`
	// List of regular expression patterns to match the names of replication slots to be excluded (by default empty).
	ExcludePatterns *[]*string `field:"optional" json:"excludePatterns" yaml:"excludePatterns"`
}

