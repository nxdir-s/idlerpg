package postgresqlcnpgio


// The policy to decide which instance should perform this backup.
//
// If empty,
// it defaults to `cluster.spec.backup.target`.
// Available options are empty string, `primary` and `prefer-standby`.
// `primary` to have backups run always on primary instances,
// `prefer-standby` to have backups run preferably on the most updated
// standby, if available.
type BackupSpecTarget string

const (
	// primary.
	BackupSpecTarget_PRIMARY BackupSpecTarget = "PRIMARY"
	// prefer-standby.
	BackupSpecTarget_PREFER_HYPHEN_STANDBY BackupSpecTarget = "PREFER_HYPHEN_STANDBY"
)

