package postgresqlcnpgio


// A shared-access-signature to be used in conjunction with the storage account name.
type ClusterSpecExternalClustersBarmanObjectStoreAzureCredentialsStorageSasToken struct {
	// The key to select.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Name of the referent.
	Name *string `field:"required" json:"name" yaml:"name"`
}

