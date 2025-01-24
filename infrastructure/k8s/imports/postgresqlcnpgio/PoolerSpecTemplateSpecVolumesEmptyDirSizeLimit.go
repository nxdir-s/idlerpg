package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// sizeLimit is the total amount of local storage required for this EmptyDir volume.
//
// The size limit is also applicable for memory medium.
// The maximum usage on memory medium EmptyDir would be the minimum value between
// the SizeLimit specified here and the sum of memory limits of all containers in a pod.
// The default is nil which means that the limit is undefined.
// More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir
type PoolerSpecTemplateSpecVolumesEmptyDirSizeLimit interface {
	Value() interface{}
}

// The jsii proxy struct for PoolerSpecTemplateSpecVolumesEmptyDirSizeLimit
type jsiiProxy_PoolerSpecTemplateSpecVolumesEmptyDirSizeLimit struct {
	_ byte // padding
}

func (j *jsiiProxy_PoolerSpecTemplateSpecVolumesEmptyDirSizeLimit) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func PoolerSpecTemplateSpecVolumesEmptyDirSizeLimit_FromNumber(value *float64) PoolerSpecTemplateSpecVolumesEmptyDirSizeLimit {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecVolumesEmptyDirSizeLimit_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecVolumesEmptyDirSizeLimit

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesEmptyDirSizeLimit",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func PoolerSpecTemplateSpecVolumesEmptyDirSizeLimit_FromString(value *string) PoolerSpecTemplateSpecVolumesEmptyDirSizeLimit {
	_init_.Initialize()

	if err := validatePoolerSpecTemplateSpecVolumesEmptyDirSizeLimit_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecTemplateSpecVolumesEmptyDirSizeLimit

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesEmptyDirSizeLimit",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

