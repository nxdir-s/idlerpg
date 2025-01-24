package postgresqlcnpgio


// The configuration of the monitoring infrastructure of this cluster.
type ClusterSpecMonitoring struct {
	// The list of config maps containing the custom queries.
	CustomQueriesConfigMap *[]*ClusterSpecMonitoringCustomQueriesConfigMap `field:"optional" json:"customQueriesConfigMap" yaml:"customQueriesConfigMap"`
	// The list of secrets containing the custom queries.
	CustomQueriesSecret *[]*ClusterSpecMonitoringCustomQueriesSecret `field:"optional" json:"customQueriesSecret" yaml:"customQueriesSecret"`
	// Whether the default queries should be injected.
	//
	// Set it to `true` if you don't want to inject default queries into the cluster.
	// Default: false.
	DisableDefaultQueries *bool `field:"optional" json:"disableDefaultQueries" yaml:"disableDefaultQueries"`
	// Enable or disable the `PodMonitor`.
	EnablePodMonitor *bool `field:"optional" json:"enablePodMonitor" yaml:"enablePodMonitor"`
	// The list of metric relabelings for the `PodMonitor`.
	//
	// Applied to samples before ingestion.
	PodMonitorMetricRelabelings *[]*ClusterSpecMonitoringPodMonitorMetricRelabelings `field:"optional" json:"podMonitorMetricRelabelings" yaml:"podMonitorMetricRelabelings"`
	// The list of relabelings for the `PodMonitor`.
	//
	// Applied to samples before scraping.
	PodMonitorRelabelings *[]*ClusterSpecMonitoringPodMonitorRelabelings `field:"optional" json:"podMonitorRelabelings" yaml:"podMonitorRelabelings"`
	// Configure TLS communication for the metrics endpoint.
	//
	// Changing tls.enabled option will force a rollout of all instances.
	Tls *ClusterSpecMonitoringTls `field:"optional" json:"tls" yaml:"tls"`
}

