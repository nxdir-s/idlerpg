package postgresqlcnpgio


// The configuration of the monitoring infrastructure of this pooler.
type PoolerSpecMonitoring struct {
	// Enable or disable the `PodMonitor`.
	EnablePodMonitor *bool `field:"optional" json:"enablePodMonitor" yaml:"enablePodMonitor"`
	// The list of metric relabelings for the `PodMonitor`.
	//
	// Applied to samples before ingestion.
	PodMonitorMetricRelabelings *[]*PoolerSpecMonitoringPodMonitorMetricRelabelings `field:"optional" json:"podMonitorMetricRelabelings" yaml:"podMonitorMetricRelabelings"`
	// The list of relabelings for the `PodMonitor`.
	//
	// Applied to samples before scraping.
	PodMonitorRelabelings *[]*PoolerSpecMonitoringPodMonitorRelabelings `field:"optional" json:"podMonitorRelabelings" yaml:"podMonitorRelabelings"`
}

