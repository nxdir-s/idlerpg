package postgresqlcnpgio


// SnapshotOwnerReference indicates the type of owner reference the snapshot should have.
type ClusterSpecBackupVolumeSnapshotSnapshotOwnerReference string

const (
	// none.
	ClusterSpecBackupVolumeSnapshotSnapshotOwnerReference_NONE ClusterSpecBackupVolumeSnapshotSnapshotOwnerReference = "NONE"
	// cluster.
	ClusterSpecBackupVolumeSnapshotSnapshotOwnerReference_CLUSTER ClusterSpecBackupVolumeSnapshotSnapshotOwnerReference = "CLUSTER"
	// backup.
	ClusterSpecBackupVolumeSnapshotSnapshotOwnerReference_BACKUP ClusterSpecBackupVolumeSnapshotSnapshotOwnerReference = "BACKUP"
)

