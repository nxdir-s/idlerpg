package postgresqlcnpgio


// ManagedService represents a specific service managed by the cluster.
//
// It includes the type of service and its associated template specification.
type ClusterSpecManagedServicesAdditional struct {
	// SelectorType specifies the type of selectors that the service will have.
	//
	// Valid values are "rw", "r", and "ro", representing read-write, read, and read-only services.
	SelectorType ClusterSpecManagedServicesAdditionalSelectorType `field:"required" json:"selectorType" yaml:"selectorType"`
	// ServiceTemplate is the template specification for the service.
	ServiceTemplate *ClusterSpecManagedServicesAdditionalServiceTemplate `field:"required" json:"serviceTemplate" yaml:"serviceTemplate"`
	// UpdateStrategy describes how the service differences should be reconciled.
	UpdateStrategy ClusterSpecManagedServicesAdditionalUpdateStrategy `field:"optional" json:"updateStrategy" yaml:"updateStrategy"`
}

