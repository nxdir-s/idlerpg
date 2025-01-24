package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type ClusterSpecResourcesRequests interface {
	Value() interface{}
}

// The jsii proxy struct for ClusterSpecResourcesRequests
type jsiiProxy_ClusterSpecResourcesRequests struct {
	_ byte // padding
}

func (j *jsiiProxy_ClusterSpecResourcesRequests) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ClusterSpecResourcesRequests_FromNumber(value *float64) ClusterSpecResourcesRequests {
	_init_.Initialize()

	if err := validateClusterSpecResourcesRequests_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ClusterSpecResourcesRequests

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.ClusterSpecResourcesRequests",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ClusterSpecResourcesRequests_FromString(value *string) ClusterSpecResourcesRequests {
	_init_.Initialize()

	if err := validateClusterSpecResourcesRequests_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ClusterSpecResourcesRequests

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.ClusterSpecResourcesRequests",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

