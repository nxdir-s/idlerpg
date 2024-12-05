package kafka

import (
	"strconv"

	"example.com/charts/imports/k8s"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

const (
	ControllerReplicas float64 = 1
	ControllerPort     float64 = 9093
	ContainerPort      float64 = 9093
)

type ControllerProps struct {
	Image         *string
	Replicas      *float64
	Port          *float64
	ContainerPort *float64
	NodeID        *float64
	ClusterID     *string
}

func NewController(scope constructs.Construct, id *string, props *ControllerProps) constructs.Construct {
	controller := constructs.NewConstruct(scope, id)

	replicas := props.Replicas
	if replicas == nil {
		replicas = jsii.Number(ControllerReplicas)
	}

	port := props.Port
	if port == nil {
		port = jsii.Number(ControllerPort)
	}

	containerPort := props.ContainerPort
	if containerPort == nil {
		containerPort = jsii.Number(ContainerPort)
	}

	label := map[string]*string{"app": cdk8s.Names_ToLabelValue(controller, nil)}

	k8s.NewKubeService(controller, jsii.String("service"), &k8s.KubeServiceProps{
		Spec: &k8s.ServiceSpec{
			Ports:    &[]*k8s.ServicePort{{Port: port, TargetPort: k8s.IntOrString_FromNumber(containerPort)}},
			Selector: &label,
		},
	})

	k8s.NewKubeDeployment(controller, jsii.String("deployment"), &k8s.KubeDeploymentProps{
		Spec: &k8s.DeploymentSpec{
			Replicas: replicas,
			Selector: &k8s.LabelSelector{
				MatchLabels: &label,
			},
			Template: &k8s.PodTemplateSpec{
				Metadata: &k8s.ObjectMeta{Labels: &map[string]*string{
					"app":                   label["app"],
					"network/kafka-network": jsii.String("true"),
				}},
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
							Ports:           &[]*k8s.ContainerPort{{ContainerPort: containerPort}},
							Env: &[]*k8s.EnvVar{
								{
									Name:  jsii.String("KAFKA_NODE_ID"),
									Value: jsii.String(strconv.Itoa(int(*props.NodeID))),
								},
								{
									Name:  jsii.String("KAFKA_PROCESS_ROLES"),
									Value: jsii.String("controller"),
								},
								{
									Name:  jsii.String("KAFKA_CONTROLLER_QUORUM_VOTERS"),
									Value: jsii.String("1@controller-1:9093,2@controller-2:9093,3@controller-3:9093"),
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
									Name:  jsii.String("KAFKA_LISTENERS"),
									Value: jsii.String("CONTROLLER://:9093"),
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

	return controller
}
