package postgresqlcnpgio


// Compress a backup file (a tar file per tablespace) while streaming it to the object store.
//
// Available options are empty string (no
// compression, default), `gzip`, `bzip2` or `snappy`.
type ClusterSpecBackupBarmanObjectStoreDataCompression string

const (
	// gzip.
	ClusterSpecBackupBarmanObjectStoreDataCompression_GZIP ClusterSpecBackupBarmanObjectStoreDataCompression = "GZIP"
	// bzip2.
	ClusterSpecBackupBarmanObjectStoreDataCompression_BZIP2 ClusterSpecBackupBarmanObjectStoreDataCompression = "BZIP2"
	// snappy.
	ClusterSpecBackupBarmanObjectStoreDataCompression_SNAPPY ClusterSpecBackupBarmanObjectStoreDataCompression = "SNAPPY"
)

