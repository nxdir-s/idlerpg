package postgresqlcnpgio


// The template of the Pod to be created.
type PoolerSpecTemplate struct {
	// Standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *PoolerSpecTemplateMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// Specification of the desired behavior of the pod.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *PoolerSpecTemplateSpec `field:"optional" json:"spec" yaml:"spec"`
}

