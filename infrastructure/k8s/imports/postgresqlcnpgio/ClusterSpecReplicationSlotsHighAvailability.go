package postgresqlcnpgio


// Replication slots for high availability configuration.
type ClusterSpecReplicationSlotsHighAvailability struct {
	// If enabled (default), the operator will automatically manage replication slots on the primary instance and use them in streaming replication connections with all the standby instances that are part of the HA cluster.
	//
	// If disabled, the operator will not take advantage
	// of replication slots in streaming connections with the replicas.
	// This feature also controls replication slots in replica cluster,
	// from the designated primary to its cascading replicas.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Prefix for replication slots managed by the operator for HA.
	//
	// It may only contain lower case letters, numbers, and the underscore character.
	// This can only be set at creation time. By default set to `_cnpg_`.
	SlotPrefix *string `field:"optional" json:"slotPrefix" yaml:"slotPrefix"`
}

