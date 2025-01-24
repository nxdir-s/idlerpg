package postgresqlcnpgio


// nfs represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs.
type PoolerSpecTemplateSpecVolumesNfs struct {
	// path that is exported by the NFS server.
	//
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
	Path *string `field:"required" json:"path" yaml:"path"`
	// server is the hostname or IP address of the NFS server.
	//
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
	Server *string `field:"required" json:"server" yaml:"server"`
	// readOnly here will force the NFS export to be mounted with read-only permissions.
	//
	// Defaults to false.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
	// Default: false.
	//
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}

