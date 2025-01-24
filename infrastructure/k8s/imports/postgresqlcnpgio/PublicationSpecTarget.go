package postgresqlcnpgio


// Target of the publication as expected by PostgreSQL `CREATE PUBLICATION` command.
type PublicationSpecTarget struct {
	// Marks the publication as one that replicates changes for all tables in the database, including tables created in the future.
	//
	// Corresponding to `FOR ALL TABLES` in PostgreSQL.
	AllTables *bool `field:"optional" json:"allTables" yaml:"allTables"`
	// Just the following schema objects.
	Objects *[]*PublicationSpecTargetObjects `field:"optional" json:"objects" yaml:"objects"`
}

