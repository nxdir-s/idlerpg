package postgresqlcnpgio


// Bind+Search authentication configuration.
type ClusterSpecPostgresqlLdapBindSearchAuth struct {
	// Root DN to begin the user search.
	BaseDn *string `field:"optional" json:"baseDn" yaml:"baseDn"`
	// DN of the user to bind to the directory.
	BindDn *string `field:"optional" json:"bindDn" yaml:"bindDn"`
	// Secret with the password for the user to bind to the directory.
	BindPassword *ClusterSpecPostgresqlLdapBindSearchAuthBindPassword `field:"optional" json:"bindPassword" yaml:"bindPassword"`
	// Attribute to match against the username.
	SearchAttribute *string `field:"optional" json:"searchAttribute" yaml:"searchAttribute"`
	// Search filter to use when doing the search+bind authentication.
	SearchFilter *string `field:"optional" json:"searchFilter" yaml:"searchFilter"`
}

