package postgresqlcnpgio


// Whenever to force the encryption of files (if the bucket is not already configured for that).
//
// Allowed options are empty string (use the bucket policy, default),
// `AES256` and `aws:kms`.
type ClusterSpecBackupBarmanObjectStoreDataEncryption string

const (
	// AES256.
	ClusterSpecBackupBarmanObjectStoreDataEncryption_AES256 ClusterSpecBackupBarmanObjectStoreDataEncryption = "AES256"
	// aws:kms.
	ClusterSpecBackupBarmanObjectStoreDataEncryption_AWS_KMS ClusterSpecBackupBarmanObjectStoreDataEncryption = "AWS_KMS"
)

