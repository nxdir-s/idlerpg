package postgresqlcnpgio


// Defines the major PostgreSQL version we want to use within an ImageCatalog.
type ClusterSpecImageCatalogRef struct {
	// Kind is the type of resource being referenced.
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// The major version of PostgreSQL we want to use from the ImageCatalog.
	Major *float64 `field:"required" json:"major" yaml:"major"`
	// Name is the name of resource being referenced.
	Name *string `field:"required" json:"name" yaml:"name"`
	// APIGroup is the group for the resource being referenced.
	//
	// If APIGroup is not specified, the specified Kind must be in the core API group.
	// For any other third-party types, APIGroup is required.
	ApiGroup *string `field:"optional" json:"apiGroup" yaml:"apiGroup"`
}

