package postgresqlcnpgio


// The backup object containing the physical base backup from which to initiate the recovery procedure.
//
// Mutually exclusive with `source` and `volumeSnapshots`.
type ClusterSpecBootstrapRecoveryBackup struct {
	// Name of the referent.
	Name *string `field:"required" json:"name" yaml:"name"`
	// EndpointCA store the CA bundle of the barman endpoint.
	//
	// Useful when using self-signed certificates to avoid
	// errors with certificate issuer and barman-cloud-wal-archive.
	EndpointCa *ClusterSpecBootstrapRecoveryBackupEndpointCa `field:"optional" json:"endpointCa" yaml:"endpointCa"`
}

