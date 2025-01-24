package postgresqlcnpgio


// The credentials to use to upload data to Google Cloud Storage.
type ClusterSpecExternalClustersBarmanObjectStoreGoogleCredentials struct {
	// The secret containing the Google Cloud Storage JSON file with the credentials.
	ApplicationCredentials *ClusterSpecExternalClustersBarmanObjectStoreGoogleCredentialsApplicationCredentials `field:"optional" json:"applicationCredentials" yaml:"applicationCredentials"`
	// If set to true, will presume that it's running inside a GKE environment, default to false.
	GkeEnvironment *bool `field:"optional" json:"gkeEnvironment" yaml:"gkeEnvironment"`
}

