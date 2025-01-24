package postgresqlcnpgio


// Compress a WAL file before sending it to the object store.
//
// Available
// options are empty string (no compression, default), `gzip`, `bzip2` or `snappy`.
type ClusterSpecBackupBarmanObjectStoreWalCompression string

const (
	// gzip.
	ClusterSpecBackupBarmanObjectStoreWalCompression_GZIP ClusterSpecBackupBarmanObjectStoreWalCompression = "GZIP"
	// bzip2.
	ClusterSpecBackupBarmanObjectStoreWalCompression_BZIP2 ClusterSpecBackupBarmanObjectStoreWalCompression = "BZIP2"
	// snappy.
	ClusterSpecBackupBarmanObjectStoreWalCompression_SNAPPY ClusterSpecBackupBarmanObjectStoreWalCompression = "SNAPPY"
)

