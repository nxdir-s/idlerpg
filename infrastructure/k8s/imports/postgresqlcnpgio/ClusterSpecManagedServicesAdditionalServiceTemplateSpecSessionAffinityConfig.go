package postgresqlcnpgio


// sessionAffinityConfig contains the configurations of session affinity.
type ClusterSpecManagedServicesAdditionalServiceTemplateSpecSessionAffinityConfig struct {
	// clientIP contains the configurations of Client IP based session affinity.
	ClientIp *ClusterSpecManagedServicesAdditionalServiceTemplateSpecSessionAffinityConfigClientIp `field:"optional" json:"clientIp" yaml:"clientIp"`
}

