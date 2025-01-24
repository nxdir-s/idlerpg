package postgresqlcnpgio


// Compress a backup file (a tar file per tablespace) while streaming it to the object store.
//
// Available options are empty string (no
// compression, default), `gzip`, `bzip2` or `snappy`.
type ClusterSpecExternalClustersBarmanObjectStoreDataCompression string

const (
	// gzip.
	ClusterSpecExternalClustersBarmanObjectStoreDataCompression_GZIP ClusterSpecExternalClustersBarmanObjectStoreDataCompression = "GZIP"
	// bzip2.
	ClusterSpecExternalClustersBarmanObjectStoreDataCompression_BZIP2 ClusterSpecExternalClustersBarmanObjectStoreDataCompression = "BZIP2"
	// snappy.
	ClusterSpecExternalClustersBarmanObjectStoreDataCompression_SNAPPY ClusterSpecExternalClustersBarmanObjectStoreDataCompression = "SNAPPY"
)

