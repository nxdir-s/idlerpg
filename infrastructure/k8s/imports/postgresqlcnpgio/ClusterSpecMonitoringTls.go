package postgresqlcnpgio


// Configure TLS communication for the metrics endpoint.
//
// Changing tls.enabled option will force a rollout of all instances.
type ClusterSpecMonitoringTls struct {
	// Enable TLS for the monitoring endpoint.
	//
	// Changing this option will force a rollout of all instances.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

