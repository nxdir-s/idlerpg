package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Number or name of the port to access on the container.
//
// Number must be in the range 1 to 65535.
// Name must be an IANA_SVC_NAME.
type PoolerSpecTemplateSpecEphemeralContainersReadinessProbeTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for PoolerSpecTemplateSpecEphemeralContainersReadinessProbeTcpSocketPort
type jsiiProxy_PoolerSpecTemplateSpecEphemeralContainersReadinessProbeTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_PoolerSpecTemplateSpecEphemeralContainersReadinessProbeTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func PoolerSpecTemplateSpecEphemeralContainersReadinessProbeTcpSocketPort_FromNumber(value *float64) PoolerSpecTemplateSpecEphemeralContainersReadinessProbeTcpSocketPort {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecEphemeralContainersReadinessProbeTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecEphemeralContainersReadinessProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersReadinessProbeTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func PoolerSpecTemplateSpecEphemeralContainersReadinessProbeTcpSocketPort_FromString(value *string) PoolerSpecTemplateSpecEphemeralContainersReadinessProbeTcpSocketPort {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecEphemeralContainersReadinessProbeTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecEphemeralContainersReadinessProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersReadinessProbeTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

