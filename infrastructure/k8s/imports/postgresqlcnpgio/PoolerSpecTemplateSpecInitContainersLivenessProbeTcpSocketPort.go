package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Number or name of the port to access on the container.
//
// Number must be in the range 1 to 65535.
// Name must be an IANA_SVC_NAME.
type PoolerSpecTemplateSpecInitContainersLivenessProbeTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for PoolerSpecTemplateSpecInitContainersLivenessProbeTcpSocketPort
type jsiiProxy_PoolerSpecTemplateSpecInitContainersLivenessProbeTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_PoolerSpecTemplateSpecInitContainersLivenessProbeTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func PoolerSpecTemplateSpecInitContainersLivenessProbeTcpSocketPort_FromNumber(value *float64) PoolerSpecTemplateSpecInitContainersLivenessProbeTcpSocketPort {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecInitContainersLivenessProbeTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecInitContainersLivenessProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersLivenessProbeTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func PoolerSpecTemplateSpecInitContainersLivenessProbeTcpSocketPort_FromString(value *string) PoolerSpecTemplateSpecInitContainersLivenessProbeTcpSocketPort {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecInitContainersLivenessProbeTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecInitContainersLivenessProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersLivenessProbeTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

