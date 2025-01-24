package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Specifies the output format of the exposed resources, defaults to "1".
type PoolerSpecTemplateSpecEphemeralContainersEnvValueFromResourceFieldRefDivisor interface {
	Value() interface{}
}

// The jsii proxy struct for PoolerSpecTemplateSpecEphemeralContainersEnvValueFromResourceFieldRefDivisor
type jsiiProxy_PoolerSpecTemplateSpecEphemeralContainersEnvValueFromResourceFieldRefDivisor struct {
	_ byte // padding
}

func (j *jsiiProxy_PoolerSpecTemplateSpecEphemeralContainersEnvValueFromResourceFieldRefDivisor) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func PoolerSpecTemplateSpecEphemeralContainersEnvValueFromResourceFieldRefDivisor_FromNumber(value *float64) PoolerSpecTemplateSpecEphemeralContainersEnvValueFromResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecEphemeralContainersEnvValueFromResourceFieldRefDivisor_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecEphemeralContainersEnvValueFromResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersEnvValueFromResourceFieldRefDivisor",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func PoolerSpecTemplateSpecEphemeralContainersEnvValueFromResourceFieldRefDivisor_FromString(value *string) PoolerSpecTemplateSpecEphemeralContainersEnvValueFromResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecEphemeralContainersEnvValueFromResourceFieldRefDivisor_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecEphemeralContainersEnvValueFromResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersEnvValueFromResourceFieldRefDivisor",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

