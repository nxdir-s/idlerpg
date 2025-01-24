package postgresqlcnpgio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// ImageCatalog is the Schema for the imagecatalogs API.
type ImageCatalogProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// Specification of the desired behavior of the ImageCatalog.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *ImageCatalogSpec `field:"required" json:"spec" yaml:"spec"`
}

