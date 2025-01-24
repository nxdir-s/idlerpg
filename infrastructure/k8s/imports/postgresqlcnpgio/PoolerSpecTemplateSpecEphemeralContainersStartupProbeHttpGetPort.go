package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Name or number of the port to access on the container.
//
// Number must be in the range 1 to 65535.
// Name must be an IANA_SVC_NAME.
type PoolerSpecTemplateSpecEphemeralContainersStartupProbeHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for PoolerSpecTemplateSpecEphemeralContainersStartupProbeHttpGetPort
type jsiiProxy_PoolerSpecTemplateSpecEphemeralContainersStartupProbeHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_PoolerSpecTemplateSpecEphemeralContainersStartupProbeHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func PoolerSpecTemplateSpecEphemeralContainersStartupProbeHttpGetPort_FromNumber(value *float64) PoolerSpecTemplateSpecEphemeralContainersStartupProbeHttpGetPort {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecEphemeralContainersStartupProbeHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecEphemeralContainersStartupProbeHttpGetPort

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersStartupProbeHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func PoolerSpecTemplateSpecEphemeralContainersStartupProbeHttpGetPort_FromString(value *string) PoolerSpecTemplateSpecEphemeralContainersStartupProbeHttpGetPort {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecEphemeralContainersStartupProbeHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecEphemeralContainersStartupProbeHttpGetPort

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersStartupProbeHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

