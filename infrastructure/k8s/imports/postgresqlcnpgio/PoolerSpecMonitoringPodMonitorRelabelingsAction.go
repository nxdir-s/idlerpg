package postgresqlcnpgio


// Action to perform based on the regex matching.
//
// `Uppercase` and `Lowercase` actions require Prometheus >= v2.36.0.
// `DropEqual` and `KeepEqual` actions require Prometheus >= v2.41.0.
//
// Default: "Replace".
type PoolerSpecMonitoringPodMonitorRelabelingsAction string

const (
	// replace.
	PoolerSpecMonitoringPodMonitorRelabelingsAction_REPLACE PoolerSpecMonitoringPodMonitorRelabelingsAction = "REPLACE"
	// keep.
	PoolerSpecMonitoringPodMonitorRelabelingsAction_KEEP PoolerSpecMonitoringPodMonitorRelabelingsAction = "KEEP"
	// drop.
	PoolerSpecMonitoringPodMonitorRelabelingsAction_DROP PoolerSpecMonitoringPodMonitorRelabelingsAction = "DROP"
	// hashmod.
	PoolerSpecMonitoringPodMonitorRelabelingsAction_HASHMOD PoolerSpecMonitoringPodMonitorRelabelingsAction = "HASHMOD"
	// labelmap.
	PoolerSpecMonitoringPodMonitorRelabelingsAction_LABELMAP PoolerSpecMonitoringPodMonitorRelabelingsAction = "LABELMAP"
	// labeldrop.
	PoolerSpecMonitoringPodMonitorRelabelingsAction_LABELDROP PoolerSpecMonitoringPodMonitorRelabelingsAction = "LABELDROP"
	// labelkeep.
	PoolerSpecMonitoringPodMonitorRelabelingsAction_LABELKEEP PoolerSpecMonitoringPodMonitorRelabelingsAction = "LABELKEEP"
	// lowercase.
	PoolerSpecMonitoringPodMonitorRelabelingsAction_LOWERCASE PoolerSpecMonitoringPodMonitorRelabelingsAction = "LOWERCASE"
	// uppercase.
	PoolerSpecMonitoringPodMonitorRelabelingsAction_UPPERCASE PoolerSpecMonitoringPodMonitorRelabelingsAction = "UPPERCASE"
	// keepequal.
	PoolerSpecMonitoringPodMonitorRelabelingsAction_KEEPEQUAL PoolerSpecMonitoringPodMonitorRelabelingsAction = "KEEPEQUAL"
	// dropequal.
	PoolerSpecMonitoringPodMonitorRelabelingsAction_DROPEQUAL PoolerSpecMonitoringPodMonitorRelabelingsAction = "DROPEQUAL"
)

