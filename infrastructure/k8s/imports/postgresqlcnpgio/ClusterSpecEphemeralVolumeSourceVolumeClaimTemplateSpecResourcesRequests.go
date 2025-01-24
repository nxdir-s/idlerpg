package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecResourcesRequests interface {
	Value() interface{}
}

// The jsii proxy struct for ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecResourcesRequests
type jsiiProxy_ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecResourcesRequests struct {
	_ byte // padding
}

func (j *jsiiProxy_ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecResourcesRequests) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecResourcesRequests_FromNumber(value *float64) ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecResourcesRequests {
	_init_.Initialize()

	if err := validateClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecResourcesRequests_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecResourcesRequests

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecResourcesRequests",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecResourcesRequests_FromString(value *string) ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecResourcesRequests {
	_init_.Initialize()

	if err := validateClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecResourcesRequests_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecResourcesRequests

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecResourcesRequests",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

