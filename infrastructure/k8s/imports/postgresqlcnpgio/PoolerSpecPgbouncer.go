package postgresqlcnpgio


// The PgBouncer configuration.
type PoolerSpecPgbouncer struct {
	// The query that will be used to download the hash of the password of a certain user.
	//
	// Default: "SELECT usename, passwd FROM public.user_search($1)".
	// In case it is specified, also an AuthQuerySecret has to be specified and
	// no automatic CNPG Cluster integration will be triggered.
	AuthQuery *string `field:"optional" json:"authQuery" yaml:"authQuery"`
	// The credentials of the user that need to be used for the authentication query.
	//
	// In case it is specified, also an AuthQuery
	// (e.g. "SELECT usename, passwd FROM pg_catalog.pg_shadow WHERE usename=$1")
	// has to be specified and no automatic CNPG Cluster integration will be triggered.
	AuthQuerySecret *PoolerSpecPgbouncerAuthQuerySecret `field:"optional" json:"authQuerySecret" yaml:"authQuerySecret"`
	// Additional parameters to be passed to PgBouncer - please check the CNPG documentation for a list of options you can configure.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// When set to `true`, PgBouncer will disconnect from the PostgreSQL server, first waiting for all queries to complete, and pause all new client connections until this value is set to `false` (default).
	//
	// Internally,
	// the operator calls PgBouncer's `PAUSE` and `RESUME` commands.
	Paused *bool `field:"optional" json:"paused" yaml:"paused"`
	// PostgreSQL Host Based Authentication rules (lines to be appended to the pg_hba.conf file).
	PgHba *[]*string `field:"optional" json:"pgHba" yaml:"pgHba"`
	// The pool mode.
	//
	// Default: `session`.
	PoolMode PoolerSpecPgbouncerPoolMode `field:"optional" json:"poolMode" yaml:"poolMode"`
}

