package postgresqlcnpgio


// Whenever to force the encryption of files (if the bucket is not already configured for that).
//
// Allowed options are empty string (use the bucket policy, default),
// `AES256` and `aws:kms`.
type ClusterSpecExternalClustersBarmanObjectStoreWalEncryption string

const (
	// AES256.
	ClusterSpecExternalClustersBarmanObjectStoreWalEncryption_AES256 ClusterSpecExternalClustersBarmanObjectStoreWalEncryption = "AES256"
	// aws:kms.
	ClusterSpecExternalClustersBarmanObjectStoreWalEncryption_AWS_KMS ClusterSpecExternalClustersBarmanObjectStoreWalEncryption = "AWS_KMS"
)

