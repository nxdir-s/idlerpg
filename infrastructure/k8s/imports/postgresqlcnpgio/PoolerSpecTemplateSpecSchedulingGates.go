package postgresqlcnpgio


// PodSchedulingGate is associated to a Pod to guard its scheduling.
type PoolerSpecTemplateSpecSchedulingGates struct {
	// Name of the scheduling gate.
	//
	// Each scheduling gate must have a unique name field.
	Name *string `field:"required" json:"name" yaml:"name"`
}

