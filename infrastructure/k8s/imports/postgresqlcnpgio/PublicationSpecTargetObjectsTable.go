package postgresqlcnpgio


// Specifies a list of tables to add to the publication.
//
// Corresponding
// to `FOR TABLE` in PostgreSQL.
type PublicationSpecTargetObjectsTable struct {
	// The table name.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The columns to publish.
	Columns *[]*string `field:"optional" json:"columns" yaml:"columns"`
	// Whether to limit to the table only or include all its descendants.
	Only *bool `field:"optional" json:"only" yaml:"only"`
	// The schema name.
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

