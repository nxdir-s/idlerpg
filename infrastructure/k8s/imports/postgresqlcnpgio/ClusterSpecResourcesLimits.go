package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type ClusterSpecResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for ClusterSpecResourcesLimits
type jsiiProxy_ClusterSpecResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_ClusterSpecResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ClusterSpecResourcesLimits_FromNumber(value *float64) ClusterSpecResourcesLimits {
	_init_.Initialize()

	if err := validateClusterSpecResourcesLimits_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ClusterSpecResourcesLimits

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.ClusterSpecResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ClusterSpecResourcesLimits_FromString(value *string) ClusterSpecResourcesLimits {
	_init_.Initialize()

	if err := validateClusterSpecResourcesLimits_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ClusterSpecResourcesLimits

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.ClusterSpecResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

