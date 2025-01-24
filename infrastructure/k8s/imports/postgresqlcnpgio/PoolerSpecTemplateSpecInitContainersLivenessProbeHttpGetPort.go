package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Name or number of the port to access on the container.
//
// Number must be in the range 1 to 65535.
// Name must be an IANA_SVC_NAME.
type PoolerSpecTemplateSpecInitContainersLivenessProbeHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for PoolerSpecTemplateSpecInitContainersLivenessProbeHttpGetPort
type jsiiProxy_PoolerSpecTemplateSpecInitContainersLivenessProbeHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_PoolerSpecTemplateSpecInitContainersLivenessProbeHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func PoolerSpecTemplateSpecInitContainersLivenessProbeHttpGetPort_FromNumber(value *float64) PoolerSpecTemplateSpecInitContainersLivenessProbeHttpGetPort {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecInitContainersLivenessProbeHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecInitContainersLivenessProbeHttpGetPort

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersLivenessProbeHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func PoolerSpecTemplateSpecInitContainersLivenessProbeHttpGetPort_FromString(value *string) PoolerSpecTemplateSpecInitContainersLivenessProbeHttpGetPort {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecInitContainersLivenessProbeHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecInitContainersLivenessProbeHttpGetPort

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersLivenessProbeHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

