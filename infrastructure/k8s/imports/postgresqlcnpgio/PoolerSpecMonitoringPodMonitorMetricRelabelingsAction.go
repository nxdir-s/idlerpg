package postgresqlcnpgio


// Action to perform based on the regex matching.
//
// `Uppercase` and `Lowercase` actions require Prometheus >= v2.36.0.
// `DropEqual` and `KeepEqual` actions require Prometheus >= v2.41.0.
//
// Default: "Replace".
type PoolerSpecMonitoringPodMonitorMetricRelabelingsAction string

const (
	// replace.
	PoolerSpecMonitoringPodMonitorMetricRelabelingsAction_REPLACE PoolerSpecMonitoringPodMonitorMetricRelabelingsAction = "REPLACE"
	// keep.
	PoolerSpecMonitoringPodMonitorMetricRelabelingsAction_KEEP PoolerSpecMonitoringPodMonitorMetricRelabelingsAction = "KEEP"
	// drop.
	PoolerSpecMonitoringPodMonitorMetricRelabelingsAction_DROP PoolerSpecMonitoringPodMonitorMetricRelabelingsAction = "DROP"
	// hashmod.
	PoolerSpecMonitoringPodMonitorMetricRelabelingsAction_HASHMOD PoolerSpecMonitoringPodMonitorMetricRelabelingsAction = "HASHMOD"
	// labelmap.
	PoolerSpecMonitoringPodMonitorMetricRelabelingsAction_LABELMAP PoolerSpecMonitoringPodMonitorMetricRelabelingsAction = "LABELMAP"
	// labeldrop.
	PoolerSpecMonitoringPodMonitorMetricRelabelingsAction_LABELDROP PoolerSpecMonitoringPodMonitorMetricRelabelingsAction = "LABELDROP"
	// labelkeep.
	PoolerSpecMonitoringPodMonitorMetricRelabelingsAction_LABELKEEP PoolerSpecMonitoringPodMonitorMetricRelabelingsAction = "LABELKEEP"
	// lowercase.
	PoolerSpecMonitoringPodMonitorMetricRelabelingsAction_LOWERCASE PoolerSpecMonitoringPodMonitorMetricRelabelingsAction = "LOWERCASE"
	// uppercase.
	PoolerSpecMonitoringPodMonitorMetricRelabelingsAction_UPPERCASE PoolerSpecMonitoringPodMonitorMetricRelabelingsAction = "UPPERCASE"
	// keepequal.
	PoolerSpecMonitoringPodMonitorMetricRelabelingsAction_KEEPEQUAL PoolerSpecMonitoringPodMonitorMetricRelabelingsAction = "KEEPEQUAL"
	// dropequal.
	PoolerSpecMonitoringPodMonitorMetricRelabelingsAction_DROPEQUAL PoolerSpecMonitoringPodMonitorMetricRelabelingsAction = "DROPEQUAL"
)

