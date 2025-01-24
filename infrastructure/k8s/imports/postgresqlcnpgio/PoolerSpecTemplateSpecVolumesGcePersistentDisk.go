package postgresqlcnpgio


// gcePersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod.
//
// Deprecated: GCEPersistentDisk is deprecated. All operations for the in-tree
// gcePersistentDisk type are redirected to the pd.csi.storage.gke.io CSI driver.
// More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
type PoolerSpecTemplateSpecVolumesGcePersistentDisk struct {
	// pdName is unique name of the PD resource in GCE.
	//
	// Used to identify the disk in GCE.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	PdName *string `field:"required" json:"pdName" yaml:"pdName"`
	// fsType is filesystem type of the volume that you want to mount.
	//
	// Tip: Ensure that the filesystem type is supported by the host operating system.
	// Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// partition is the partition in the volume that you want to mount.
	//
	// If omitted, the default is to mount by volume name.
	// Examples: For volume /dev/sda1, you specify the partition as "1".
	// Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty).
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	Partition *float64 `field:"optional" json:"partition" yaml:"partition"`
	// readOnly here will force the ReadOnly setting in VolumeMounts.
	//
	// Defaults to false.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	// Default: false.
	//
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}

