package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Specifies the output format of the exposed resources, defaults to "1".
type ClusterSpecProjectedVolumeTemplateSourcesDownwardApiItemsResourceFieldRefDivisor interface {
	Value() interface{}
}

// The jsii proxy struct for ClusterSpecProjectedVolumeTemplateSourcesDownwardApiItemsResourceFieldRefDivisor
type jsiiProxy_ClusterSpecProjectedVolumeTemplateSourcesDownwardApiItemsResourceFieldRefDivisor struct {
	_ byte // padding
}

func (j *jsiiProxy_ClusterSpecProjectedVolumeTemplateSourcesDownwardApiItemsResourceFieldRefDivisor) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ClusterSpecProjectedVolumeTemplateSourcesDownwardApiItemsResourceFieldRefDivisor_FromNumber(value *float64) ClusterSpecProjectedVolumeTemplateSourcesDownwardApiItemsResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateClusterSpecProjectedVolumeTemplateSourcesDownwardApiItemsResourceFieldRefDivisor_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ClusterSpecProjectedVolumeTemplateSourcesDownwardApiItemsResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.ClusterSpecProjectedVolumeTemplateSourcesDownwardApiItemsResourceFieldRefDivisor",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ClusterSpecProjectedVolumeTemplateSourcesDownwardApiItemsResourceFieldRefDivisor_FromString(value *string) ClusterSpecProjectedVolumeTemplateSourcesDownwardApiItemsResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateClusterSpecProjectedVolumeTemplateSourcesDownwardApiItemsResourceFieldRefDivisor_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ClusterSpecProjectedVolumeTemplateSourcesDownwardApiItemsResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.ClusterSpecProjectedVolumeTemplateSourcesDownwardApiItemsResourceFieldRefDivisor",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

