package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type ClusterSpecWalStoragePvcTemplateResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for ClusterSpecWalStoragePvcTemplateResourcesLimits
type jsiiProxy_ClusterSpecWalStoragePvcTemplateResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_ClusterSpecWalStoragePvcTemplateResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ClusterSpecWalStoragePvcTemplateResourcesLimits_FromNumber(value *float64) ClusterSpecWalStoragePvcTemplateResourcesLimits {
	_init_.Initialize()

	if err := validateClusterSpecWalStoragePvcTemplateResourcesLimits_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ClusterSpecWalStoragePvcTemplateResourcesLimits

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.ClusterSpecWalStoragePvcTemplateResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ClusterSpecWalStoragePvcTemplateResourcesLimits_FromString(value *string) ClusterSpecWalStoragePvcTemplateResourcesLimits {
	_init_.Initialize()

	if err := validateClusterSpecWalStoragePvcTemplateResourcesLimits_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ClusterSpecWalStoragePvcTemplateResourcesLimits

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.ClusterSpecWalStoragePvcTemplateResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

