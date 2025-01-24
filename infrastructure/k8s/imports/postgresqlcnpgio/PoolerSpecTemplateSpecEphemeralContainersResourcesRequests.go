package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type PoolerSpecTemplateSpecEphemeralContainersResourcesRequests interface {
	Value() interface{}
}

// The jsii proxy struct for PoolerSpecTemplateSpecEphemeralContainersResourcesRequests
type jsiiProxy_PoolerSpecTemplateSpecEphemeralContainersResourcesRequests struct {
	_ byte // padding
}

func (j *jsiiProxy_PoolerSpecTemplateSpecEphemeralContainersResourcesRequests) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func PoolerSpecTemplateSpecEphemeralContainersResourcesRequests_FromNumber(value *float64) PoolerSpecTemplateSpecEphemeralContainersResourcesRequests {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecEphemeralContainersResourcesRequests_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecEphemeralContainersResourcesRequests

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersResourcesRequests",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func PoolerSpecTemplateSpecEphemeralContainersResourcesRequests_FromString(value *string) PoolerSpecTemplateSpecEphemeralContainersResourcesRequests {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecEphemeralContainersResourcesRequests_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecEphemeralContainersResourcesRequests

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersResourcesRequests",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

