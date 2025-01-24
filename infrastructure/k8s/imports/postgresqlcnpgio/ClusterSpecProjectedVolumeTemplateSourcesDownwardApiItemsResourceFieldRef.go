package postgresqlcnpgio


// Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.
type ClusterSpecProjectedVolumeTemplateSourcesDownwardApiItemsResourceFieldRef struct {
	// Required: resource to select.
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// Container name: required for volumes, optional for env vars.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// Specifies the output format of the exposed resources, defaults to "1".
	Divisor ClusterSpecProjectedVolumeTemplateSourcesDownwardApiItemsResourceFieldRefDivisor `field:"optional" json:"divisor" yaml:"divisor"`
}

