package postgresqlcnpgio


// Bootstrap the cluster from a backup.
type ClusterSpecBootstrapRecovery struct {
	// The backup object containing the physical base backup from which to initiate the recovery procedure.
	//
	// Mutually exclusive with `source` and `volumeSnapshots`.
	Backup *ClusterSpecBootstrapRecoveryBackup `field:"optional" json:"backup" yaml:"backup"`
	// Name of the database used by the application.
	//
	// Default: `app`.
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Name of the owner of the database in the instance to be used by applications.
	//
	// Defaults to the value of the `database` key.
	// Default: the value of the `database` key.
	//
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// By default, the recovery process applies all the available WAL files in the archive (full recovery).
	//
	// However, you can also
	// end the recovery as soon as a consistent state is reached or
	// recover to a point-in-time (PITR) by specifying a `RecoveryTarget` object,
	// as expected by PostgreSQL (i.e., timestamp, transaction Id, LSN, ...).
	// More info: https://www.postgresql.org/docs/current/runtime-config-wal.html#RUNTIME-CONFIG-WAL-RECOVERY-TARGET
	RecoveryTarget *ClusterSpecBootstrapRecoveryRecoveryTarget `field:"optional" json:"recoveryTarget" yaml:"recoveryTarget"`
	// Name of the secret containing the initial credentials for the owner of the user database.
	//
	// If empty a new secret will be
	// created from scratch.
	Secret *ClusterSpecBootstrapRecoverySecret `field:"optional" json:"secret" yaml:"secret"`
	// The external cluster whose backup we will restore.
	//
	// This is also
	// used as the name of the folder under which the backup is stored,
	// so it must be set to the name of the source cluster
	// Mutually exclusive with `backup`.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// The static PVC data source(s) from which to initiate the recovery procedure.
	//
	// Currently supporting `VolumeSnapshot`
	// and `PersistentVolumeClaim` resources that map an existing
	// PVC group, compatible with CloudNativePG, and taken with
	// a cold backup copy on a fenced Postgres instance (limitation
	// which will be removed in the future when online backup
	// will be implemented).
	// Mutually exclusive with `backup`.
	VolumeSnapshots *ClusterSpecBootstrapRecoveryVolumeSnapshots `field:"optional" json:"volumeSnapshots" yaml:"volumeSnapshots"`
}

