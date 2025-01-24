package postgresqlcnpgio


// The configuration that is used by the portions of PostgreSQL that are managed by the instance manager.
type ClusterSpecManaged struct {
	// Database roles managed by the `Cluster`.
	Roles *[]*ClusterSpecManagedRoles `field:"optional" json:"roles" yaml:"roles"`
	// Services roles managed by the `Cluster`.
	Services *ClusterSpecManagedServices `field:"optional" json:"services" yaml:"services"`
}

