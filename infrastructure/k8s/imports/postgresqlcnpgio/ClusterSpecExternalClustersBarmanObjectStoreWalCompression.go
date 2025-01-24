package postgresqlcnpgio


// Compress a WAL file before sending it to the object store.
//
// Available
// options are empty string (no compression, default), `gzip`, `bzip2` or `snappy`.
type ClusterSpecExternalClustersBarmanObjectStoreWalCompression string

const (
	// gzip.
	ClusterSpecExternalClustersBarmanObjectStoreWalCompression_GZIP ClusterSpecExternalClustersBarmanObjectStoreWalCompression = "GZIP"
	// bzip2.
	ClusterSpecExternalClustersBarmanObjectStoreWalCompression_BZIP2 ClusterSpecExternalClustersBarmanObjectStoreWalCompression = "BZIP2"
	// snappy.
	ClusterSpecExternalClustersBarmanObjectStoreWalCompression_SNAPPY ClusterSpecExternalClustersBarmanObjectStoreWalCompression = "SNAPPY"
)

