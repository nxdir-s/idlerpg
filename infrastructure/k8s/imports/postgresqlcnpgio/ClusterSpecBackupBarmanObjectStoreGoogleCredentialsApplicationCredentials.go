package postgresqlcnpgio


// The secret containing the Google Cloud Storage JSON file with the credentials.
type ClusterSpecBackupBarmanObjectStoreGoogleCredentialsApplicationCredentials struct {
	// The key to select.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Name of the referent.
	Name *string `field:"required" json:"name" yaml:"name"`
}

