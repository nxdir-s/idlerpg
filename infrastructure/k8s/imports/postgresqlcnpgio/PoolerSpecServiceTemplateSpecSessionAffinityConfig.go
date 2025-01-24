package postgresqlcnpgio


// sessionAffinityConfig contains the configurations of session affinity.
type PoolerSpecServiceTemplateSpecSessionAffinityConfig struct {
	// clientIP contains the configurations of Client IP based session affinity.
	ClientIp *PoolerSpecServiceTemplateSpecSessionAffinityConfigClientIp `field:"optional" json:"clientIp" yaml:"clientIp"`
}

