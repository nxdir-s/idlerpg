package postgresqlcnpgio


// fc represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.
type PoolerSpecTemplateSpecVolumesFc struct {
	// fsType is the filesystem type to mount.
	//
	// Must be a filesystem type supported by the host operating system.
	// Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// lun is Optional: FC target lun number.
	Lun *float64 `field:"optional" json:"lun" yaml:"lun"`
	// readOnly is Optional: Defaults to false (read/write).
	//
	// ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// Default: false (read/write). ReadOnly here will force
	//
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	// targetWWNs is Optional: FC target worldwide names (WWNs).
	TargetWwNs *[]*string `field:"optional" json:"targetWwNs" yaml:"targetWwNs"`
	// wwids Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.
	Wwids *[]*string `field:"optional" json:"wwids" yaml:"wwids"`
}

