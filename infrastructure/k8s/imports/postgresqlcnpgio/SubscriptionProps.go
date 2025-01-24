package postgresqlcnpgio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Subscription is the Schema for the subscriptions API.
type SubscriptionProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// SubscriptionSpec defines the desired state of Subscription.
	Spec *SubscriptionSpec `field:"required" json:"spec" yaml:"spec"`
}

