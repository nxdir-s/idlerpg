package postgresqlcnpgio


// The configuration for the backup of the WAL stream.
//
// When not defined, WAL files will be stored uncompressed and may be
// unencrypted in the object store, according to the bucket default policy.
type ClusterSpecExternalClustersBarmanObjectStoreWal struct {
	// Additional arguments that can be appended to the 'barman-cloud-wal-archive' command-line invocation.
	//
	// These arguments provide flexibility to customize
	// the WAL archive process further, according to specific requirements or configurations.
	//
	// Example:
	// In a scenario where specialized backup options are required, such as setting
	// a specific timeout or defining custom behavior, users can use this field
	// to specify additional command arguments.
	//
	// Note:
	// It's essential to ensure that the provided arguments are valid and supported
	// by the 'barman-cloud-wal-archive' command, to avoid potential errors or unintended
	// behavior during execution.
	ArchiveAdditionalCommandArgs *[]*string `field:"optional" json:"archiveAdditionalCommandArgs" yaml:"archiveAdditionalCommandArgs"`
	// Compress a WAL file before sending it to the object store.
	//
	// Available
	// options are empty string (no compression, default), `gzip`, `bzip2` or `snappy`.
	Compression ClusterSpecExternalClustersBarmanObjectStoreWalCompression `field:"optional" json:"compression" yaml:"compression"`
	// Whenever to force the encryption of files (if the bucket is not already configured for that).
	//
	// Allowed options are empty string (use the bucket policy, default),
	// `AES256` and `aws:kms`.
	Encryption ClusterSpecExternalClustersBarmanObjectStoreWalEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// Number of WAL files to be either archived in parallel (when the PostgreSQL instance is archiving to a backup object store) or restored in parallel (when a PostgreSQL standby is fetching WAL files from a recovery object store).
	//
	// If not specified, WAL files
	// will be processed one at a time. It accepts a positive integer as a
	// value - with 1 being the minimum accepted value.
	MaxParallel *float64 `field:"optional" json:"maxParallel" yaml:"maxParallel"`
	// Additional arguments that can be appended to the 'barman-cloud-wal-restore' command-line invocation.
	//
	// These arguments provide flexibility to customize
	// the WAL restore process further, according to specific requirements or configurations.
	//
	// Example:
	// In a scenario where specialized backup options are required, such as setting
	// a specific timeout or defining custom behavior, users can use this field
	// to specify additional command arguments.
	//
	// Note:
	// It's essential to ensure that the provided arguments are valid and supported
	// by the 'barman-cloud-wal-restore' command, to avoid potential errors or unintended
	// behavior during execution.
	RestoreAdditionalCommandArgs *[]*string `field:"optional" json:"restoreAdditionalCommandArgs" yaml:"restoreAdditionalCommandArgs"`
}

