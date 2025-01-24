package postgresqlcnpgio


// VolumeSnapshot provides the configuration for the execution of volume snapshot backups.
type ClusterSpecBackupVolumeSnapshot struct {
	// Annotations key-value pairs that will be added to .metadata.annotations snapshot resources.
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// ClassName specifies the Snapshot Class to be used for PG_DATA PersistentVolumeClaim.
	//
	// It is the default class for the other types if no specific class is present.
	ClassName *string `field:"optional" json:"className" yaml:"className"`
	// Labels are key-value pairs that will be added to .metadata.labels snapshot resources.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Whether the default type of backup with volume snapshots is online/hot (`true`, default) or offline/cold (`false`).
	Online *bool `field:"optional" json:"online" yaml:"online"`
	// Configuration parameters to control the online/hot backup with volume snapshots.
	OnlineConfiguration *ClusterSpecBackupVolumeSnapshotOnlineConfiguration `field:"optional" json:"onlineConfiguration" yaml:"onlineConfiguration"`
	// SnapshotOwnerReference indicates the type of owner reference the snapshot should have.
	SnapshotOwnerReference ClusterSpecBackupVolumeSnapshotSnapshotOwnerReference `field:"optional" json:"snapshotOwnerReference" yaml:"snapshotOwnerReference"`
	// TablespaceClassName specifies the Snapshot Class to be used for the tablespaces.
	//
	// defaults to the PGDATA Snapshot Class, if set.
	TablespaceClassName *map[string]*string `field:"optional" json:"tablespaceClassName" yaml:"tablespaceClassName"`
	// WalClassName specifies the Snapshot Class to be used for the PG_WAL PersistentVolumeClaim.
	WalClassName *string `field:"optional" json:"walClassName" yaml:"walClassName"`
}

