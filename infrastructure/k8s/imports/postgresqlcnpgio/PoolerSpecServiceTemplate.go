package postgresqlcnpgio


// Template for the Service to be created.
type PoolerSpecServiceTemplate struct {
	// Standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *PoolerSpecServiceTemplateMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// Specification of the desired behavior of the service.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *PoolerSpecServiceTemplateSpec `field:"optional" json:"spec" yaml:"spec"`
}

