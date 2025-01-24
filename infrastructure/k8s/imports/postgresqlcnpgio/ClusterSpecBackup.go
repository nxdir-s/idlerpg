package postgresqlcnpgio


// The configuration to be used for backups.
type ClusterSpecBackup struct {
	// The configuration for the barman-cloud tool suite.
	BarmanObjectStore *ClusterSpecBackupBarmanObjectStore `field:"optional" json:"barmanObjectStore" yaml:"barmanObjectStore"`
	// RetentionPolicy is the retention policy to be used for backups and WALs (i.e. '60d'). The retention policy is expressed in the form of `XXu` where `XX` is a positive integer and `u` is in `[dwm]` - days, weeks, months. It's currently only applicable when using the BarmanObjectStore method.
	RetentionPolicy *string `field:"optional" json:"retentionPolicy" yaml:"retentionPolicy"`
	// The policy to decide which instance should perform backups.
	//
	// Available
	// options are empty string, which will default to `prefer-standby` policy,
	// `primary` to have backups run always on primary instances, `prefer-standby`
	// to have backups run preferably on the most updated standby, if available.
	Target ClusterSpecBackupTarget `field:"optional" json:"target" yaml:"target"`
	// VolumeSnapshot provides the configuration for the execution of volume snapshot backups.
	VolumeSnapshot *ClusterSpecBackupVolumeSnapshot `field:"optional" json:"volumeSnapshot" yaml:"volumeSnapshot"`
}

