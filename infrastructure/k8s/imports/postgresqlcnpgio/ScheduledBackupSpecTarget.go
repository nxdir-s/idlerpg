package postgresqlcnpgio


// The policy to decide which instance should perform this backup.
//
// If empty,
// it defaults to `cluster.spec.backup.target`.
// Available options are empty string, `primary` and `prefer-standby`.
// `primary` to have backups run always on primary instances,
// `prefer-standby` to have backups run preferably on the most updated
// standby, if available.
type ScheduledBackupSpecTarget string

const (
	// primary.
	ScheduledBackupSpecTarget_PRIMARY ScheduledBackupSpecTarget = "PRIMARY"
	// prefer-standby.
	ScheduledBackupSpecTarget_PREFER_HYPHEN_STANDBY ScheduledBackupSpecTarget = "PREFER_HYPHEN_STANDBY"
)

