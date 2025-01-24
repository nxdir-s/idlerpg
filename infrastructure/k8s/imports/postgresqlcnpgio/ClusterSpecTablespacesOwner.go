package postgresqlcnpgio


// Owner is the PostgreSQL user owning the tablespace.
type ClusterSpecTablespacesOwner struct {
	Name *string `field:"optional" json:"name" yaml:"name"`
}

