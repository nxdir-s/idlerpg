package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type ClusterSpecWalStoragePvcTemplateResourcesRequests interface {
	Value() interface{}
}

// The jsii proxy struct for ClusterSpecWalStoragePvcTemplateResourcesRequests
type jsiiProxy_ClusterSpecWalStoragePvcTemplateResourcesRequests struct {
	_ byte // padding
}

func (j *jsiiProxy_ClusterSpecWalStoragePvcTemplateResourcesRequests) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ClusterSpecWalStoragePvcTemplateResourcesRequests_FromNumber(value *float64) ClusterSpecWalStoragePvcTemplateResourcesRequests {
	_init_.Initialize()

	if err := validateClusterSpecWalStoragePvcTemplateResourcesRequests_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ClusterSpecWalStoragePvcTemplateResourcesRequests

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.ClusterSpecWalStoragePvcTemplateResourcesRequests",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ClusterSpecWalStoragePvcTemplateResourcesRequests_FromString(value *string) ClusterSpecWalStoragePvcTemplateResourcesRequests {
	_init_.Initialize()

	if err := validateClusterSpecWalStoragePvcTemplateResourcesRequests_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ClusterSpecWalStoragePvcTemplateResourcesRequests

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.ClusterSpecWalStoragePvcTemplateResourcesRequests",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

