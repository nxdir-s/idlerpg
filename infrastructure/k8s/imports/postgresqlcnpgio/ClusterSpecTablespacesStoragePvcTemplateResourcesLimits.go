package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type ClusterSpecTablespacesStoragePvcTemplateResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for ClusterSpecTablespacesStoragePvcTemplateResourcesLimits
type jsiiProxy_ClusterSpecTablespacesStoragePvcTemplateResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_ClusterSpecTablespacesStoragePvcTemplateResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ClusterSpecTablespacesStoragePvcTemplateResourcesLimits_FromNumber(value *float64) ClusterSpecTablespacesStoragePvcTemplateResourcesLimits {
	_init_.Initialize()

	if err := validateClusterSpecTablespacesStoragePvcTemplateResourcesLimits_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ClusterSpecTablespacesStoragePvcTemplateResourcesLimits

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.ClusterSpecTablespacesStoragePvcTemplateResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ClusterSpecTablespacesStoragePvcTemplateResourcesLimits_FromString(value *string) ClusterSpecTablespacesStoragePvcTemplateResourcesLimits {
	_init_.Initialize()

	if err := validateClusterSpecTablespacesStoragePvcTemplateResourcesLimits_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ClusterSpecTablespacesStoragePvcTemplateResourcesLimits

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.ClusterSpecTablespacesStoragePvcTemplateResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

