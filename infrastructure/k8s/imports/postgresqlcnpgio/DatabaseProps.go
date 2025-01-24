package postgresqlcnpgio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Database is the Schema for the databases API.
type DatabaseProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// Specification of the desired Database.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *DatabaseSpec `field:"required" json:"spec" yaml:"spec"`
}

