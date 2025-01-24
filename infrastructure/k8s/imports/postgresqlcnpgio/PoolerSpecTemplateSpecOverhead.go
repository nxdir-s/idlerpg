package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type PoolerSpecTemplateSpecOverhead interface {
	Value() interface{}
}

// The jsii proxy struct for PoolerSpecTemplateSpecOverhead
type jsiiProxy_PoolerSpecTemplateSpecOverhead struct {
	_ byte // padding
}

func (j *jsiiProxy_PoolerSpecTemplateSpecOverhead) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func PoolerSpecTemplateSpecOverhead_FromNumber(value *float64) PoolerSpecTemplateSpecOverhead {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecOverhead_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecOverhead

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecOverhead",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func PoolerSpecTemplateSpecOverhead_FromString(value *string) PoolerSpecTemplateSpecOverhead {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecOverhead_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecOverhead

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecOverhead",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

