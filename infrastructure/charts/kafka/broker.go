package kafka

import (
	"fmt"
	"strconv"

	"example.com/charts/imports/k8s"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

const (
	BrokerReplicas      float64 = 1
	BrokerPort          float64 = 29092
	BrokerInternalPort  float64 = 19092
	BrokerContainerPort float64 = 9092
)

type BrokerProps struct {
	Image          *string
	Replicas       *float64
	Port           *float64
	ContainerPort  *float64
	InternalPort   *float64
	NodeID         *float64
	ClusterID      *string
	ControllerPort *float64
}

func NewBroker(scope constructs.Construct, id *string, props *BrokerProps) constructs.Construct {
	broker := constructs.NewConstruct(scope, id)

	replicas := props.Replicas
	if replicas == nil {
		replicas = jsii.Number(BrokerReplicas)
	}

	port := props.Port
	if port == nil {
		port = jsii.Number(BrokerPort)
	}

	containerPort := props.ContainerPort
	if containerPort == nil {
		containerPort = jsii.Number(BrokerContainerPort)
	}

	controllerPort := props.ControllerPort
	if controllerPort == nil {
		controllerPort = jsii.Number(ControllerPort)
	}

	internalPort := props.InternalPort
	if internalPort == nil {
		internalPort = jsii.Number(BrokerInternalPort)
	}

	label := map[string]*string{"app": cdk8s.Names_ToLabelValue(broker, nil)}

	configMap := NewBrokerCfg(broker, jsii.String(*id+"-cm"), &BrokerCfgProps{
		BrokerID:       id,
		Port:           port,
		ControllerPort: controllerPort,
		ContainerPort:  containerPort,
		InternalPort:   internalPort,
		NodeID:         props.NodeID,
		ClusterID:      props.ClusterID,
	})

	k8s.NewKubeService(broker, id, &k8s.KubeServiceProps{
		Spec: &k8s.ServiceSpec{
			ClusterIp: jsii.String("None"),
			Ports: &[]*k8s.ServicePort{
				{
					Port:       port,
					TargetPort: k8s.IntOrString_FromNumber(containerPort),
				},
				{
					Port:       internalPort,
					TargetPort: k8s.IntOrString_FromNumber(internalPort),
				},
			},
			Selector: &label,
		},
	})

	k8s.NewKubeStatefulSet(broker, id, &k8s.KubeStatefulSetProps{
		Spec: &k8s.StatefulSetSpec{
			Replicas: replicas,
			Selector: &k8s.LabelSelector{
				MatchLabels: &label,
			},
			Template: &k8s.PodTemplateSpec{
				Metadata: &k8s.ObjectMeta{
					Labels: &map[string]*string{
						"app":                   label["app"],
						"network/kafka-network": jsii.String("true"),
					},
				},
				Spec: &k8s.PodSpec{
					SecurityContext: &k8s.PodSecurityContext{
						FsGroup: jsii.Number(1000),
					},
					EnableServiceLinks: jsii.Bool(false),
					RestartPolicy:      jsii.String("Always"),
					Containers: &[]*k8s.Container{
						{
							Name:            id,
							Image:           props.Image,
							ImagePullPolicy: jsii.String("Always"),
							EnvFrom: &[]*k8s.EnvFromSource{
								{
									ConfigMapRef: &k8s.ConfigMapEnvSource{
										Name: configMap.Name(),
									},
								},
							},
							Ports: &[]*k8s.ContainerPort{
								{
									ContainerPort: containerPort,
								},
								{
									ContainerPort: internalPort,
								},
								{
									ContainerPort: port,
								},
							},
						},
					},
				},
			},
		},
	})

	return broker
}

type BrokerCfgProps struct {
	BrokerID       *string
	Port           *float64
	ControllerPort *float64
	ContainerPort  *float64
	InternalPort   *float64
	NodeID         *float64
	ClusterID      *string
}

func NewBrokerCfg(scope constructs.Construct, id *string, props *BrokerCfgProps) k8s.KubeConfigMap {
	return k8s.NewKubeConfigMap(scope, id, &k8s.KubeConfigMapProps{
		Immutable: jsii.Bool(true),
		Data: &map[string]*string{
			"KAFKA_NODE_ID":                                  jsii.String(strconv.Itoa(int(*props.NodeID))),
			"KAFKA_PROCESS_ROLES":                            jsii.String("broker"),
			"KAFKA_CONTROLLER_QUORUM_VOTERS":                 jsii.String(fmt.Sprintf("1@controller-1:%d,2@controller-2:%d,3@controller-3:%d", props.ControllerPort, props.ControllerPort, props.ControllerPort)),
			"KAFKA_LISTENERS":                                jsii.String(fmt.Sprintf("PLAINTEXT://:%d,PLAINTEXT_HOST://:%d", props.InternalPort, props.ContainerPort)),
			"KAFKA_LISTENER_SECURITY_PROTOCOL_MAP":           jsii.String("CONTROLLER:PLAINTEXT,PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT"),
			"KAFKA_INTER_BROKER_LISTENER_NAME":               jsii.String("PLAINTEXT"),
			"KAFKA_ADVERTISED_LISTENERS":                     jsii.String(fmt.Sprintf("PLAINTEXT://%s:%d,PLAINTEXT_HOST://localhost:%d", *props.BrokerID, props.InternalPort, props.Port)),
			"KAFKA_CONTROLLER_LISTENER_NAMES":                jsii.String("CONTROLLER"),
			"CLUSTER_ID":                                     props.ClusterID,
			"KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR":         jsii.String("3"),
			"KAFKA_GROUP_INITIAL_REBALANCE_DELAY_MS":         jsii.String("0"),
			"KAFKA_TRANSACTION_STATE_LOG_MIN_ISR":            jsii.String("2"),
			"KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR": jsii.String("3"),
			"KAFKA_LOG_DIRS":                                 jsii.String("/tmp/kraft-combined-logs"),
		},
	})
}
