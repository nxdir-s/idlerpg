package postgresqlcnpgio


// Configure the generation of the service account.
type ClusterSpecServiceAccountTemplate struct {
	// Metadata are the metadata to be used for the generated service account.
	Metadata *ClusterSpecServiceAccountTemplateMetadata `field:"required" json:"metadata" yaml:"metadata"`
}

