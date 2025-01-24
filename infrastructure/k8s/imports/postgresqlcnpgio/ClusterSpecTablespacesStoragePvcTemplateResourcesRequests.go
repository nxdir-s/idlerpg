package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type ClusterSpecTablespacesStoragePvcTemplateResourcesRequests interface {
	Value() interface{}
}

// The jsii proxy struct for ClusterSpecTablespacesStoragePvcTemplateResourcesRequests
type jsiiProxy_ClusterSpecTablespacesStoragePvcTemplateResourcesRequests struct {
	_ byte // padding
}

func (j *jsiiProxy_ClusterSpecTablespacesStoragePvcTemplateResourcesRequests) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ClusterSpecTablespacesStoragePvcTemplateResourcesRequests_FromNumber(value *float64) ClusterSpecTablespacesStoragePvcTemplateResourcesRequests {
	_init_.Initialize()

	if err := validateClusterSpecTablespacesStoragePvcTemplateResourcesRequests_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ClusterSpecTablespacesStoragePvcTemplateResourcesRequests

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.ClusterSpecTablespacesStoragePvcTemplateResourcesRequests",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ClusterSpecTablespacesStoragePvcTemplateResourcesRequests_FromString(value *string) ClusterSpecTablespacesStoragePvcTemplateResourcesRequests {
	_init_.Initialize()

	if err := validateClusterSpecTablespacesStoragePvcTemplateResourcesRequests_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ClusterSpecTablespacesStoragePvcTemplateResourcesRequests

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.ClusterSpecTablespacesStoragePvcTemplateResourcesRequests",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

