package postgresqlcnpgio


// secretRef is Optional: secretRef is reference to the secret object containing sensitive information to pass to the plugin scripts.
//
// This may be
// empty if no secret object is specified. If the secret object
// contains more than one secret, all secrets are passed to the plugin
// scripts.
type PoolerSpecTemplateSpecVolumesFlexVolumeSecretRef struct {
	// Name of the referent.
	//
	// This field is effectively required, but due to backwards compatibility is
	// allowed to be empty. Instances of this type with an empty value here are
	// almost certainly wrong.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name *string `field:"optional" json:"name" yaml:"name"`
}

