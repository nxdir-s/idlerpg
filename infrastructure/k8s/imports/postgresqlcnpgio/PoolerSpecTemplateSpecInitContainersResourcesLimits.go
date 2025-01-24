package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type PoolerSpecTemplateSpecInitContainersResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for PoolerSpecTemplateSpecInitContainersResourcesLimits
type jsiiProxy_PoolerSpecTemplateSpecInitContainersResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_PoolerSpecTemplateSpecInitContainersResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func PoolerSpecTemplateSpecInitContainersResourcesLimits_FromNumber(value *float64) PoolerSpecTemplateSpecInitContainersResourcesLimits {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecInitContainersResourcesLimits_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecInitContainersResourcesLimits

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func PoolerSpecTemplateSpecInitContainersResourcesLimits_FromString(value *string) PoolerSpecTemplateSpecInitContainersResourcesLimits {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecInitContainersResourcesLimits_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecInitContainersResourcesLimits

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

