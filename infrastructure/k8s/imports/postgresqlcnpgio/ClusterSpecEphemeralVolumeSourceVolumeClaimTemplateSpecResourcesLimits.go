package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecResourcesLimits
type jsiiProxy_ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecResourcesLimits_FromNumber(value *float64) ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecResourcesLimits {
	_init_.Initialize()

	if err := validateClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecResourcesLimits_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecResourcesLimits

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecResourcesLimits_FromString(value *string) ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecResourcesLimits {
	_init_.Initialize()

	if err := validateClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecResourcesLimits_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecResourcesLimits

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

