package postgresqlcnpgio


// nodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls.
//
// This field is optional, and  may be empty if no secret is required. If the
// secret object contains more than one secret, all secret references are passed.
type PoolerSpecTemplateSpecVolumesCsiNodePublishSecretRef struct {
	// Name of the referent.
	//
	// This field is effectively required, but due to backwards compatibility is
	// allowed to be empty. Instances of this type with an empty value here are
	// almost certainly wrong.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name *string `field:"optional" json:"name" yaml:"name"`
}

