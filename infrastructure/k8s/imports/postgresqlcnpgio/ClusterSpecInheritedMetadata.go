package postgresqlcnpgio


// Metadata that will be inherited by all objects related to the Cluster.
type ClusterSpecInheritedMetadata struct {
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
}

