package postgresqlcnpgio


// The configuration for the CA and related certificates.
type ClusterSpecCertificates struct {
	// The secret containing the Client CA certificate.
	//
	// If not defined, a new secret will be created
	// with a self-signed CA and will be used to generate all the client certificates.<br />
	// <br />
	// Contains:<br />
	// <br />
	// - `ca.crt`: CA that should be used to validate the client certificates,
	// used as `ssl_ca_file` of all the instances.<br />
	// - `ca.key`: key used to generate client certificates, if ReplicationTLSSecret is provided,
	// this can be omitted.<br />
	ClientCaSecret *string `field:"optional" json:"clientCaSecret" yaml:"clientCaSecret"`
	// The secret of type kubernetes.io/tls containing the client certificate to authenticate as the `streaming_replica` user. If not defined, ClientCASecret must provide also `ca.key`, and a new secret will be created using the provided CA.
	ReplicationTlsSecret *string `field:"optional" json:"replicationTlsSecret" yaml:"replicationTlsSecret"`
	// The list of the server alternative DNS names to be added to the generated server TLS certificates, when required.
	ServerAltDnsNames *[]*string `field:"optional" json:"serverAltDnsNames" yaml:"serverAltDnsNames"`
	// The secret containing the Server CA certificate.
	//
	// If not defined, a new secret will be created
	// with a self-signed CA and will be used to generate the TLS certificate ServerTLSSecret.<br />
	// <br />
	// Contains:<br />
	// <br />
	// - `ca.crt`: CA that should be used to validate the server certificate,
	// used as `sslrootcert` in client connection strings.<br />
	// - `ca.key`: key used to generate Server SSL certs, if ServerTLSSecret is provided,
	// this can be omitted.<br />
	ServerCaSecret *string `field:"optional" json:"serverCaSecret" yaml:"serverCaSecret"`
	// The secret of type kubernetes.io/tls containing the server TLS certificate and key that will be set as `ssl_cert_file` and `ssl_key_file` so that clients can connect to postgres securely. If not defined, ServerCASecret must provide also `ca.key` and a new secret will be created using the provided CA.
	ServerTlsSecret *string `field:"optional" json:"serverTlsSecret" yaml:"serverTlsSecret"`
}

