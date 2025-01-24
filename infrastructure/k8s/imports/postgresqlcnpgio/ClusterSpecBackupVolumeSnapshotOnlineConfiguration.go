package postgresqlcnpgio


// Configuration parameters to control the online/hot backup with volume snapshots.
type ClusterSpecBackupVolumeSnapshotOnlineConfiguration struct {
	// Control whether the I/O workload for the backup initial checkpoint will be limited, according to the `checkpoint_completion_target` setting on the PostgreSQL server.
	//
	// If set to true, an immediate checkpoint will be
	// used, meaning PostgreSQL will complete the checkpoint as soon as
	// possible. `false` by default.
	ImmediateCheckpoint *bool `field:"optional" json:"immediateCheckpoint" yaml:"immediateCheckpoint"`
	// If false, the function will return immediately after the backup is completed, without waiting for WAL to be archived.
	//
	// This behavior is only useful with backup software that independently monitors WAL archiving.
	// Otherwise, WAL required to make the backup consistent might be missing and make the backup useless.
	// By default, or when this parameter is true, pg_backup_stop will wait for WAL to be archived when archiving is
	// enabled.
	// On a standby, this means that it will wait only when archive_mode = always.
	// If write activity on the primary is low, it may be useful to run pg_switch_wal on the primary in order to trigger
	// an immediate segment switch.
	WaitForArchive *bool `field:"optional" json:"waitForArchive" yaml:"waitForArchive"`
}

