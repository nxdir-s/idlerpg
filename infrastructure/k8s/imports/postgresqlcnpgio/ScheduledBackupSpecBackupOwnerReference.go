package postgresqlcnpgio


// Indicates which ownerReference should be put inside the created backup resources.<br /> - none: no owner reference for created backup objects (same behavior as before the field was introduced)<br /> - self: sets the Scheduled backup object as owner of the backup<br /> - cluster: set the cluster as owner of the backup<br />.
type ScheduledBackupSpecBackupOwnerReference string

const (
	// none.
	ScheduledBackupSpecBackupOwnerReference_NONE ScheduledBackupSpecBackupOwnerReference = "NONE"
	// self.
	ScheduledBackupSpecBackupOwnerReference_SELF ScheduledBackupSpecBackupOwnerReference = "SELF"
	// cluster.
	ScheduledBackupSpecBackupOwnerReference_CLUSTER ScheduledBackupSpecBackupOwnerReference = "CLUSTER"
)

