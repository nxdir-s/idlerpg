package postgresqlcnpgio


// The configuration to be used to backup the data files When not defined, base backups files will be stored uncompressed and may be unencrypted in the object store, according to the bucket default policy.
type ClusterSpecExternalClustersBarmanObjectStoreData struct {
	// AdditionalCommandArgs represents additional arguments that can be appended to the 'barman-cloud-backup' command-line invocation.
	//
	// These arguments
	// provide flexibility to customize the backup process further according to
	// specific requirements or configurations.
	//
	// Example:
	// In a scenario where specialized backup options are required, such as setting
	// a specific timeout or defining custom behavior, users can use this field
	// to specify additional command arguments.
	//
	// Note:
	// It's essential to ensure that the provided arguments are valid and supported
	// by the 'barman-cloud-backup' command, to avoid potential errors or unintended
	// behavior during execution.
	AdditionalCommandArgs *[]*string `field:"optional" json:"additionalCommandArgs" yaml:"additionalCommandArgs"`
	// Compress a backup file (a tar file per tablespace) while streaming it to the object store.
	//
	// Available options are empty string (no
	// compression, default), `gzip`, `bzip2` or `snappy`.
	Compression ClusterSpecExternalClustersBarmanObjectStoreDataCompression `field:"optional" json:"compression" yaml:"compression"`
	// Whenever to force the encryption of files (if the bucket is not already configured for that).
	//
	// Allowed options are empty string (use the bucket policy, default),
	// `AES256` and `aws:kms`.
	Encryption ClusterSpecExternalClustersBarmanObjectStoreDataEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// Control whether the I/O workload for the backup initial checkpoint will be limited, according to the `checkpoint_completion_target` setting on the PostgreSQL server.
	//
	// If set to true, an immediate checkpoint will be
	// used, meaning PostgreSQL will complete the checkpoint as soon as
	// possible. `false` by default.
	ImmediateCheckpoint *bool `field:"optional" json:"immediateCheckpoint" yaml:"immediateCheckpoint"`
	// The number of parallel jobs to be used to upload the backup, defaults to 2.
	Jobs *float64 `field:"optional" json:"jobs" yaml:"jobs"`
}

