package postgresqlcnpgio


// The policy to decide which instance should perform backups.
//
// Available
// options are empty string, which will default to `prefer-standby` policy,
// `primary` to have backups run always on primary instances, `prefer-standby`
// to have backups run preferably on the most updated standby, if available.
type ClusterSpecBackupTarget string

const (
	// primary.
	ClusterSpecBackupTarget_PRIMARY ClusterSpecBackupTarget = "PRIMARY"
	// prefer-standby.
	ClusterSpecBackupTarget_PREFER_HYPHEN_STANDBY ClusterSpecBackupTarget = "PREFER_HYPHEN_STANDBY"
)

