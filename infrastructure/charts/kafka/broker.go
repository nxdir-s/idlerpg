package kafka

import (
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
	Image         *string
	Replicas      *float64
	Port          *float64
	ContainerPort *float64
	InternalPort  *float64
	NodeID        *float64
	ClusterID     *string
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

	internalPort := props.InternalPort
	if containerPort == nil {
		containerPort = jsii.Number(BrokerInternalPort)
	}

	label := map[string]*string{"app": cdk8s.Names_ToLabelValue(broker, nil)}

	k8s.NewKubeService(broker, jsii.String("service"), &k8s.KubeServiceProps{
		Spec: &k8s.ServiceSpec{
			ClusterIp: jsii.String("None"),
			Ports: &[]*k8s.ServicePort{
				{
					Port:       port,
					TargetPort: k8s.IntOrString_FromNumber(containerPort),
				},
			},
			Selector: &label,
		},
	})

	k8s.NewKubeStatefulSet(broker, jsii.String("deployment"), &k8s.KubeStatefulSetProps{
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
							Ports: &[]*k8s.ContainerPort{
								{
									ContainerPort: containerPort,
								},
								{
									ContainerPort: internalPort,
								},
							},
							Env: &[]*k8s.EnvVar{
								{
									Name:  jsii.String("KAFKA_NODE_ID"),
									Value: jsii.String(strconv.Itoa(int(*props.NodeID))),
								},
								{
									Name:  jsii.String("KAFKA_PROCESS_ROLES"),
									Value: jsii.String("broker"),
								},
								{
									Name:  jsii.String("KAFKA_CONTROLLER_QUORUM_VOTERS"),
									Value: jsii.String("1@controller-1:9093,2@controller-2:9093,3@controller-3:9093"),
								},
								{
									Name:  jsii.String("KAFKA_LISTENERS"),
									Value: jsii.String("PLAINTEXT://:19092,PLAINTEXT_HOST://:9092"),
								},
								{
									Name:  jsii.String("KAFKA_LISTENER_SECURITY_PROTOCOL_MAP"),
									Value: jsii.String("CONTROLLER:PLAINTEXT,PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT"),
								},
								{
									Name:  jsii.String("KAFKA_INTER_BROKER_LISTENER_NAME"),
									Value: jsii.String("PLAINTEXT"),
								},
								{
									Name:  jsii.String("KAFKA_CONTROLLER_LISTENER_NAMES"),
									Value: jsii.String("CONTROLLER"),
								},
								{
									Name:  jsii.String("CLUSTER_ID"),
									Value: props.ClusterID,
								},
								{
									Name:  jsii.String("KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR"),
									Value: jsii.String("3"),
								},
								{
									Name:  jsii.String("KAFKA_GROUP_INITIAL_REBALANCE_DELAY_MS"),
									Value: jsii.String("0"),
								},
								{
									Name:  jsii.String("KAFKA_TRANSACTION_STATE_LOG_MIN_ISR"),
									Value: jsii.String("2"),
								},
								{
									Name:  jsii.String("KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR"),
									Value: jsii.String("3"),
								},
								{
									Name:  jsii.String("KAFKA_LOG_DIRS"),
									Value: jsii.String("/tmp/kraft-combined-logs"),
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
