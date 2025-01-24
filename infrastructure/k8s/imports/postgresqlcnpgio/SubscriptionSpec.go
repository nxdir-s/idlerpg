package postgresqlcnpgio


// SubscriptionSpec defines the desired state of Subscription.
type SubscriptionSpec struct {
	// The name of the PostgreSQL cluster that identifies the "subscriber".
	Cluster *SubscriptionSpecCluster `field:"required" json:"cluster" yaml:"cluster"`
	// The name of the database where the publication will be installed in the "subscriber" cluster.
	Dbname *string `field:"required" json:"dbname" yaml:"dbname"`
	// The name of the external cluster with the publication ("publisher").
	ExternalClusterName *string `field:"required" json:"externalClusterName" yaml:"externalClusterName"`
	// The name of the subscription inside PostgreSQL.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the publication inside the PostgreSQL database in the "publisher".
	PublicationName *string `field:"required" json:"publicationName" yaml:"publicationName"`
	// Subscription parameters part of the `WITH` clause as expected by PostgreSQL `CREATE SUBSCRIPTION` command.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// The name of the database containing the publication on the external cluster.
	//
	// Defaults to the one in the external cluster definition.
	// Default: the one in the external cluster definition.
	//
	PublicationDbName *string `field:"optional" json:"publicationDbName" yaml:"publicationDbName"`
	// The policy for end-of-life maintenance of this subscription.
	SubscriptionReclaimPolicy SubscriptionSpecSubscriptionReclaimPolicy `field:"optional" json:"subscriptionReclaimPolicy" yaml:"subscriptionReclaimPolicy"`
}

