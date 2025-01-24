package postgresqlcnpgio


// The storage account key to be used in conjunction with the storage account name.
type ClusterSpecBackupBarmanObjectStoreAzureCredentialsStorageKey struct {
	// The key to select.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Name of the referent.
	Name *string `field:"required" json:"name" yaml:"name"`
}

