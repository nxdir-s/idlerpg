package postgresqlcnpgio


// flocker represents a Flocker volume attached to a kubelet's host machine.
//
// This depends on the Flocker control service being running.
// Deprecated: Flocker is deprecated and the in-tree flocker type is no longer supported.
type PoolerSpecTemplateSpecVolumesFlocker struct {
	// datasetName is Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated.
	DatasetName *string `field:"optional" json:"datasetName" yaml:"datasetName"`
	// datasetUUID is the UUID of the dataset.
	//
	// This is unique identifier of a Flocker dataset.
	DatasetUuid *string `field:"optional" json:"datasetUuid" yaml:"datasetUuid"`
}

