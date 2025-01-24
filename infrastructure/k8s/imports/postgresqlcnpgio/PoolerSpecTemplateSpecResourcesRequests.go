package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type PoolerSpecTemplateSpecResourcesRequests interface {
	Value() interface{}
}

// The jsii proxy struct for PoolerSpecTemplateSpecResourcesRequests
type jsiiProxy_PoolerSpecTemplateSpecResourcesRequests struct {
	_ byte // padding
}

func (j *jsiiProxy_PoolerSpecTemplateSpecResourcesRequests) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func PoolerSpecTemplateSpecResourcesRequests_FromNumber(value *float64) PoolerSpecTemplateSpecResourcesRequests {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecResourcesRequests_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecResourcesRequests

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecResourcesRequests",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func PoolerSpecTemplateSpecResourcesRequests_FromString(value *string) PoolerSpecTemplateSpecResourcesRequests {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecResourcesRequests_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecResourcesRequests

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecResourcesRequests",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

