package postgresqlcnpgio


// Secret containing the password of the role (if present) If null, the password will be ignored unless DisablePassword is set.
type ClusterSpecManagedRolesPasswordSecret struct {
	// Name of the referent.
	Name *string `field:"required" json:"name" yaml:"name"`
}

