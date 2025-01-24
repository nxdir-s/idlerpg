package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Specifies the output format of the exposed resources, defaults to "1".
type PoolerSpecTemplateSpecVolumesDownwardApiItemsResourceFieldRefDivisor interface {
	Value() interface{}
}

// The jsii proxy struct for PoolerSpecTemplateSpecVolumesDownwardApiItemsResourceFieldRefDivisor
type jsiiProxy_PoolerSpecTemplateSpecVolumesDownwardApiItemsResourceFieldRefDivisor struct {
	_ byte // padding
}

func (j *jsiiProxy_PoolerSpecTemplateSpecVolumesDownwardApiItemsResourceFieldRefDivisor) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func PoolerSpecTemplateSpecVolumesDownwardApiItemsResourceFieldRefDivisor_FromNumber(value *float64) PoolerSpecTemplateSpecVolumesDownwardApiItemsResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecVolumesDownwardApiItemsResourceFieldRefDivisor_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecVolumesDownwardApiItemsResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesDownwardApiItemsResourceFieldRefDivisor",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func PoolerSpecTemplateSpecVolumesDownwardApiItemsResourceFieldRefDivisor_FromString(value *string) PoolerSpecTemplateSpecVolumesDownwardApiItemsResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecVolumesDownwardApiItemsResourceFieldRefDivisor_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecVolumesDownwardApiItemsResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesDownwardApiItemsResourceFieldRefDivisor",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

