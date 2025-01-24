package postgresqlcnpgio


// awsElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod.
//
// Deprecated: AWSElasticBlockStore is deprecated. All operations for the in-tree
// awsElasticBlockStore type are redirected to the ebs.csi.aws.com CSI driver.
// More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
type PoolerSpecTemplateSpecVolumesAwsElasticBlockStore struct {
	// volumeID is unique ID of the persistent disk resource in AWS (Amazon EBS volume).
	//
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
	// fsType is the filesystem type of the volume that you want to mount.
	//
	// Tip: Ensure that the filesystem type is supported by the host operating system.
	// Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// partition is the partition in the volume that you want to mount.
	//
	// If omitted, the default is to mount by volume name.
	// Examples: For volume /dev/sda1, you specify the partition as "1".
	// Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty).
	Partition *float64 `field:"optional" json:"partition" yaml:"partition"`
	// readOnly value true will force the readOnly setting in VolumeMounts.
	//
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}

