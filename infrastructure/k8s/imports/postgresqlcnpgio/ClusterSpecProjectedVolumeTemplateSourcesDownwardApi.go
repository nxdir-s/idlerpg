package postgresqlcnpgio


// downwardAPI information about the downwardAPI data to project.
type ClusterSpecProjectedVolumeTemplateSourcesDownwardApi struct {
	// Items is a list of DownwardAPIVolume file.
	Items *[]*ClusterSpecProjectedVolumeTemplateSourcesDownwardApiItems `field:"optional" json:"items" yaml:"items"`
}

