package postgresqlcnpgio


// The reference to the secret access key.
type ClusterSpecBackupBarmanObjectStoreS3CredentialsSecretAccessKey struct {
	// The key to select.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Name of the referent.
	Name *string `field:"required" json:"name" yaml:"name"`
}

