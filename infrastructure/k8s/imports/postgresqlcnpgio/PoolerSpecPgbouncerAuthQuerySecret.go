package postgresqlcnpgio


// The credentials of the user that need to be used for the authentication query.
//
// In case it is specified, also an AuthQuery
// (e.g. "SELECT usename, passwd FROM pg_catalog.pg_shadow WHERE usename=$1")
// has to be specified and no automatic CNPG Cluster integration will be triggered.
type PoolerSpecPgbouncerAuthQuerySecret struct {
	// Name of the referent.
	Name *string `field:"required" json:"name" yaml:"name"`
}

