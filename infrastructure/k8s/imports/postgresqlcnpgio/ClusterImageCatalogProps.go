package postgresqlcnpgio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// ClusterImageCatalog is the Schema for the clusterimagecatalogs API.
type ClusterImageCatalogProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// Specification of the desired behavior of the ClusterImageCatalog.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *ClusterImageCatalogSpec `field:"required" json:"spec" yaml:"spec"`
}

