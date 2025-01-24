package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type PoolerSpecTemplateSpecContainersResourcesRequests interface {
	Value() interface{}
}

// The jsii proxy struct for PoolerSpecTemplateSpecContainersResourcesRequests
type jsiiProxy_PoolerSpecTemplateSpecContainersResourcesRequests struct {
	_ byte // padding
}

func (j *jsiiProxy_PoolerSpecTemplateSpecContainersResourcesRequests) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func PoolerSpecTemplateSpecContainersResourcesRequests_FromNumber(value *float64) PoolerSpecTemplateSpecContainersResourcesRequests {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecContainersResourcesRequests_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecContainersResourcesRequests

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersResourcesRequests",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func PoolerSpecTemplateSpecContainersResourcesRequests_FromString(value *string) PoolerSpecTemplateSpecContainersResourcesRequests {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecContainersResourcesRequests_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecContainersResourcesRequests

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersResourcesRequests",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

