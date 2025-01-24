package postgresqlcnpgio


// This is the cluster reference on which the Pooler will work.
//
// Pooler name should never match with any cluster name within the same namespace.
type PoolerSpecCluster struct {
	// Name of the referent.
	Name *string `field:"required" json:"name" yaml:"name"`
}

