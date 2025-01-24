package postgresqlcnpgio


// The configuration for the barman-cloud tool suite.
type ClusterSpecBackupBarmanObjectStore struct {
	// The path where to store the backup (i.e. s3://bucket/path/to/folder) this path, with different destination folders, will be used for WALs and for data.
	DestinationPath *string `field:"required" json:"destinationPath" yaml:"destinationPath"`
	// The credentials to use to upload data to Azure Blob Storage.
	AzureCredentials *ClusterSpecBackupBarmanObjectStoreAzureCredentials `field:"optional" json:"azureCredentials" yaml:"azureCredentials"`
	// The configuration to be used to backup the data files When not defined, base backups files will be stored uncompressed and may be unencrypted in the object store, according to the bucket default policy.
	Data *ClusterSpecBackupBarmanObjectStoreData `field:"optional" json:"data" yaml:"data"`
	// EndpointCA store the CA bundle of the barman endpoint.
	//
	// Useful when using self-signed certificates to avoid
	// errors with certificate issuer and barman-cloud-wal-archive.
	EndpointCa *ClusterSpecBackupBarmanObjectStoreEndpointCa `field:"optional" json:"endpointCa" yaml:"endpointCa"`
	// Endpoint to be used to upload data to the cloud, overriding the automatic endpoint discovery.
	EndpointUrl *string `field:"optional" json:"endpointUrl" yaml:"endpointUrl"`
	// The credentials to use to upload data to Google Cloud Storage.
	GoogleCredentials *ClusterSpecBackupBarmanObjectStoreGoogleCredentials `field:"optional" json:"googleCredentials" yaml:"googleCredentials"`
	// HistoryTags is a list of key value pairs that will be passed to the Barman --history-tags option.
	HistoryTags *map[string]*string `field:"optional" json:"historyTags" yaml:"historyTags"`
	// The credentials to use to upload data to S3.
	S3Credentials *ClusterSpecBackupBarmanObjectStoreS3Credentials `field:"optional" json:"s3Credentials" yaml:"s3Credentials"`
	// The server name on S3, the cluster name is used if this parameter is omitted.
	ServerName *string `field:"optional" json:"serverName" yaml:"serverName"`
	// Tags is a list of key value pairs that will be passed to the Barman --tags option.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The configuration for the backup of the WAL stream.
	//
	// When not defined, WAL files will be stored uncompressed and may be
	// unencrypted in the object store, according to the bucket default policy.
	Wal *ClusterSpecBackupBarmanObjectStoreWal `field:"optional" json:"wal" yaml:"wal"`
}

