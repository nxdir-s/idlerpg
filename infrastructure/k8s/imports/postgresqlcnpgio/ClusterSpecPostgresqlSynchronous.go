package postgresqlcnpgio


// Configuration of the PostgreSQL synchronous replication feature.
type ClusterSpecPostgresqlSynchronous struct {
	// Method to select synchronous replication standbys from the listed servers, accepting 'any' (quorum-based synchronous replication) or 'first' (priority-based synchronous replication) as values.
	Method ClusterSpecPostgresqlSynchronousMethod `field:"required" json:"method" yaml:"method"`
	// Specifies the number of synchronous standby servers that transactions must wait for responses from.
	Number *float64 `field:"required" json:"number" yaml:"number"`
	// If set to "required", data durability is strictly enforced.
	//
	// Write operations
	// with synchronous commit settings (`on`, `remote_write`, or `remote_apply`) will
	// block if there are insufficient healthy replicas, ensuring data persistence.
	// If set to "preferred", data durability is maintained when healthy replicas
	// are available, but the required number of instances will adjust dynamically
	// if replicas become unavailable. This setting relaxes strict durability enforcement
	// to allow for operational continuity. This setting is only applicable if both
	// `standbyNamesPre` and `standbyNamesPost` are unset (empty).
	DataDurability ClusterSpecPostgresqlSynchronousDataDurability `field:"optional" json:"dataDurability" yaml:"dataDurability"`
	// Specifies the maximum number of local cluster pods that can be automatically included in the `synchronous_standby_names` option in PostgreSQL.
	MaxStandbyNamesFromCluster *float64 `field:"optional" json:"maxStandbyNamesFromCluster" yaml:"maxStandbyNamesFromCluster"`
	// A user-defined list of application names to be added to `synchronous_standby_names` after local cluster pods (the order is only useful for priority-based synchronous replication).
	StandbyNamesPost *[]*string `field:"optional" json:"standbyNamesPost" yaml:"standbyNamesPost"`
	// A user-defined list of application names to be added to `synchronous_standby_names` before local cluster pods (the order is only useful for priority-based synchronous replication).
	StandbyNamesPre *[]*string `field:"optional" json:"standbyNamesPre" yaml:"standbyNamesPre"`
}

