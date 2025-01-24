package postgresqlcnpgio


// Replica cluster configuration.
type ClusterSpecReplica struct {
	// The name of the external cluster which is the replication origin.
	Source *string `field:"required" json:"source" yaml:"source"`
	// If replica mode is enabled, this cluster will be a replica of an existing cluster.
	//
	// Replica cluster can be created from a recovery
	// object store or via streaming through pg_basebackup.
	// Refer to the Replica clusters page of the documentation for more information.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// When replica mode is enabled, this parameter allows you to replay transactions only when the system time is at least the configured time past the commit time.
	//
	// This provides an opportunity to correct
	// data loss errors. Note that when this parameter is set, a promotion
	// token cannot be used.
	MinApplyDelay *string `field:"optional" json:"minApplyDelay" yaml:"minApplyDelay"`
	// Primary defines which Cluster is defined to be the primary in the distributed PostgreSQL cluster, based on the topology specified in externalClusters.
	Primary *string `field:"optional" json:"primary" yaml:"primary"`
	// A demotion token generated by an external cluster used to check if the promotion requirements are met.
	PromotionToken *string `field:"optional" json:"promotionToken" yaml:"promotionToken"`
	// Self defines the name of this cluster.
	//
	// It is used to determine if this is a primary
	// or a replica cluster, comparing it with `primary`.
	Self *string `field:"optional" json:"self" yaml:"self"`
}

