package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type PoolerSpecTemplateSpecContainersResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for PoolerSpecTemplateSpecContainersResourcesLimits
type jsiiProxy_PoolerSpecTemplateSpecContainersResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_PoolerSpecTemplateSpecContainersResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func PoolerSpecTemplateSpecContainersResourcesLimits_FromNumber(value *float64) PoolerSpecTemplateSpecContainersResourcesLimits {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecContainersResourcesLimits_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecContainersResourcesLimits

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func PoolerSpecTemplateSpecContainersResourcesLimits_FromString(value *string) PoolerSpecTemplateSpecContainersResourcesLimits {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecContainersResourcesLimits_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecContainersResourcesLimits

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

