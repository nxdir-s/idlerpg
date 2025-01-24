package postgresqlcnpgio


// The source of the import.
type ClusterSpecBootstrapInitdbImportSource struct {
	// The name of the externalCluster used for import.
	ExternalCluster *string `field:"required" json:"externalCluster" yaml:"externalCluster"`
}

