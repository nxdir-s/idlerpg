package postgresqlcnpgio


// Services roles managed by the `Cluster`.
type ClusterSpecManagedServices struct {
	// Additional is a list of additional managed services specified by the user.
	Additional *[]*ClusterSpecManagedServicesAdditional `field:"optional" json:"additional" yaml:"additional"`
	// DisabledDefaultServices is a list of service types that are disabled by default.
	//
	// Valid values are "r", and "ro", representing read, and read-only services.
	DisabledDefaultServices *[]ClusterSpecManagedServicesDisabledDefaultServices `field:"optional" json:"disabledDefaultServices" yaml:"disabledDefaultServices"`
}

