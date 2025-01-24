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
type ClusterSpecManagedServicesAdditionalServiceTemplateSpecPortsTargetPort interface {
	Value() interface{}
}

// The jsii proxy struct for ClusterSpecManagedServicesAdditionalServiceTemplateSpecPortsTargetPort
type jsiiProxy_ClusterSpecManagedServicesAdditionalServiceTemplateSpecPortsTargetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_ClusterSpecManagedServicesAdditionalServiceTemplateSpecPortsTargetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ClusterSpecManagedServicesAdditionalServiceTemplateSpecPortsTargetPort_FromNumber(value *float64) ClusterSpecManagedServicesAdditionalServiceTemplateSpecPortsTargetPort {
	_init_.Initialize()

	if err := validateClusterSpecManagedServicesAdditionalServiceTemplateSpecPortsTargetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ClusterSpecManagedServicesAdditionalServiceTemplateSpecPortsTargetPort

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.ClusterSpecManagedServicesAdditionalServiceTemplateSpecPortsTargetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ClusterSpecManagedServicesAdditionalServiceTemplateSpecPortsTargetPort_FromString(value *string) ClusterSpecManagedServicesAdditionalServiceTemplateSpecPortsTargetPort {
	_init_.Initialize()

	if err := validateClusterSpecManagedServicesAdditionalServiceTemplateSpecPortsTargetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ClusterSpecManagedServicesAdditionalServiceTemplateSpecPortsTargetPort

	_jsii_.StaticInvoke(
		"postgresqlcnpgio.ClusterSpecManagedServicesAdditionalServiceTemplateSpecPortsTargetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

