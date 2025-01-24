package postgresqlcnpgio


// The credentials to use to upload data to Azure Blob Storage.
type ClusterSpecExternalClustersBarmanObjectStoreAzureCredentials struct {
	// The connection string to be used.
	ConnectionString *ClusterSpecExternalClustersBarmanObjectStoreAzureCredentialsConnectionString `field:"optional" json:"connectionString" yaml:"connectionString"`
	// Use the Azure AD based authentication without providing explicitly the keys.
	InheritFromAzureAd *bool `field:"optional" json:"inheritFromAzureAd" yaml:"inheritFromAzureAd"`
	// The storage account where to upload data.
	StorageAccount *ClusterSpecExternalClustersBarmanObjectStoreAzureCredentialsStorageAccount `field:"optional" json:"storageAccount" yaml:"storageAccount"`
	// The storage account key to be used in conjunction with the storage account name.
	StorageKey *ClusterSpecExternalClustersBarmanObjectStoreAzureCredentialsStorageKey `field:"optional" json:"storageKey" yaml:"storageKey"`
	// A shared-access-signature to be used in conjunction with the storage account name.
	StorageSasToken *ClusterSpecExternalClustersBarmanObjectStoreAzureCredentialsStorageSasToken `field:"optional" json:"storageSasToken" yaml:"storageSasToken"`
}

