package postgresqlcnpgio


// The reference to the password to be used to connect to the server.
//
// If a password is provided, CloudNativePG creates a PostgreSQL
// passfile at `/controller/external/NAME/pass` (where "NAME" is the
// cluster's name). This passfile is automatically referenced in the
// connection string when establishing a connection to the remote
// PostgreSQL server from the current PostgreSQL `Cluster`. This ensures
// secure and efficient password management for external clusters.
type ClusterSpecExternalClustersPassword struct {
	// The key of the secret to select from.
	//
	// Must be a valid secret key.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Name of the referent.
	//
	// This field is effectively required, but due to backwards compatibility is
	// allowed to be empty. Instances of this type with an empty value here are
	// almost certainly wrong.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specify whether the Secret or its key must be defined.
	Optional *bool `field:"optional" json:"optional" yaml:"optional"`
}

