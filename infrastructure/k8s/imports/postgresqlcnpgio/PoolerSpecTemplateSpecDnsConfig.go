package postgresqlcnpgio


// Specifies the DNS parameters of a pod.
//
// Parameters specified here will be merged to the generated DNS
// configuration based on DNSPolicy.
type PoolerSpecTemplateSpecDnsConfig struct {
	// A list of DNS name server IP addresses.
	//
	// This will be appended to the base nameservers generated from DNSPolicy.
	// Duplicated nameservers will be removed.
	Nameservers *[]*string `field:"optional" json:"nameservers" yaml:"nameservers"`
	// A list of DNS resolver options.
	//
	// This will be merged with the base options generated from DNSPolicy.
	// Duplicated entries will be removed. Resolution options given in Options
	// will override those that appear in the base DNSPolicy.
	Options *[]*PoolerSpecTemplateSpecDnsConfigOptions `field:"optional" json:"options" yaml:"options"`
	// A list of DNS search domains for host-name lookup.
	//
	// This will be appended to the base search paths generated from DNSPolicy.
	// Duplicated search paths will be removed.
	Searches *[]*string `field:"optional" json:"searches" yaml:"searches"`
}

