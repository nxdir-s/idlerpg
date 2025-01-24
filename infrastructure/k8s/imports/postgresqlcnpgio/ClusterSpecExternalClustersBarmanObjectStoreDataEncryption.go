package postgresqlcnpgio


// Whenever to force the encryption of files (if the bucket is not already configured for that).
//
// Allowed options are empty string (use the bucket policy, default),
// `AES256` and `aws:kms`.
type ClusterSpecExternalClustersBarmanObjectStoreDataEncryption string

const (
	// AES256.
	ClusterSpecExternalClustersBarmanObjectStoreDataEncryption_AES256 ClusterSpecExternalClustersBarmanObjectStoreDataEncryption = "AES256"
	// aws:kms.
	ClusterSpecExternalClustersBarmanObjectStoreDataEncryption_AWS_KMS ClusterSpecExternalClustersBarmanObjectStoreDataEncryption = "AWS_KMS"
)

