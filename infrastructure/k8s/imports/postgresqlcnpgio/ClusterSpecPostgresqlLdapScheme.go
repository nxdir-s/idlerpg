package postgresqlcnpgio


// LDAP schema to be used, possible options are `ldap` and `ldaps`.
type ClusterSpecPostgresqlLdapScheme string

const (
	// ldap.
	ClusterSpecPostgresqlLdapScheme_LDAP ClusterSpecPostgresqlLdapScheme = "LDAP"
	// ldaps.
	ClusterSpecPostgresqlLdapScheme_LDAPS ClusterSpecPostgresqlLdapScheme = "LDAPS"
)

