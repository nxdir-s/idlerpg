package postgresqlcnpgio


// PodDNSConfigOption defines DNS resolver options of a pod.
type PoolerSpecTemplateSpecDnsConfigOptions struct {
	// Name is this DNS resolver option's name.
	//
	// Required.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Value is this DNS resolver option's value.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

