package postgresqlcnpgio


// Name of the secret containing the initial credentials for the owner of the user database.
//
// If empty a new secret will be
// created from scratch.
type ClusterSpecBootstrapPgBasebackupSecret struct {
	// Name of the referent.
	Name *string `field:"required" json:"name" yaml:"name"`
}

