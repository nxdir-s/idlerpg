package postgresqlcnpgio


// PublicationTargetObject is an object to publish.
type PublicationSpecTargetObjects struct {
	// Specifies a list of tables to add to the publication.
	//
	// Corresponding
	// to `FOR TABLE` in PostgreSQL.
	Table *PublicationSpecTargetObjectsTable `field:"optional" json:"table" yaml:"table"`
	// Marks the publication as one that replicates changes for all tables in the specified list of schemas, including tables created in the future.
	//
	// Corresponding to `FOR TABLES IN SCHEMA` in PostgreSQL.
	TablesInSchema *string `field:"optional" json:"tablesInSchema" yaml:"tablesInSchema"`
}

