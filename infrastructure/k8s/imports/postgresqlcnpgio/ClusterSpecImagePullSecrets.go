package postgresqlcnpgio


// LocalObjectReference contains enough information to let you locate a local object with a known type inside the same namespace.
type ClusterSpecImagePullSecrets struct {
	// Name of the referent.
	Name *string `field:"required" json:"name" yaml:"name"`
}

