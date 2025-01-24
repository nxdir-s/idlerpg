package postgresqlcnpgio


// SecurityContext defines the security options the container should be run with.
//
// If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext.
// More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
type PoolerSpecTemplateSpecContainersSecurityContext struct {
	// AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process.
	//
	// This bool directly controls if
	// the no_new_privs flag will be set on the container process.
	// AllowPrivilegeEscalation is true always when the container is:
	// 1) run as Privileged
	// 2) has CAP_SYS_ADMIN
	// Note that this field cannot be set when spec.os.name is windows.
	AllowPrivilegeEscalation *bool `field:"optional" json:"allowPrivilegeEscalation" yaml:"allowPrivilegeEscalation"`
	// appArmorProfile is the AppArmor options to use by this container.
	//
	// If set, this profile
	// overrides the pod's appArmorProfile.
	// Note that this field cannot be set when spec.os.name is windows.
	AppArmorProfile *PoolerSpecTemplateSpecContainersSecurityContextAppArmorProfile `field:"optional" json:"appArmorProfile" yaml:"appArmorProfile"`
	// The capabilities to add/drop when running containers.
	//
	// Defaults to the default set of capabilities granted by the container runtime.
	// Note that this field cannot be set when spec.os.name is windows.
	// Default: the default set of capabilities granted by the container runtime.
	//
	Capabilities *PoolerSpecTemplateSpecContainersSecurityContextCapabilities `field:"optional" json:"capabilities" yaml:"capabilities"`
	// Run container in privileged mode.
	//
	// Processes in privileged containers are essentially equivalent to root on the host.
	// Defaults to false.
	// Note that this field cannot be set when spec.os.name is windows.
	// Default: false.
	//
	Privileged *bool `field:"optional" json:"privileged" yaml:"privileged"`
	// procMount denotes the type of proc mount to use for the containers.
	//
	// The default value is Default which uses the container runtime defaults for
	// readonly paths and masked paths.
	// This requires the ProcMountType feature flag to be enabled.
	// Note that this field cannot be set when spec.os.name is windows.
	ProcMount *string `field:"optional" json:"procMount" yaml:"procMount"`
	// Whether this container has a read-only root filesystem.
	//
	// Default is false.
	// Note that this field cannot be set when spec.os.name is windows.
	// Default: false.
	//
	ReadOnlyRootFilesystem *bool `field:"optional" json:"readOnlyRootFilesystem" yaml:"readOnlyRootFilesystem"`
	// The GID to run the entrypoint of the container process.
	//
	// Uses runtime default if unset.
	// May also be set in PodSecurityContext.  If set in both SecurityContext and
	// PodSecurityContext, the value specified in SecurityContext takes precedence.
	// Note that this field cannot be set when spec.os.name is windows.
	RunAsGroup *float64 `field:"optional" json:"runAsGroup" yaml:"runAsGroup"`
	// Indicates that the container must run as a non-root user.
	//
	// If true, the Kubelet will validate the image at runtime to ensure that it
	// does not run as UID 0 (root) and fail to start the container if it does.
	// If unset or false, no such validation will be performed.
	// May also be set in PodSecurityContext.  If set in both SecurityContext and
	// PodSecurityContext, the value specified in SecurityContext takes precedence.
	RunAsNonRoot *bool `field:"optional" json:"runAsNonRoot" yaml:"runAsNonRoot"`
	// The UID to run the entrypoint of the container process.
	//
	// Defaults to user specified in image metadata if unspecified.
	// May also be set in PodSecurityContext.  If set in both SecurityContext and
	// PodSecurityContext, the value specified in SecurityContext takes precedence.
	// Note that this field cannot be set when spec.os.name is windows.
	// Default: user specified in image metadata if unspecified.
	//
	RunAsUser *float64 `field:"optional" json:"runAsUser" yaml:"runAsUser"`
	// The seccomp options to use by this container.
	//
	// If seccomp options are
	// provided at both the pod & container level, the container options
	// override the pod options.
	// Note that this field cannot be set when spec.os.name is windows.
	SeccompProfile *PoolerSpecTemplateSpecContainersSecurityContextSeccompProfile `field:"optional" json:"seccompProfile" yaml:"seccompProfile"`
	// The SELinux context to be applied to the container.
	//
	// If unspecified, the container runtime will allocate a random SELinux context for each
	// container.  May also be set in PodSecurityContext.  If set in both SecurityContext and
	// PodSecurityContext, the value specified in SecurityContext takes precedence.
	// Note that this field cannot be set when spec.os.name is windows.
	SeLinuxOptions *PoolerSpecTemplateSpecContainersSecurityContextSeLinuxOptions `field:"optional" json:"seLinuxOptions" yaml:"seLinuxOptions"`
	// The Windows specific settings applied to all containers.
	//
	// If unspecified, the options from the PodSecurityContext will be used.
	// If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	// Note that this field cannot be set when spec.os.name is linux.
	WindowsOptions *PoolerSpecTemplateSpecContainersSecurityContextWindowsOptions `field:"optional" json:"windowsOptions" yaml:"windowsOptions"`
}

