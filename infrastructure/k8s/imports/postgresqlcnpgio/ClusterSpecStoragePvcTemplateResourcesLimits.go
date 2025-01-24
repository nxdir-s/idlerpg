package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type ClusterSpecStoragePvcTemplateResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for ClusterSpecStoragePvcTemplateResourcesLimits
type jsiiProxy_ClusterSpecStoragePvcTemplateResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_ClusterSpecStoragePvcTemplateResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ClusterSpecStoragePvcTemplateResourcesLimits_FromNumber(value *float64) ClusterSpecStoragePvcTemplateResourcesLimits {
	_init_.Initialize()

	if err := validateClusterSpecStoragePvcTemplateResourcesLimits_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ClusterSpecStoragePvcTemplateResourcesLimits

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.ClusterSpecStoragePvcTemplateResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ClusterSpecStoragePvcTemplateResourcesLimits_FromString(value *string) ClusterSpecStoragePvcTemplateResourcesLimits {
	_init_.Initialize()

	if err := validateClusterSpecStoragePvcTemplateResourcesLimits_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ClusterSpecStoragePvcTemplateResourcesLimits

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.ClusterSpecStoragePvcTemplateResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

