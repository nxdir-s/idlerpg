package postgresqlcnpgio


// The secret containing the superuser password.
//
// If not defined a new
// secret will be created with a randomly generated password.
type ClusterSpecSuperuserSecret struct {
	// Name of the referent.
	Name *string `field:"required" json:"name" yaml:"name"`
}

