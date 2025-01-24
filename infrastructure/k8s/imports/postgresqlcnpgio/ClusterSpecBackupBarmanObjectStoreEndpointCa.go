package postgresqlcnpgio


// EndpointCA store the CA bundle of the barman endpoint.
//
// Useful when using self-signed certificates to avoid
// errors with certificate issuer and barman-cloud-wal-archive.
type ClusterSpecBackupBarmanObjectStoreEndpointCa struct {
	// The key to select.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Name of the referent.
	Name *string `field:"required" json:"name" yaml:"name"`
}

