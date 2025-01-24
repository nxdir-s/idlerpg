package postgresqlcnpgio


// Bootstrap the cluster via initdb.
type ClusterSpecBootstrapInitdb struct {
	// Specifies the locale name when the builtin provider is used.
	//
	// This option requires `localeProvider` to be set to `builtin`.
	// Available from PostgreSQL 17.
	BuiltinLocale *string `field:"optional" json:"builtinLocale" yaml:"builtinLocale"`
	// Name of the database used by the application.
	//
	// Default: `app`.
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Whether the `-k` option should be passed to initdb, enabling checksums on data pages (default: `false`).
	DataChecksums *bool `field:"optional" json:"dataChecksums" yaml:"dataChecksums"`
	// The value to be passed as option `--encoding` for initdb (default:`UTF8`).
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// Specifies the ICU locale when the ICU provider is used.
	//
	// This option requires `localeProvider` to be set to `icu`.
	// Available from PostgreSQL 15.
	IcuLocale *string `field:"optional" json:"icuLocale" yaml:"icuLocale"`
	// Specifies additional collation rules to customize the behavior of the default collation.
	//
	// This option requires `localeProvider` to be set to `icu`.
	// Available from PostgreSQL 16.
	IcuRules *string `field:"optional" json:"icuRules" yaml:"icuRules"`
	// Bootstraps the new cluster by importing data from an existing PostgreSQL instance using logical backup (`pg_dump` and `pg_restore`).
	Import *ClusterSpecBootstrapInitdbImport `field:"optional" json:"import" yaml:"import"`
	// Sets the default collation order and character classification in the new database.
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
	// The value to be passed as option `--lc-collate` for initdb (default:`C`).
	LocaleCollate *string `field:"optional" json:"localeCollate" yaml:"localeCollate"`
	// The value to be passed as option `--lc-ctype` for initdb (default:`C`).
	LocaleCType *string `field:"optional" json:"localeCType" yaml:"localeCType"`
	// This option sets the locale provider for databases created in the new cluster.
	//
	// Available from PostgreSQL 16.
	LocaleProvider *string `field:"optional" json:"localeProvider" yaml:"localeProvider"`
	// The list of options that must be passed to initdb when creating the cluster.
	//
	// Deprecated: This could lead to inconsistent configurations,
	// please use the explicit provided parameters instead.
	// If defined, explicit values will be ignored.
	Options *[]*string `field:"optional" json:"options" yaml:"options"`
	// Name of the owner of the database in the instance to be used by applications.
	//
	// Defaults to the value of the `database` key.
	// Default: the value of the `database` key.
	//
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// List of SQL queries to be executed as a superuser in the application database right after the cluster has been created - to be used with extreme care (by default empty).
	PostInitApplicationSql *[]*string `field:"optional" json:"postInitApplicationSql" yaml:"postInitApplicationSql"`
	// List of references to ConfigMaps or Secrets containing SQL files to be executed as a superuser in the application database right after the cluster has been created.
	//
	// The references are processed in a specific order:
	// first, all Secrets are processed, followed by all ConfigMaps.
	// Within each group, the processing order follows the sequence specified
	// in their respective arrays.
	// (by default empty).
	PostInitApplicationSqlRefs *ClusterSpecBootstrapInitdbPostInitApplicationSqlRefs `field:"optional" json:"postInitApplicationSqlRefs" yaml:"postInitApplicationSqlRefs"`
	// List of SQL queries to be executed as a superuser in the `postgres` database right after the cluster has been created - to be used with extreme care (by default empty).
	PostInitSql *[]*string `field:"optional" json:"postInitSql" yaml:"postInitSql"`
	// List of references to ConfigMaps or Secrets containing SQL files to be executed as a superuser in the `postgres` database right after the cluster has been created.
	//
	// The references are processed in a specific order:
	// first, all Secrets are processed, followed by all ConfigMaps.
	// Within each group, the processing order follows the sequence specified
	// in their respective arrays.
	// (by default empty).
	PostInitSqlRefs *ClusterSpecBootstrapInitdbPostInitSqlRefs `field:"optional" json:"postInitSqlRefs" yaml:"postInitSqlRefs"`
	// List of SQL queries to be executed as a superuser in the `template1` database right after the cluster has been created - to be used with extreme care (by default empty).
	PostInitTemplateSql *[]*string `field:"optional" json:"postInitTemplateSql" yaml:"postInitTemplateSql"`
	// List of references to ConfigMaps or Secrets containing SQL files to be executed as a superuser in the `template1` database right after the cluster has been created.
	//
	// The references are processed in a specific order:
	// first, all Secrets are processed, followed by all ConfigMaps.
	// Within each group, the processing order follows the sequence specified
	// in their respective arrays.
	// (by default empty).
	PostInitTemplateSqlRefs *ClusterSpecBootstrapInitdbPostInitTemplateSqlRefs `field:"optional" json:"postInitTemplateSqlRefs" yaml:"postInitTemplateSqlRefs"`
	// Name of the secret containing the initial credentials for the owner of the user database.
	//
	// If empty a new secret will be
	// created from scratch.
	Secret *ClusterSpecBootstrapInitdbSecret `field:"optional" json:"secret" yaml:"secret"`
	// The value in megabytes (1 to 1024) to be passed to the `--wal-segsize` option for initdb (default: empty, resulting in PostgreSQL default: 16MB).
	WalSegmentSize *float64 `field:"optional" json:"walSegmentSize" yaml:"walSegmentSize"`
}

