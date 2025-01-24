package postgresqlcnpgio


// The storage account where to upload data.
type ClusterSpecExternalClustersBarmanObjectStoreAzureCredentialsStorageAccount struct {
	// The key to select.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Name of the referent.
	Name *string `field:"required" json:"name" yaml:"name"`
}

