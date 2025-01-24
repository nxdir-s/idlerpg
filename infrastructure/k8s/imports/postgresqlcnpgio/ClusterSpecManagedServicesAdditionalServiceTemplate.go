package postgresqlcnpgio


// ServiceTemplate is the template specification for the service.
type ClusterSpecManagedServicesAdditionalServiceTemplate struct {
	// Standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *ClusterSpecManagedServicesAdditionalServiceTemplateMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// Specification of the desired behavior of the service.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *ClusterSpecManagedServicesAdditionalServiceTemplateSpec `field:"optional" json:"spec" yaml:"spec"`
}

