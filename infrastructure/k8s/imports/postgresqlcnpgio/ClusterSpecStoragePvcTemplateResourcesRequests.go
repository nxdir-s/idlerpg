package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type ClusterSpecStoragePvcTemplateResourcesRequests interface {
	Value() interface{}
}

// The jsii proxy struct for ClusterSpecStoragePvcTemplateResourcesRequests
type jsiiProxy_ClusterSpecStoragePvcTemplateResourcesRequests struct {
	_ byte // padding
}

func (j *jsiiProxy_ClusterSpecStoragePvcTemplateResourcesRequests) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ClusterSpecStoragePvcTemplateResourcesRequests_FromNumber(value *float64) ClusterSpecStoragePvcTemplateResourcesRequests {
	_init_.Initialize()

	if err := validateClusterSpecStoragePvcTemplateResourcesRequests_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ClusterSpecStoragePvcTemplateResourcesRequests

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.ClusterSpecStoragePvcTemplateResourcesRequests",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ClusterSpecStoragePvcTemplateResourcesRequests_FromString(value *string) ClusterSpecStoragePvcTemplateResourcesRequests {
	_init_.Initialize()

	if err := validateClusterSpecStoragePvcTemplateResourcesRequests_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ClusterSpecStoragePvcTemplateResourcesRequests

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.ClusterSpecStoragePvcTemplateResourcesRequests",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

