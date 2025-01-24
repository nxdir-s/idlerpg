package postgresqlcnpgio


// Specification of the desired behavior of the backup.
//
// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
type BackupSpec struct {
	// The cluster to backup.
	Cluster *BackupSpecCluster `field:"required" json:"cluster" yaml:"cluster"`
	// The backup method to be used, possible options are `barmanObjectStore`, `volumeSnapshot` or `plugin`.
	//
	// Defaults to: `barmanObjectStore`.
	// Default: barmanObjectStore`.
	//
	Method BackupSpecMethod `field:"optional" json:"method" yaml:"method"`
	// Whether the default type of backup with volume snapshots is online/hot (`true`, default) or offline/cold (`false`) Overrides the default setting specified in the cluster field '.spec.backup.volumeSnapshot.online'.
	Online *bool `field:"optional" json:"online" yaml:"online"`
	// Configuration parameters to control the online/hot backup with volume snapshots Overrides the default settings specified in the cluster '.backup.volumeSnapshot.onlineConfiguration' stanza.
	OnlineConfiguration *BackupSpecOnlineConfiguration `field:"optional" json:"onlineConfiguration" yaml:"onlineConfiguration"`
	// Configuration parameters passed to the plugin managing this backup.
	PluginConfiguration *BackupSpecPluginConfiguration `field:"optional" json:"pluginConfiguration" yaml:"pluginConfiguration"`
	// The policy to decide which instance should perform this backup.
	//
	// If empty,
	// it defaults to `cluster.spec.backup.target`.
	// Available options are empty string, `primary` and `prefer-standby`.
	// `primary` to have backups run always on primary instances,
	// `prefer-standby` to have backups run preferably on the most updated
	// standby, if available.
	Target BackupSpecTarget `field:"optional" json:"target" yaml:"target"`
}

