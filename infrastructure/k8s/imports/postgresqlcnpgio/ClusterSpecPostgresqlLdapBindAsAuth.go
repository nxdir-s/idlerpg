package postgresqlcnpgio


// Bind as authentication configuration.
type ClusterSpecPostgresqlLdapBindAsAuth struct {
	// Prefix for the bind authentication option.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Suffix for the bind authentication option.
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

