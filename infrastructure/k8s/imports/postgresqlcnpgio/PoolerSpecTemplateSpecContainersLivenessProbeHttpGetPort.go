package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Name or number of the port to access on the container.
//
// Number must be in the range 1 to 65535.
// Name must be an IANA_SVC_NAME.
type PoolerSpecTemplateSpecContainersLivenessProbeHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for PoolerSpecTemplateSpecContainersLivenessProbeHttpGetPort
type jsiiProxy_PoolerSpecTemplateSpecContainersLivenessProbeHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_PoolerSpecTemplateSpecContainersLivenessProbeHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func PoolerSpecTemplateSpecContainersLivenessProbeHttpGetPort_FromNumber(value *float64) PoolerSpecTemplateSpecContainersLivenessProbeHttpGetPort {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecContainersLivenessProbeHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecContainersLivenessProbeHttpGetPort

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersLivenessProbeHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func PoolerSpecTemplateSpecContainersLivenessProbeHttpGetPort_FromString(value *string) PoolerSpecTemplateSpecContainersLivenessProbeHttpGetPort {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecContainersLivenessProbeHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecContainersLivenessProbeHttpGetPort

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersLivenessProbeHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

