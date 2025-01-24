package postgresqlcnpgio


// Specification of the desired behavior of the ScheduledBackup.
//
// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
type ScheduledBackupSpec struct {
	// The cluster to backup.
	Cluster *ScheduledBackupSpecCluster `field:"required" json:"cluster" yaml:"cluster"`
	// The schedule does not follow the same format used in Kubernetes CronJobs as it includes an additional seconds specifier, see https://pkg.go.dev/github.com/robfig/cron#hdr-CRON_Expression_Format.
	Schedule *string `field:"required" json:"schedule" yaml:"schedule"`
	// Indicates which ownerReference should be put inside the created backup resources.<br /> - none: no owner reference for created backup objects (same behavior as before the field was introduced)<br /> - self: sets the Scheduled backup object as owner of the backup<br /> - cluster: set the cluster as owner of the backup<br />.
	BackupOwnerReference ScheduledBackupSpecBackupOwnerReference `field:"optional" json:"backupOwnerReference" yaml:"backupOwnerReference"`
	// If the first backup has to be immediately start after creation or not.
	Immediate *bool `field:"optional" json:"immediate" yaml:"immediate"`
	// The backup method to be used, possible options are `barmanObjectStore`, `volumeSnapshot` or `plugin`.
	//
	// Defaults to: `barmanObjectStore`.
	// Default: barmanObjectStore`.
	//
	Method ScheduledBackupSpecMethod `field:"optional" json:"method" yaml:"method"`
	// Whether the default type of backup with volume snapshots is online/hot (`true`, default) or offline/cold (`false`) Overrides the default setting specified in the cluster field '.spec.backup.volumeSnapshot.online'.
	Online *bool `field:"optional" json:"online" yaml:"online"`
	// Configuration parameters to control the online/hot backup with volume snapshots Overrides the default settings specified in the cluster '.backup.volumeSnapshot.onlineConfiguration' stanza.
	OnlineConfiguration *ScheduledBackupSpecOnlineConfiguration `field:"optional" json:"onlineConfiguration" yaml:"onlineConfiguration"`
	// Configuration parameters passed to the plugin managing this backup.
	PluginConfiguration *ScheduledBackupSpecPluginConfiguration `field:"optional" json:"pluginConfiguration" yaml:"pluginConfiguration"`
	// If this backup is suspended or not.
	Suspend *bool `field:"optional" json:"suspend" yaml:"suspend"`
	// The policy to decide which instance should perform this backup.
	//
	// If empty,
	// it defaults to `cluster.spec.backup.target`.
	// Available options are empty string, `primary` and `prefer-standby`.
	// `primary` to have backups run always on primary instances,
	// `prefer-standby` to have backups run preferably on the most updated
	// standby, if available.
	Target ScheduledBackupSpecTarget `field:"optional" json:"target" yaml:"target"`
}

