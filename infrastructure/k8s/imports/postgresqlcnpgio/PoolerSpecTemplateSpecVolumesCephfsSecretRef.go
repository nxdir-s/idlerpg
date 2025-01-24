package postgresqlcnpgio


// secretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty.
//
// More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
type PoolerSpecTemplateSpecVolumesCephfsSecretRef struct {
	// Name of the referent.
	//
	// This field is effectively required, but due to backwards compatibility is
	// allowed to be empty. Instances of this type with an empty value here are
	// almost certainly wrong.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name *string `field:"optional" json:"name" yaml:"name"`
}

