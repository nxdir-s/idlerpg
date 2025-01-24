package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The maximum number of pods that can be unavailable during the update.
//
// Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).
// Absolute number is calculated from percentage by rounding down.
// This can not be 0 if MaxSurge is 0.
// Defaults to 25%.
// Example: when this is set to 30%, the old ReplicaSet can be scaled down to 70% of desired pods
// immediately when the rolling update starts. Once new pods are ready, old ReplicaSet
// can be scaled down further, followed by scaling up the new ReplicaSet, ensuring
// that the total number of pods available at all times during the update is at
// least 70% of desired pods.
// Default: 25%.
//
type PoolerSpecDeploymentStrategyRollingUpdateMaxUnavailable interface {
	Value() interface{}
}

// The jsii proxy struct for PoolerSpecDeploymentStrategyRollingUpdateMaxUnavailable
type jsiiProxy_PoolerSpecDeploymentStrategyRollingUpdateMaxUnavailable struct {
	_ byte // padding
}

func (j *jsiiProxy_PoolerSpecDeploymentStrategyRollingUpdateMaxUnavailable) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func PoolerSpecDeploymentStrategyRollingUpdateMaxUnavailable_FromNumber(value *float64) PoolerSpecDeploymentStrategyRollingUpdateMaxUnavailable {
	_init_.Initialize()

	if err := validatePoolerSpecDeploymentStrategyRollingUpdateMaxUnavailable_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecDeploymentStrategyRollingUpdateMaxUnavailable

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecDeploymentStrategyRollingUpdateMaxUnavailable",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func PoolerSpecDeploymentStrategyRollingUpdateMaxUnavailable_FromString(value *string) PoolerSpecDeploymentStrategyRollingUpdateMaxUnavailable {
	_init_.Initialize()

	if err := validatePoolerSpecDeploymentStrategyRollingUpdateMaxUnavailable_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecDeploymentStrategyRollingUpdateMaxUnavailable

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecDeploymentStrategyRollingUpdateMaxUnavailable",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

