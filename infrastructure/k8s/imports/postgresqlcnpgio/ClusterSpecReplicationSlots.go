package postgresqlcnpgio


// Replication slots management configuration.
type ClusterSpecReplicationSlots struct {
	// Replication slots for high availability configuration.
	HighAvailability *ClusterSpecReplicationSlotsHighAvailability `field:"optional" json:"highAvailability" yaml:"highAvailability"`
	// Configures the synchronization of the user defined physical replication slots.
	SynchronizeReplicas *ClusterSpecReplicationSlotsSynchronizeReplicas `field:"optional" json:"synchronizeReplicas" yaml:"synchronizeReplicas"`
	// Standby will update the status of the local replication slots every `updateInterval` seconds (default 30).
	UpdateInterval *float64 `field:"optional" json:"updateInterval" yaml:"updateInterval"`
}

