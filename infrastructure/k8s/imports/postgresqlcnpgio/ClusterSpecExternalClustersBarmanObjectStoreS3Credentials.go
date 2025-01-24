package postgresqlcnpgio


// The credentials to use to upload data to S3.
type ClusterSpecExternalClustersBarmanObjectStoreS3Credentials struct {
	// The reference to the access key id.
	AccessKeyId *ClusterSpecExternalClustersBarmanObjectStoreS3CredentialsAccessKeyId `field:"optional" json:"accessKeyId" yaml:"accessKeyId"`
	// Use the role based authentication without providing explicitly the keys.
	InheritFromIamRole *bool `field:"optional" json:"inheritFromIamRole" yaml:"inheritFromIamRole"`
	// The reference to the secret containing the region name.
	Region *ClusterSpecExternalClustersBarmanObjectStoreS3CredentialsRegion `field:"optional" json:"region" yaml:"region"`
	// The reference to the secret access key.
	SecretAccessKey *ClusterSpecExternalClustersBarmanObjectStoreS3CredentialsSecretAccessKey `field:"optional" json:"secretAccessKey" yaml:"secretAccessKey"`
	// The references to the session key.
	SessionToken *ClusterSpecExternalClustersBarmanObjectStoreS3CredentialsSessionToken `field:"optional" json:"sessionToken" yaml:"sessionToken"`
}

