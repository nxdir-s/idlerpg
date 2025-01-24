package postgresqlcnpgio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// ScheduledBackup is the Schema for the scheduledbackups API.
type ScheduledBackupProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// Specification of the desired behavior of the ScheduledBackup.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *ScheduledBackupSpec `field:"required" json:"spec" yaml:"spec"`
}

