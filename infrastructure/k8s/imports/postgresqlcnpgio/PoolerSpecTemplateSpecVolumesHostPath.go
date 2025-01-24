package postgresqlcnpgio


// hostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container.
//
// This is generally
// used for system agents or other privileged things that are allowed
// to see the host machine. Most containers will NOT need this.
// More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
type PoolerSpecTemplateSpecVolumesHostPath struct {
	// path of the directory on the host.
	//
	// If the path is a symlink, it will follow the link to the real path.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
	Path *string `field:"required" json:"path" yaml:"path"`
	// type for HostPath Volume Defaults to "" More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath.
	// Default: More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

