package postgresqlcnpgio


// csi (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers.
type PoolerSpecTemplateSpecVolumesCsi struct {
	// driver is the name of the CSI driver that handles this volume.
	//
	// Consult with your admin for the correct name as registered in the cluster.
	Driver *string `field:"required" json:"driver" yaml:"driver"`
	// fsType to mount.
	//
	// Ex. "ext4", "xfs", "ntfs".
	// If not provided, the empty value is passed to the associated CSI driver
	// which will determine the default filesystem to apply.
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// nodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls.
	//
	// This field is optional, and  may be empty if no secret is required. If the
	// secret object contains more than one secret, all secret references are passed.
	NodePublishSecretRef *PoolerSpecTemplateSpecVolumesCsiNodePublishSecretRef `field:"optional" json:"nodePublishSecretRef" yaml:"nodePublishSecretRef"`
	// readOnly specifies a read-only configuration for the volume.
	//
	// Defaults to false (read/write).
	// Default: false (read/write).
	//
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	// volumeAttributes stores driver-specific properties that are passed to the CSI driver.
	//
	// Consult your driver's documentation for supported values.
	VolumeAttributes *map[string]*string `field:"optional" json:"volumeAttributes" yaml:"volumeAttributes"`
}

