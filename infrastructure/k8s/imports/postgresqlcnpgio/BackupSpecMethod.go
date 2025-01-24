package postgresqlcnpgio


// The backup method to be used, possible options are `barmanObjectStore`, `volumeSnapshot` or `plugin`.
//
// Defaults to: `barmanObjectStore`.
// Default: barmanObjectStore`.
//
type BackupSpecMethod string

const (
	// barmanObjectStore.
	BackupSpecMethod_BARMAN_OBJECT_STORE BackupSpecMethod = "BARMAN_OBJECT_STORE"
	// volumeSnapshot.
	BackupSpecMethod_VOLUME_SNAPSHOT BackupSpecMethod = "VOLUME_SNAPSHOT"
	// plugin.
	BackupSpecMethod_PLUGIN BackupSpecMethod = "PLUGIN"
)

