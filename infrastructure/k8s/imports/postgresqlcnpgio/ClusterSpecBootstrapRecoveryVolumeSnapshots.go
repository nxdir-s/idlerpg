package postgresqlcnpgio


// The static PVC data source(s) from which to initiate the recovery procedure.
//
// Currently supporting `VolumeSnapshot`
// and `PersistentVolumeClaim` resources that map an existing
// PVC group, compatible with CloudNativePG, and taken with
// a cold backup copy on a fenced Postgres instance (limitation
// which will be removed in the future when online backup
// will be implemented).
// Mutually exclusive with `backup`.
type ClusterSpecBootstrapRecoveryVolumeSnapshots struct {
	// Configuration of the storage of the instances.
	Storage *ClusterSpecBootstrapRecoveryVolumeSnapshotsStorage `field:"required" json:"storage" yaml:"storage"`
	// Configuration of the storage for PostgreSQL tablespaces.
	TablespaceStorage *map[string]*ClusterSpecBootstrapRecoveryVolumeSnapshotsTablespaceStorage `field:"optional" json:"tablespaceStorage" yaml:"tablespaceStorage"`
	// Configuration of the storage for PostgreSQL WAL (Write-Ahead Log).
	WalStorage *ClusterSpecBootstrapRecoveryVolumeSnapshotsWalStorage `field:"optional" json:"walStorage" yaml:"walStorage"`
}

