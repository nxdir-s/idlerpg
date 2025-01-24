package postgresqlcnpgio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Publication is the Schema for the publications API.
type PublicationProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// PublicationSpec defines the desired state of Publication.
	Spec *PublicationSpec `field:"required" json:"spec" yaml:"spec"`
}

