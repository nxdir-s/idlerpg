package postgresqlcnpgio


// Specifies the OS of the containers in the pod. Some pod and container fields are restricted if this is set.
//
// If the OS field is set to linux, the following fields must be unset:
// -securityContext.windowsOptions
//
// If the OS field is set to windows, following fields must be unset:
// - spec.hostPID
// - spec.hostIPC
// - spec.hostUsers
// - spec.securityContext.appArmorProfile
// - spec.securityContext.seLinuxOptions
// - spec.securityContext.seccompProfile
// - spec.securityContext.fsGroup
// - spec.securityContext.fsGroupChangePolicy
// - spec.securityContext.sysctls
// - spec.shareProcessNamespace
// - spec.securityContext.runAsUser
// - spec.securityContext.runAsGroup
// - spec.securityContext.supplementalGroups
// - spec.securityContext.supplementalGroupsPolicy
// - spec.containers[*].securityContext.appArmorProfile
// - spec.containers[*].securityContext.seLinuxOptions
// - spec.containers[*].securityContext.seccompProfile
// - spec.containers[*].securityContext.capabilities
// - spec.containers[*].securityContext.readOnlyRootFilesystem
// - spec.containers[*].securityContext.privileged
// - spec.containers[*].securityContext.allowPrivilegeEscalation
// - spec.containers[*].securityContext.procMount
// - spec.containers[*].securityContext.runAsUser
// - spec.containers[*].securityContext.runAsGroup
type PoolerSpecTemplateSpecOs struct {
	// Name is the name of the operating system.
	//
	// The currently supported values are linux and windows.
	// Additional value may be defined in future and can be one of:
	// https://github.com/opencontainers/runtime-spec/blob/master/config.md#platform-specific-configuration
	// Clients should expect to handle additional values and treat unrecognized values in this field as os: null.
	Name *string `field:"required" json:"name" yaml:"name"`
}

