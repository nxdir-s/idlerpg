package postgresqlcnpgio


// vsphereVolume represents a vSphere volume attached and mounted on kubelets host machine.
//
// Deprecated: VsphereVolume is deprecated. All operations for the in-tree vsphereVolume type
// are redirected to the csi.vsphere.vmware.com CSI driver.
type PoolerSpecTemplateSpecVolumesVsphereVolume struct {
	// volumePath is the path that identifies vSphere volume vmdk.
	VolumePath *string `field:"required" json:"volumePath" yaml:"volumePath"`
	// fsType is filesystem type to mount.
	//
	// Must be a filesystem type supported by the host operating system.
	// Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// storagePolicyID is the storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.
	StoragePolicyId *string `field:"optional" json:"storagePolicyId" yaml:"storagePolicyId"`
	// storagePolicyName is the storage Policy Based Management (SPBM) profile name.
	StoragePolicyName *string `field:"optional" json:"storagePolicyName" yaml:"storagePolicyName"`
}

