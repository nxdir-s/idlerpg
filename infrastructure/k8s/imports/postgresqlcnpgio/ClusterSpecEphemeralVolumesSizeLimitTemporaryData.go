package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// TemporaryData is the size limit of the temporary data volume.
type ClusterSpecEphemeralVolumesSizeLimitTemporaryData interface {
	Value() interface{}
}

// The jsii proxy struct for ClusterSpecEphemeralVolumesSizeLimitTemporaryData
type jsiiProxy_ClusterSpecEphemeralVolumesSizeLimitTemporaryData struct {
	_ byte // padding
}

func (j *jsiiProxy_ClusterSpecEphemeralVolumesSizeLimitTemporaryData) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ClusterSpecEphemeralVolumesSizeLimitTemporaryData_FromNumber(value *float64) ClusterSpecEphemeralVolumesSizeLimitTemporaryData {
	_init_.Initialize()

	if err := validateClusterSpecEphemeralVolumesSizeLimitTemporaryData_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ClusterSpecEphemeralVolumesSizeLimitTemporaryData

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.ClusterSpecEphemeralVolumesSizeLimitTemporaryData",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ClusterSpecEphemeralVolumesSizeLimitTemporaryData_FromString(value *string) ClusterSpecEphemeralVolumesSizeLimitTemporaryData {
	_init_.Initialize()

	if err := validateClusterSpecEphemeralVolumesSizeLimitTemporaryData_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ClusterSpecEphemeralVolumesSizeLimitTemporaryData

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.ClusterSpecEphemeralVolumesSizeLimitTemporaryData",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

