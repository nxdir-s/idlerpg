package postgresqlcnpgio


// SecretKeySelector contains enough information to let you locate the key of a Secret.
type ClusterSpecMonitoringCustomQueriesSecret struct {
	// The key to select.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Name of the referent.
	Name *string `field:"required" json:"name" yaml:"name"`
}

