package postgresqlcnpgio


// PublicationSpec defines the desired state of Publication.
type PublicationSpec struct {
	// The name of the PostgreSQL cluster that identifies the "publisher".
	Cluster *PublicationSpecCluster `field:"required" json:"cluster" yaml:"cluster"`
	// The name of the database where the publication will be installed in the "publisher" cluster.
	Dbname *string `field:"required" json:"dbname" yaml:"dbname"`
	// The name of the publication inside PostgreSQL.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Target of the publication as expected by PostgreSQL `CREATE PUBLICATION` command.
	Target *PublicationSpecTarget `field:"required" json:"target" yaml:"target"`
	// Publication parameters part of the `WITH` clause as expected by PostgreSQL `CREATE PUBLICATION` command.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// The policy for end-of-life maintenance of this publication.
	PublicationReclaimPolicy PublicationSpecPublicationReclaimPolicy `field:"optional" json:"publicationReclaimPolicy" yaml:"publicationReclaimPolicy"`
}

