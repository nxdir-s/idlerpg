package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Specifies the output format of the exposed resources, defaults to "1".
type ClusterSpecEnvValueFromResourceFieldRefDivisor interface {
	Value() interface{}
}

// The jsii proxy struct for ClusterSpecEnvValueFromResourceFieldRefDivisor
type jsiiProxy_ClusterSpecEnvValueFromResourceFieldRefDivisor struct {
	_ byte // padding
}

func (j *jsiiProxy_ClusterSpecEnvValueFromResourceFieldRefDivisor) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ClusterSpecEnvValueFromResourceFieldRefDivisor_FromNumber(value *float64) ClusterSpecEnvValueFromResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateClusterSpecEnvValueFromResourceFieldRefDivisor_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ClusterSpecEnvValueFromResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.ClusterSpecEnvValueFromResourceFieldRefDivisor",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ClusterSpecEnvValueFromResourceFieldRefDivisor_FromString(value *string) ClusterSpecEnvValueFromResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateClusterSpecEnvValueFromResourceFieldRefDivisor_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ClusterSpecEnvValueFromResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.ClusterSpecEnvValueFromResourceFieldRefDivisor",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

