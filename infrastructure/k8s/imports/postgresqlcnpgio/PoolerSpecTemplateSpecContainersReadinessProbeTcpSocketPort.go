package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Number or name of the port to access on the container.
//
// Number must be in the range 1 to 65535.
// Name must be an IANA_SVC_NAME.
type PoolerSpecTemplateSpecContainersReadinessProbeTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for PoolerSpecTemplateSpecContainersReadinessProbeTcpSocketPort
type jsiiProxy_PoolerSpecTemplateSpecContainersReadinessProbeTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_PoolerSpecTemplateSpecContainersReadinessProbeTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func PoolerSpecTemplateSpecContainersReadinessProbeTcpSocketPort_FromNumber(value *float64) PoolerSpecTemplateSpecContainersReadinessProbeTcpSocketPort {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecContainersReadinessProbeTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecContainersReadinessProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersReadinessProbeTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func PoolerSpecTemplateSpecContainersReadinessProbeTcpSocketPort_FromString(value *string) PoolerSpecTemplateSpecContainersReadinessProbeTcpSocketPort {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecContainersReadinessProbeTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecContainersReadinessProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersReadinessProbeTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

