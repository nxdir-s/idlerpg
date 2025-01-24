package postgresqlcnpgio


// Resources is the total amount of CPU and Memory resources required by all containers in the pod.
//
// It supports specifying Requests and Limits for
// "cpu" and "memory" resource names only. ResourceClaims are not supported.
//
// This field enables fine-grained control over resource allocation for the
// entire pod, allowing resource sharing among containers in a pod.
//
// This is an alpha field and requires enabling the PodLevelResources feature
// gate.
type PoolerSpecTemplateSpecResources struct {
	// Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.
	//
	// This is an alpha field and requires enabling the
	// DynamicResourceAllocation feature gate.
	//
	// This field is immutable. It can only be set for containers.
	Claims *[]*PoolerSpecTemplateSpecResourcesClaims `field:"optional" json:"claims" yaml:"claims"`
	// Limits describes the maximum amount of compute resources allowed.
	//
	// More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	Limits *map[string]PoolerSpecTemplateSpecResourcesLimits `field:"optional" json:"limits" yaml:"limits"`
	// Requests describes the minimum amount of compute resources required.
	//
	// If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,
	// otherwise to an implementation-defined value. Requests cannot exceed Limits.
	// More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	Requests *map[string]PoolerSpecTemplateSpecResourcesRequests `field:"optional" json:"requests" yaml:"requests"`
}

