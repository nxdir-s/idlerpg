package postgresqlcnpgio


// resources represents the minimum resources the volume should have.
//
// If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements
// that are lower than previous value but must still be higher than capacity recorded in the
// status field of the claim.
// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources
type ClusterSpecTablespacesStoragePvcTemplateResources struct {
	// Limits describes the maximum amount of compute resources allowed.
	//
	// More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	Limits *map[string]ClusterSpecTablespacesStoragePvcTemplateResourcesLimits `field:"optional" json:"limits" yaml:"limits"`
	// Requests describes the minimum amount of compute resources required.
	//
	// If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,
	// otherwise to an implementation-defined value. Requests cannot exceed Limits.
	// More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	Requests *map[string]ClusterSpecTablespacesStoragePvcTemplateResourcesRequests `field:"optional" json:"requests" yaml:"requests"`
}

