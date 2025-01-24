package postgresqlcnpgio

import (
	_init_ "example.com/charts/imports/postgresqlcnpgio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Number or name of the port to access on the pods targeted by the service.
//
// Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
// If this is a string, it will be looked up as a named port in the
// target Pod's container ports. If this is not specified, the value
// of the 'port' field is used (an identity map).
// This field is ignored for services with clusterIP=None, and should be
// omitted or set equal to the 'port' field.
// More info: https://kubernetes.io/docs/concepts/services-networking/service/#defining-a-service
type PoolerSpecServiceTemplateSpecPortsTargetPort interface {
	Value() interface{}
}

// The jsii proxy struct for PoolerSpecServiceTemplateSpecPortsTargetPort
type jsiiProxy_PoolerSpecServiceTemplateSpecPortsTargetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_PoolerSpecServiceTemplateSpecPortsTargetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func PoolerSpecServiceTemplateSpecPortsTargetPort_FromNumber(value *float64) PoolerSpecServiceTemplateSpecPortsTargetPort {
	_init_.Initialize()

	if err := validatePoolerSpecServiceTemplateSpecPortsTargetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecServiceTemplateSpecPortsTargetPort

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecServiceTemplateSpecPortsTargetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func PoolerSpecServiceTemplateSpecPortsTargetPort_FromString(value *string) PoolerSpecServiceTemplateSpecPortsTargetPort {
	_init_.Initialize()

	if err := validatePoolerSpecServiceTemplateSpecPortsTargetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns PoolerSpecServiceTemplateSpecPortsTargetPort

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.PoolerSpecServiceTemplateSpecPortsTargetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

