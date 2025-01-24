package postgresqlcnpgio


// secretRef is name of the authentication secret for RBDUser.
//
// If provided
// overrides keyring.
// Default is nil.
// More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
// Default: nil.
//
type PoolerSpecTemplateSpecVolumesRbdSecretRef struct {
	// Name of the referent.
	//
	// This field is effectively required, but due to backwards compatibility is
	// allowed to be empty. Instances of this type with an empty value here are
	// almost certainly wrong.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name *string `field:"optional" json:"name" yaml:"name"`
}

