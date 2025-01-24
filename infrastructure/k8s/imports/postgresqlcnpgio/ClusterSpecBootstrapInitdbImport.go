package postgresqlcnpgio


// Bootstraps the new cluster by importing data from an existing PostgreSQL instance using logical backup (`pg_dump` and `pg_restore`).
type ClusterSpecBootstrapInitdbImport struct {
	// The databases to import.
	Databases *[]*string `field:"required" json:"databases" yaml:"databases"`
	// The source of the import.
	Source *ClusterSpecBootstrapInitdbImportSource `field:"required" json:"source" yaml:"source"`
	// The import type.
	//
	// Can be `microservice` or `monolith`.
	Type ClusterSpecBootstrapInitdbImportType `field:"required" json:"type" yaml:"type"`
	// List of custom options to pass to the `pg_dump` command.
	//
	// IMPORTANT:
	// Use these options with caution and at your own risk, as the operator
	// does not validate their content. Be aware that certain options may
	// conflict with the operator's intended functionality or design.
	PgDumpExtraOptions *[]*string `field:"optional" json:"pgDumpExtraOptions" yaml:"pgDumpExtraOptions"`
	// List of custom options to pass to the `pg_restore` command.
	//
	// IMPORTANT:
	// Use these options with caution and at your own risk, as the operator
	// does not validate their content. Be aware that certain options may
	// conflict with the operator's intended functionality or design.
	PgRestoreExtraOptions *[]*string `field:"optional" json:"pgRestoreExtraOptions" yaml:"pgRestoreExtraOptions"`
	// List of SQL queries to be executed as a superuser in the application database right after is imported - to be used with extreme care (by default empty).
	//
	// Only available in microservice type.
	PostImportApplicationSql *[]*string `field:"optional" json:"postImportApplicationSql" yaml:"postImportApplicationSql"`
	// The roles to import.
	Roles *[]*string `field:"optional" json:"roles" yaml:"roles"`
	// When set to true, only the `pre-data` and `post-data` sections of `pg_restore` are invoked, avoiding data import.
	//
	// Default: `false`.
	SchemaOnly *bool `field:"optional" json:"schemaOnly" yaml:"schemaOnly"`
}

