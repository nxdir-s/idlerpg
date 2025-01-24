package postgresqlcnpgio


// The deployment strategy to use for pgbouncer to replace existing pods with new ones.
type PoolerSpecDeploymentStrategy struct {
	// Rolling update config params.
	//
	// Present only if DeploymentStrategyType =
	// RollingUpdate.
	RollingUpdate *PoolerSpecDeploymentStrategyRollingUpdate `field:"optional" json:"rollingUpdate" yaml:"rollingUpdate"`
	// Type of deployment.
	//
	// Can be "Recreate" or "RollingUpdate". Default is RollingUpdate.
	// Default: RollingUpdate.
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

