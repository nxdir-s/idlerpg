package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Specifies the output format of the exposed resources, defaults to "1".
type PoolerSpecTemplateSpecContainersEnvValueFromResourceFieldRefDivisor interface {
	Value() interface{}
}

// The jsii proxy struct for PoolerSpecTemplateSpecContainersEnvValueFromResourceFieldRefDivisor
type jsiiProxy_PoolerSpecTemplateSpecContainersEnvValueFromResourceFieldRefDivisor struct {
	_ byte // padding
}

func (j *jsiiProxy_PoolerSpecTemplateSpecContainersEnvValueFromResourceFieldRefDivisor) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func PoolerSpecTemplateSpecContainersEnvValueFromResourceFieldRefDivisor_FromNumber(value *float64) PoolerSpecTemplateSpecContainersEnvValueFromResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecContainersEnvValueFromResourceFieldRefDivisor_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecContainersEnvValueFromResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersEnvValueFromResourceFieldRefDivisor",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func PoolerSpecTemplateSpecContainersEnvValueFromResourceFieldRefDivisor_FromString(value *string) PoolerSpecTemplateSpecContainersEnvValueFromResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecContainersEnvValueFromResourceFieldRefDivisor_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecContainersEnvValueFromResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersEnvValueFromResourceFieldRefDivisor",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

