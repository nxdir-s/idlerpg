package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type PoolerSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for PoolerSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits
type jsiiProxy_PoolerSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_PoolerSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func PoolerSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits_FromNumber(value *float64) PoolerSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func PoolerSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits_FromString(value *string) PoolerSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

