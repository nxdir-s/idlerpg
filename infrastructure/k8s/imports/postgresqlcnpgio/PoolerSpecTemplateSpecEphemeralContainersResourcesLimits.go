package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type PoolerSpecTemplateSpecEphemeralContainersResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for PoolerSpecTemplateSpecEphemeralContainersResourcesLimits
type jsiiProxy_PoolerSpecTemplateSpecEphemeralContainersResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_PoolerSpecTemplateSpecEphemeralContainersResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func PoolerSpecTemplateSpecEphemeralContainersResourcesLimits_FromNumber(value *float64) PoolerSpecTemplateSpecEphemeralContainersResourcesLimits {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecEphemeralContainersResourcesLimits_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecEphemeralContainersResourcesLimits

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func PoolerSpecTemplateSpecEphemeralContainersResourcesLimits_FromString(value *string) PoolerSpecTemplateSpecEphemeralContainersResourcesLimits {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecEphemeralContainersResourcesLimits_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecEphemeralContainersResourcesLimits

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

