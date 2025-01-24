package postgresqlcnpgio


// EphemeralVolumesSizeLimit allows the user to set the limits for the ephemeral volumes.
type ClusterSpecEphemeralVolumesSizeLimit struct {
	// Shm is the size limit of the shared memory volume.
	Shm ClusterSpecEphemeralVolumesSizeLimitShm `field:"optional" json:"shm" yaml:"shm"`
	// TemporaryData is the size limit of the temporary data volume.
	TemporaryData ClusterSpecEphemeralVolumesSizeLimitTemporaryData `field:"optional" json:"temporaryData" yaml:"temporaryData"`
}

