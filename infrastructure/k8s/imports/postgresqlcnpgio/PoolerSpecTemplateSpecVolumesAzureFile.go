package postgresqlcnpgio


// azureFile represents an Azure File Service mount on the host and bind mount to the pod.
//
// Deprecated: AzureFile is deprecated. All operations for the in-tree azureFile type
// are redirected to the file.csi.azure.com CSI driver.
type PoolerSpecTemplateSpecVolumesAzureFile struct {
	// secretName is the  name of secret that contains Azure Storage Account Name and Key.
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
	// shareName is the azure share Name.
	ShareName *string `field:"required" json:"shareName" yaml:"shareName"`
	// readOnly defaults to false (read/write).
	//
	// ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}

