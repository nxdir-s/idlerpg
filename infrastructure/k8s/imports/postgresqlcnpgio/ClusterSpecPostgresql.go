package postgresqlcnpgio


// Configuration of the PostgreSQL server.
type ClusterSpecPostgresql struct {
	// If this parameter is true, the user will be able to invoke `ALTER SYSTEM` on this CloudNativePG Cluster.
	//
	// This should only be used for debugging and troubleshooting.
	// Defaults to false.
	// Default: false.
	//
	EnableAlterSystem *bool `field:"optional" json:"enableAlterSystem" yaml:"enableAlterSystem"`
	// Options to specify LDAP configuration.
	Ldap *ClusterSpecPostgresqlLdap `field:"optional" json:"ldap" yaml:"ldap"`
	// PostgreSQL configuration options (postgresql.conf).
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// PostgreSQL Host Based Authentication rules (lines to be appended to the pg_hba.conf file).
	PgHba *[]*string `field:"optional" json:"pgHba" yaml:"pgHba"`
	// PostgreSQL User Name Maps rules (lines to be appended to the pg_ident.conf file).
	PgIdent *[]*string `field:"optional" json:"pgIdent" yaml:"pgIdent"`
	// Specifies the maximum number of seconds to wait when promoting an instance to primary.
	//
	// Default value is 40000000, greater than one year in seconds,
	// big enough to simulate an infinite timeout.
	PromotionTimeout *float64 `field:"optional" json:"promotionTimeout" yaml:"promotionTimeout"`
	// Lists of shared preload libraries to add to the default ones.
	SharedPreloadLibraries *[]*string `field:"optional" json:"sharedPreloadLibraries" yaml:"sharedPreloadLibraries"`
	// Configuration of the PostgreSQL synchronous replication feature.
	Synchronous *ClusterSpecPostgresqlSynchronous `field:"optional" json:"synchronous" yaml:"synchronous"`
	// Requirements to be met by sync replicas.
	//
	// This will affect how the "synchronous_standby_names" parameter will be
	// set up.
	SyncReplicaElectionConstraint *ClusterSpecPostgresqlSyncReplicaElectionConstraint `field:"optional" json:"syncReplicaElectionConstraint" yaml:"syncReplicaElectionConstraint"`
}

