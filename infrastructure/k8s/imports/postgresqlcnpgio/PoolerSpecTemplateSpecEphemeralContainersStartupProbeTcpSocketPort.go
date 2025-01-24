package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Number or name of the port to access on the container.
//
// Number must be in the range 1 to 65535.
// Name must be an IANA_SVC_NAME.
type PoolerSpecTemplateSpecEphemeralContainersStartupProbeTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for PoolerSpecTemplateSpecEphemeralContainersStartupProbeTcpSocketPort
type jsiiProxy_PoolerSpecTemplateSpecEphemeralContainersStartupProbeTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_PoolerSpecTemplateSpecEphemeralContainersStartupProbeTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func PoolerSpecTemplateSpecEphemeralContainersStartupProbeTcpSocketPort_FromNumber(value *float64) PoolerSpecTemplateSpecEphemeralContainersStartupProbeTcpSocketPort {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecEphemeralContainersStartupProbeTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecEphemeralContainersStartupProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersStartupProbeTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func PoolerSpecTemplateSpecEphemeralContainersStartupProbeTcpSocketPort_FromString(value *string) PoolerSpecTemplateSpecEphemeralContainersStartupProbeTcpSocketPort {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecEphemeralContainersStartupProbeTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecEphemeralContainersStartupProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersStartupProbeTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

