package postgresqlcnpgio


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
type ClusterSpecPostgresqlSynchronousDataDurability string

const (
	// required.
	ClusterSpecPostgresqlSynchronousDataDurability_REQUIRED ClusterSpecPostgresqlSynchronousDataDurability = "REQUIRED"
	// preferred.
	ClusterSpecPostgresqlSynchronousDataDurability_PREFERRED ClusterSpecPostgresqlSynchronousDataDurability = "PREFERRED"
)

