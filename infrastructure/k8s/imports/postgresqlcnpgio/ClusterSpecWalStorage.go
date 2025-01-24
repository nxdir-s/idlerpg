package postgresqlcnpgio


// Configuration of the storage for PostgreSQL WAL (Write-Ahead Log).
type ClusterSpecWalStorage struct {
	// Template to be used to generate the Persistent Volume Claim.
	PvcTemplate *ClusterSpecWalStoragePvcTemplate `field:"optional" json:"pvcTemplate" yaml:"pvcTemplate"`
	// Resize existent PVCs, defaults to true.
	ResizeInUseVolumes *bool `field:"optional" json:"resizeInUseVolumes" yaml:"resizeInUseVolumes"`
	// Size of the storage.
	//
	// Required if not already specified in the PVC template.
	// Changes to this field are automatically reapplied to the created PVCs.
	// Size cannot be decreased.
	Size *string `field:"optional" json:"size" yaml:"size"`
	// StorageClass to use for PVCs.
	//
	// Applied after
	// evaluating the PVC template, if available.
	// If not specified, the generated PVCs will use the
	// default storage class.
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
}

