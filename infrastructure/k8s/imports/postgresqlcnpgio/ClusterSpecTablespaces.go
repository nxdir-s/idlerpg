package postgresqlcnpgio


// TablespaceConfiguration is the configuration of a tablespace, and includes the storage specification for the tablespace.
type ClusterSpecTablespaces struct {
	// The name of the tablespace.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The storage configuration for the tablespace.
	Storage *ClusterSpecTablespacesStorage `field:"required" json:"storage" yaml:"storage"`
	// Owner is the PostgreSQL user owning the tablespace.
	Owner *ClusterSpecTablespacesOwner `field:"optional" json:"owner" yaml:"owner"`
	// When set to true, the tablespace will be added as a `temp_tablespaces` entry in PostgreSQL, and will be available to automatically house temp database objects, or other temporary files.
	//
	// Please refer to PostgreSQL
	// documentation for more information on the `temp_tablespaces` GUC.
	Temporary *bool `field:"optional" json:"temporary" yaml:"temporary"`
}

