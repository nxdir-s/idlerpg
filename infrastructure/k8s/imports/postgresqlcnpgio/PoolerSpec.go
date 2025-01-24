package postgresqlcnpgio


// Specification of the desired behavior of the Pooler.
//
// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
type PoolerSpec struct {
	// This is the cluster reference on which the Pooler will work.
	//
	// Pooler name should never match with any cluster name within the same namespace.
	Cluster *PoolerSpecCluster `field:"required" json:"cluster" yaml:"cluster"`
	// The PgBouncer configuration.
	Pgbouncer *PoolerSpecPgbouncer `field:"required" json:"pgbouncer" yaml:"pgbouncer"`
	// The deployment strategy to use for pgbouncer to replace existing pods with new ones.
	DeploymentStrategy *PoolerSpecDeploymentStrategy `field:"optional" json:"deploymentStrategy" yaml:"deploymentStrategy"`
	// The number of replicas we want.
	//
	// Default: 1.
	Instances *float64 `field:"optional" json:"instances" yaml:"instances"`
	// The configuration of the monitoring infrastructure of this pooler.
	Monitoring *PoolerSpecMonitoring `field:"optional" json:"monitoring" yaml:"monitoring"`
	// Template for the Service to be created.
	ServiceTemplate *PoolerSpecServiceTemplate `field:"optional" json:"serviceTemplate" yaml:"serviceTemplate"`
	// The template of the Pod to be created.
	Template *PoolerSpecTemplate `field:"optional" json:"template" yaml:"template"`
	// Type of service to forward traffic to.
	//
	// Default: `rw`.
	Type PoolerSpecType `field:"optional" json:"type" yaml:"type"`
}

