package postgresqlcnpgio


// ExternalCluster represents the connection parameters to an external cluster which is used in the other sections of the configuration.
type ClusterSpecExternalClusters struct {
	// The server name, required.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The configuration for the barman-cloud tool suite.
	BarmanObjectStore *ClusterSpecExternalClustersBarmanObjectStore `field:"optional" json:"barmanObjectStore" yaml:"barmanObjectStore"`
	// The list of connection parameters, such as dbname, host, username, etc.
	ConnectionParameters *map[string]*string `field:"optional" json:"connectionParameters" yaml:"connectionParameters"`
	// The reference to the password to be used to connect to the server.
	//
	// If a password is provided, CloudNativePG creates a PostgreSQL
	// passfile at `/controller/external/NAME/pass` (where "NAME" is the
	// cluster's name). This passfile is automatically referenced in the
	// connection string when establishing a connection to the remote
	// PostgreSQL server from the current PostgreSQL `Cluster`. This ensures
	// secure and efficient password management for external clusters.
	Password *ClusterSpecExternalClustersPassword `field:"optional" json:"password" yaml:"password"`
	// The configuration of the plugin that is taking care of WAL archiving and backups for this external cluster.
	Plugin *ClusterSpecExternalClustersPlugin `field:"optional" json:"plugin" yaml:"plugin"`
	// The reference to an SSL certificate to be used to connect to this instance.
	SslCert *ClusterSpecExternalClustersSslCert `field:"optional" json:"sslCert" yaml:"sslCert"`
	// The reference to an SSL private key to be used to connect to this instance.
	SslKey *ClusterSpecExternalClustersSslKey `field:"optional" json:"sslKey" yaml:"sslKey"`
	// The reference to an SSL CA public key to be used to connect to this instance.
	SslRootCert *ClusterSpecExternalClustersSslRootCert `field:"optional" json:"sslRootCert" yaml:"sslRootCert"`
}

