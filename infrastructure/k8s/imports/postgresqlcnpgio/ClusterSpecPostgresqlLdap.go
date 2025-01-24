package postgresqlcnpgio


// Options to specify LDAP configuration.
type ClusterSpecPostgresqlLdap struct {
	// Bind as authentication configuration.
	BindAsAuth *ClusterSpecPostgresqlLdapBindAsAuth `field:"optional" json:"bindAsAuth" yaml:"bindAsAuth"`
	// Bind+Search authentication configuration.
	BindSearchAuth *ClusterSpecPostgresqlLdapBindSearchAuth `field:"optional" json:"bindSearchAuth" yaml:"bindSearchAuth"`
	// LDAP server port.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// LDAP schema to be used, possible options are `ldap` and `ldaps`.
	Scheme ClusterSpecPostgresqlLdapScheme `field:"optional" json:"scheme" yaml:"scheme"`
	// LDAP hostname or IP address.
	Server *string `field:"optional" json:"server" yaml:"server"`
	// Set to 'true' to enable LDAP over TLS.
	//
	// 'false' is default.
	Tls *bool `field:"optional" json:"tls" yaml:"tls"`
}

