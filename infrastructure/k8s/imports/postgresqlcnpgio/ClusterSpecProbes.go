package postgresqlcnpgio


// The configuration of the probes to be injected in the PostgreSQL Pods.
type ClusterSpecProbes struct {
	// The liveness probe configuration.
	Liveness *ClusterSpecProbesLiveness `field:"optional" json:"liveness" yaml:"liveness"`
	// The readiness probe configuration.
	Readiness *ClusterSpecProbesReadiness `field:"optional" json:"readiness" yaml:"readiness"`
	// The startup probe configuration.
	Startup *ClusterSpecProbesStartup `field:"optional" json:"startup" yaml:"startup"`
}

