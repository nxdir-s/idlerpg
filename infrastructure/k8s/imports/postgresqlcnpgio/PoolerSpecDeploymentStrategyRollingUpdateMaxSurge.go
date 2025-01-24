package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The maximum number of pods that can be scheduled above the desired number of pods.
//
// Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).
// This can not be 0 if MaxUnavailable is 0.
// Absolute number is calculated from percentage by rounding up.
// Defaults to 25%.
// Example: when this is set to 30%, the new ReplicaSet can be scaled up immediately when
// the rolling update starts, such that the total number of old and new pods do not exceed
// 130% of desired pods. Once old pods have been killed,
// new ReplicaSet can be scaled up further, ensuring that total number of pods running
// at any time during the update is at most 130% of desired pods.
// Default: 25%.
//
type PoolerSpecDeploymentStrategyRollingUpdateMaxSurge interface {
	Value() interface{}
}

// The jsii proxy struct for PoolerSpecDeploymentStrategyRollingUpdateMaxSurge
type jsiiProxy_PoolerSpecDeploymentStrategyRollingUpdateMaxSurge struct {
	_ byte // padding
}

func (j *jsiiProxy_PoolerSpecDeploymentStrategyRollingUpdateMaxSurge) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func PoolerSpecDeploymentStrategyRollingUpdateMaxSurge_FromNumber(value *float64) PoolerSpecDeploymentStrategyRollingUpdateMaxSurge {
	_init_.Initialize()

	if err := validatePoolerSpecDeploymentStrategyRollingUpdateMaxSurge_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecDeploymentStrategyRollingUpdateMaxSurge

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecDeploymentStrategyRollingUpdateMaxSurge",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func PoolerSpecDeploymentStrategyRollingUpdateMaxSurge_FromString(value *string) PoolerSpecDeploymentStrategyRollingUpdateMaxSurge {
	_init_.Initialize()

	if err := validatePoolerSpecDeploymentStrategyRollingUpdateMaxSurge_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecDeploymentStrategyRollingUpdateMaxSurge

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecDeploymentStrategyRollingUpdateMaxSurge",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

