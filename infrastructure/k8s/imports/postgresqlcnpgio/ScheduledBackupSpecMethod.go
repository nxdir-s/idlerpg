package postgresqlcnpgio


// The backup method to be used, possible options are `barmanObjectStore`, `volumeSnapshot` or `plugin`.
//
// Defaults to: `barmanObjectStore`.
// Default: barmanObjectStore`.
//
type ScheduledBackupSpecMethod string

const (
	// barmanObjectStore.
	ScheduledBackupSpecMethod_BARMAN_OBJECT_STORE ScheduledBackupSpecMethod = "BARMAN_OBJECT_STORE"
	// volumeSnapshot.
	ScheduledBackupSpecMethod_VOLUME_SNAPSHOT ScheduledBackupSpecMethod = "VOLUME_SNAPSHOT"
	// plugin.
	ScheduledBackupSpecMethod_PLUGIN ScheduledBackupSpecMethod = "PLUGIN"
)

