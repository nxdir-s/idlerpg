package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type PoolerSpecTemplateSpecResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for PoolerSpecTemplateSpecResourcesLimits
type jsiiProxy_PoolerSpecTemplateSpecResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_PoolerSpecTemplateSpecResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func PoolerSpecTemplateSpecResourcesLimits_FromNumber(value *float64) PoolerSpecTemplateSpecResourcesLimits {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecResourcesLimits_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecResourcesLimits

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func PoolerSpecTemplateSpecResourcesLimits_FromString(value *string) PoolerSpecTemplateSpecResourcesLimits {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecResourcesLimits_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecResourcesLimits

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

