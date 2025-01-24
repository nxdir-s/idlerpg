package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Name or number of the port to access on the container.
//
// Number must be in the range 1 to 65535.
// Name must be an IANA_SVC_NAME.
type PoolerSpecTemplateSpecContainersLifecyclePostStartHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for PoolerSpecTemplateSpecContainersLifecyclePostStartHttpGetPort
type jsiiProxy_PoolerSpecTemplateSpecContainersLifecyclePostStartHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_PoolerSpecTemplateSpecContainersLifecyclePostStartHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func PoolerSpecTemplateSpecContainersLifecyclePostStartHttpGetPort_FromNumber(value *float64) PoolerSpecTemplateSpecContainersLifecyclePostStartHttpGetPort {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecContainersLifecyclePostStartHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecContainersLifecyclePostStartHttpGetPort

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersLifecyclePostStartHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func PoolerSpecTemplateSpecContainersLifecyclePostStartHttpGetPort_FromString(value *string) PoolerSpecTemplateSpecContainersLifecyclePostStartHttpGetPort {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecContainersLifecyclePostStartHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecContainersLifecyclePostStartHttpGetPort

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersLifecyclePostStartHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

