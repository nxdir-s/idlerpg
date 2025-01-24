package postgresqlcnpgio


// Action to perform based on the regex matching.
//
// `Uppercase` and `Lowercase` actions require Prometheus >= v2.36.0.
// `DropEqual` and `KeepEqual` actions require Prometheus >= v2.41.0.
//
// Default: "Replace".
type ClusterSpecMonitoringPodMonitorRelabelingsAction string

const (
	// replace.
	ClusterSpecMonitoringPodMonitorRelabelingsAction_REPLACE ClusterSpecMonitoringPodMonitorRelabelingsAction = "REPLACE"
	// keep.
	ClusterSpecMonitoringPodMonitorRelabelingsAction_KEEP ClusterSpecMonitoringPodMonitorRelabelingsAction = "KEEP"
	// drop.
	ClusterSpecMonitoringPodMonitorRelabelingsAction_DROP ClusterSpecMonitoringPodMonitorRelabelingsAction = "DROP"
	// hashmod.
	ClusterSpecMonitoringPodMonitorRelabelingsAction_HASHMOD ClusterSpecMonitoringPodMonitorRelabelingsAction = "HASHMOD"
	// labelmap.
	ClusterSpecMonitoringPodMonitorRelabelingsAction_LABELMAP ClusterSpecMonitoringPodMonitorRelabelingsAction = "LABELMAP"
	// labeldrop.
	ClusterSpecMonitoringPodMonitorRelabelingsAction_LABELDROP ClusterSpecMonitoringPodMonitorRelabelingsAction = "LABELDROP"
	// labelkeep.
	ClusterSpecMonitoringPodMonitorRelabelingsAction_LABELKEEP ClusterSpecMonitoringPodMonitorRelabelingsAction = "LABELKEEP"
	// lowercase.
	ClusterSpecMonitoringPodMonitorRelabelingsAction_LOWERCASE ClusterSpecMonitoringPodMonitorRelabelingsAction = "LOWERCASE"
	// uppercase.
	ClusterSpecMonitoringPodMonitorRelabelingsAction_UPPERCASE ClusterSpecMonitoringPodMonitorRelabelingsAction = "UPPERCASE"
	// keepequal.
	ClusterSpecMonitoringPodMonitorRelabelingsAction_KEEPEQUAL ClusterSpecMonitoringPodMonitorRelabelingsAction = "KEEPEQUAL"
	// dropequal.
	ClusterSpecMonitoringPodMonitorRelabelingsAction_DROPEQUAL ClusterSpecMonitoringPodMonitorRelabelingsAction = "DROPEQUAL"
)

