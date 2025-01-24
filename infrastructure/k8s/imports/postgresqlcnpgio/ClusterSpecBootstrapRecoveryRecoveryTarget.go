package postgresqlcnpgio


// By default, the recovery process applies all the available WAL files in the archive (full recovery).
//
// However, you can also
// end the recovery as soon as a consistent state is reached or
// recover to a point-in-time (PITR) by specifying a `RecoveryTarget` object,
// as expected by PostgreSQL (i.e., timestamp, transaction Id, LSN, ...).
// More info: https://www.postgresql.org/docs/current/runtime-config-wal.html#RUNTIME-CONFIG-WAL-RECOVERY-TARGET
type ClusterSpecBootstrapRecoveryRecoveryTarget struct {
	// The ID of the backup from which to start the recovery process.
	//
	// If empty (default) the operator will automatically detect the backup
	// based on targetTime or targetLSN if specified. Otherwise use the
	// latest available backup in chronological order.
	BackupId *string `field:"optional" json:"backupId" yaml:"backupId"`
	// Set the target to be exclusive.
	//
	// If omitted, defaults to false, so that
	// in Postgres, `recovery_target_inclusive` will be true.
	Exclusive *bool `field:"optional" json:"exclusive" yaml:"exclusive"`
	// End recovery as soon as a consistent state is reached.
	TargetImmediate *bool `field:"optional" json:"targetImmediate" yaml:"targetImmediate"`
	// The target LSN (Log Sequence Number).
	TargetLsn *string `field:"optional" json:"targetLsn" yaml:"targetLsn"`
	// The target name (to be previously created with `pg_create_restore_point`).
	TargetName *string `field:"optional" json:"targetName" yaml:"targetName"`
	// The target time as a timestamp in the RFC3339 standard.
	TargetTime *string `field:"optional" json:"targetTime" yaml:"targetTime"`
	// The target timeline ("latest" or a positive integer).
	TargetTli *string `field:"optional" json:"targetTli" yaml:"targetTli"`
	// The target transaction ID.
	TargetXid *string `field:"optional" json:"targetXid" yaml:"targetXid"`
}

