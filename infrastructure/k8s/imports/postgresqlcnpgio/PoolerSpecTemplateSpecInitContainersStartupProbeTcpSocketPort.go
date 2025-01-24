package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Number or name of the port to access on the container.
//
// Number must be in the range 1 to 65535.
// Name must be an IANA_SVC_NAME.
type PoolerSpecTemplateSpecInitContainersStartupProbeTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for PoolerSpecTemplateSpecInitContainersStartupProbeTcpSocketPort
type jsiiProxy_PoolerSpecTemplateSpecInitContainersStartupProbeTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_PoolerSpecTemplateSpecInitContainersStartupProbeTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func PoolerSpecTemplateSpecInitContainersStartupProbeTcpSocketPort_FromNumber(value *float64) PoolerSpecTemplateSpecInitContainersStartupProbeTcpSocketPort {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecInitContainersStartupProbeTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecInitContainersStartupProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersStartupProbeTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func PoolerSpecTemplateSpecInitContainersStartupProbeTcpSocketPort_FromString(value *string) PoolerSpecTemplateSpecInitContainersStartupProbeTcpSocketPort {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecInitContainersStartupProbeTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecInitContainersStartupProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersStartupProbeTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

