package postgresqlcnpgio


// The capabilities to add/drop when running containers.
//
// Defaults to the default set of capabilities granted by the container runtime.
// Note that this field cannot be set when spec.os.name is windows.
// Default: the default set of capabilities granted by the container runtime.
//
type PoolerSpecTemplateSpecEphemeralContainersSecurityContextCapabilities struct {
	// Added capabilities.
	Add *[]*string `field:"optional" json:"add" yaml:"add"`
	// Removed capabilities.
	Drop *[]*string `field:"optional" json:"drop" yaml:"drop"`
}

