package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Shm is the size limit of the shared memory volume.
type ClusterSpecEphemeralVolumesSizeLimitShm interface {
	Value() interface{}
}

// The jsii proxy struct for ClusterSpecEphemeralVolumesSizeLimitShm
type jsiiProxy_ClusterSpecEphemeralVolumesSizeLimitShm struct {
	_ byte // padding
}

func (j *jsiiProxy_ClusterSpecEphemeralVolumesSizeLimitShm) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ClusterSpecEphemeralVolumesSizeLimitShm_FromNumber(value *float64) ClusterSpecEphemeralVolumesSizeLimitShm {
	_init_.Initialize()

	if err := validateClusterSpecEphemeralVolumesSizeLimitShm_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ClusterSpecEphemeralVolumesSizeLimitShm

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.ClusterSpecEphemeralVolumesSizeLimitShm",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ClusterSpecEphemeralVolumesSizeLimitShm_FromString(value *string) ClusterSpecEphemeralVolumesSizeLimitShm {
	_init_.Initialize()

	if err := validateClusterSpecEphemeralVolumesSizeLimitShm_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ClusterSpecEphemeralVolumesSizeLimitShm

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.ClusterSpecEphemeralVolumesSizeLimitShm",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

