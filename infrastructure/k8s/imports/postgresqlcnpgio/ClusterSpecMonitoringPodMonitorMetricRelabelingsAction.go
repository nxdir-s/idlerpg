package postgresqlcnpgio


// Action to perform based on the regex matching.
//
// `Uppercase` and `Lowercase` actions require Prometheus >= v2.36.0.
// `DropEqual` and `KeepEqual` actions require Prometheus >= v2.41.0.
//
// Default: "Replace".
type ClusterSpecMonitoringPodMonitorMetricRelabelingsAction string

const (
	// replace.
	ClusterSpecMonitoringPodMonitorMetricRelabelingsAction_REPLACE ClusterSpecMonitoringPodMonitorMetricRelabelingsAction = "REPLACE"
	// keep.
	ClusterSpecMonitoringPodMonitorMetricRelabelingsAction_KEEP ClusterSpecMonitoringPodMonitorMetricRelabelingsAction = "KEEP"
	// drop.
	ClusterSpecMonitoringPodMonitorMetricRelabelingsAction_DROP ClusterSpecMonitoringPodMonitorMetricRelabelingsAction = "DROP"
	// hashmod.
	ClusterSpecMonitoringPodMonitorMetricRelabelingsAction_HASHMOD ClusterSpecMonitoringPodMonitorMetricRelabelingsAction = "HASHMOD"
	// labelmap.
	ClusterSpecMonitoringPodMonitorMetricRelabelingsAction_LABELMAP ClusterSpecMonitoringPodMonitorMetricRelabelingsAction = "LABELMAP"
	// labeldrop.
	ClusterSpecMonitoringPodMonitorMetricRelabelingsAction_LABELDROP ClusterSpecMonitoringPodMonitorMetricRelabelingsAction = "LABELDROP"
	// labelkeep.
	ClusterSpecMonitoringPodMonitorMetricRelabelingsAction_LABELKEEP ClusterSpecMonitoringPodMonitorMetricRelabelingsAction = "LABELKEEP"
	// lowercase.
	ClusterSpecMonitoringPodMonitorMetricRelabelingsAction_LOWERCASE ClusterSpecMonitoringPodMonitorMetricRelabelingsAction = "LOWERCASE"
	// uppercase.
	ClusterSpecMonitoringPodMonitorMetricRelabelingsAction_UPPERCASE ClusterSpecMonitoringPodMonitorMetricRelabelingsAction = "UPPERCASE"
	// keepequal.
	ClusterSpecMonitoringPodMonitorMetricRelabelingsAction_KEEPEQUAL ClusterSpecMonitoringPodMonitorMetricRelabelingsAction = "KEEPEQUAL"
	// dropequal.
	ClusterSpecMonitoringPodMonitorMetricRelabelingsAction_DROPEQUAL ClusterSpecMonitoringPodMonitorMetricRelabelingsAction = "DROPEQUAL"
)

