package postgresqlcnpgio


// Define a maintenance window for the Kubernetes nodes.
type ClusterSpecNodeMaintenanceWindow struct {
	// Is there a node maintenance activity in progress?
	InProgress *bool `field:"optional" json:"inProgress" yaml:"inProgress"`
	// Reuse the existing PVC (wait for the node to come up again) or not (recreate it elsewhere - when `instances` >1).
	ReusePvc *bool `field:"optional" json:"reusePvc" yaml:"reusePvc"`
}

