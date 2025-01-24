package postgresqlcnpgio


// Specification of the desired Database.
//
// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
type DatabaseSpec struct {
	// The name of the PostgreSQL cluster hosting the database.
	Cluster *DatabaseSpecCluster `field:"required" json:"cluster" yaml:"cluster"`
	// The name of the database to create inside PostgreSQL.
	//
	// This setting cannot be changed.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Maps to the `OWNER` parameter of `CREATE DATABASE`.
	//
	// Maps to the `OWNER TO` command of `ALTER DATABASE`.
	// The role name of the user who owns the database inside PostgreSQL.
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// Maps to the `ALLOW_CONNECTIONS` parameter of `CREATE DATABASE` and `ALTER DATABASE`.
	//
	// If false then no one can connect to this database.
	AllowConnections *bool `field:"optional" json:"allowConnections" yaml:"allowConnections"`
	// Maps to the `BUILTIN_LOCALE` parameter of `CREATE DATABASE`.
	//
	// This
	// setting cannot be changed. Specifies the locale name when the
	// builtin provider is used. This option requires `localeProvider` to
	// be set to `builtin`. Available from PostgreSQL 17.
	BuiltinLocale *string `field:"optional" json:"builtinLocale" yaml:"builtinLocale"`
	// Maps to the `COLLATION_VERSION` parameter of `CREATE DATABASE`.
	//
	// This
	// setting cannot be changed.
	CollationVersion *string `field:"optional" json:"collationVersion" yaml:"collationVersion"`
	// Maps to the `CONNECTION LIMIT` clause of `CREATE DATABASE` and `ALTER DATABASE`.
	//
	// How many concurrent connections can be made to
	// this database. -1 (the default) means no limit.
	ConnectionLimit *float64 `field:"optional" json:"connectionLimit" yaml:"connectionLimit"`
	// The policy for end-of-life maintenance of this database.
	DatabaseReclaimPolicy DatabaseSpecDatabaseReclaimPolicy `field:"optional" json:"databaseReclaimPolicy" yaml:"databaseReclaimPolicy"`
	// Maps to the `ENCODING` parameter of `CREATE DATABASE`.
	//
	// This setting
	// cannot be changed. Character set encoding to use in the database.
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// Ensure the PostgreSQL database is `present` or `absent` - defaults to "present".
	Ensure DatabaseSpecEnsure `field:"optional" json:"ensure" yaml:"ensure"`
	// Maps to the `ICU_LOCALE` parameter of `CREATE DATABASE`.
	//
	// This
	// setting cannot be changed. Specifies the ICU locale when the ICU
	// provider is used. This option requires `localeProvider` to be set to
	// `icu`. Available from PostgreSQL 15.
	IcuLocale *string `field:"optional" json:"icuLocale" yaml:"icuLocale"`
	// Maps to the `ICU_RULES` parameter of `CREATE DATABASE`.
	//
	// This setting
	// cannot be changed. Specifies additional collation rules to customize
	// the behavior of the default collation. This option requires
	// `localeProvider` to be set to `icu`. Available from PostgreSQL 16.
	IcuRules *string `field:"optional" json:"icuRules" yaml:"icuRules"`
	// Maps to the `IS_TEMPLATE` parameter of `CREATE DATABASE` and `ALTER DATABASE`.
	//
	// If true, this database is considered a template and can
	// be cloned by any user with `CREATEDB` privileges.
	IsTemplate *bool `field:"optional" json:"isTemplate" yaml:"isTemplate"`
	// Maps to the `LOCALE` parameter of `CREATE DATABASE`.
	//
	// This setting
	// cannot be changed. Sets the default collation order and character
	// classification in the new database.
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
	// Maps to the `LC_COLLATE` parameter of `CREATE DATABASE`.
	//
	// This
	// setting cannot be changed.
	LocaleCollate *string `field:"optional" json:"localeCollate" yaml:"localeCollate"`
	// Maps to the `LC_CTYPE` parameter of `CREATE DATABASE`.
	//
	// This setting
	// cannot be changed.
	LocaleCType *string `field:"optional" json:"localeCType" yaml:"localeCType"`
	// Maps to the `LOCALE_PROVIDER` parameter of `CREATE DATABASE`.
	//
	// This
	// setting cannot be changed. This option sets the locale provider for
	// databases created in the new cluster. Available from PostgreSQL 16.
	LocaleProvider *string `field:"optional" json:"localeProvider" yaml:"localeProvider"`
	// Maps to the `TABLESPACE` parameter of `CREATE DATABASE`.
	//
	// Maps to the `SET TABLESPACE` command of `ALTER DATABASE`.
	// The name of the tablespace (in PostgreSQL) that will be associated
	// with the new database. This tablespace will be the default
	// tablespace used for objects created in this database.
	Tablespace *string `field:"optional" json:"tablespace" yaml:"tablespace"`
	// Maps to the `TEMPLATE` parameter of `CREATE DATABASE`.
	//
	// This setting
	// cannot be changed. The name of the template from which to create
	// this database.
	Template *string `field:"optional" json:"template" yaml:"template"`
}

