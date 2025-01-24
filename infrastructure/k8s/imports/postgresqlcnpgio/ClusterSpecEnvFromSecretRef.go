package postgresqlcnpgio


// The Secret to select from.
type ClusterSpecEnvFromSecretRef struct {
	// Name of the referent.
	//
	// This field is effectively required, but due to backwards compatibility is
	// allowed to be empty. Instances of this type with an empty value here are
	// almost certainly wrong.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specify whether the Secret must be defined.
	Optional *bool `field:"optional" json:"optional" yaml:"optional"`
}

