package postgresqlcnpgio


// Whenever to force the encryption of files (if the bucket is not already configured for that).
//
// Allowed options are empty string (use the bucket policy, default),
// `AES256` and `aws:kms`.
type ClusterSpecBackupBarmanObjectStoreWalEncryption string

const (
	// AES256.
	ClusterSpecBackupBarmanObjectStoreWalEncryption_AES256 ClusterSpecBackupBarmanObjectStoreWalEncryption = "AES256"
	// aws:kms.
	ClusterSpecBackupBarmanObjectStoreWalEncryption_AWS_KMS ClusterSpecBackupBarmanObjectStoreWalEncryption = "AWS_KMS"
)

