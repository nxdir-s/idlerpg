package postgresqlcnpgio


// Metadata are the metadata to be used for the generated service account.
type ClusterSpecServiceAccountTemplateMetadata struct {
	// Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.
	//
	// They are not
	// queryable and should be preserved when modifying objects.
	// More info: http://kubernetes.io/docs/user-guide/annotations
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Map of string keys and values that can be used to organize and categorize (scope and select) objects.
	//
	// May match selectors of replication controllers
	// and services.
	// More info: http://kubernetes.io/docs/user-guide/labels
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The name of the resource.
	//
	// Only supported for certain types.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

