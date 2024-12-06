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

	configMap := NewControllerCfg(controller, jsii.String(*id+"-cm"), &ControllerCfgProps{
		ControllerPort: port,
		NodeID:         props.NodeID,
		ClusterID:      props.ClusterID,
	})

	k8s.NewKubeService(controller, id, &k8s.KubeServiceProps{
		Spec: &k8s.ServiceSpec{
			Ports:    &[]*k8s.ServicePort{{Port: port, TargetPort: k8s.IntOrString_FromNumber(containerPort)}},
			Selector: &label,
		},
	})

	k8s.NewKubeDeployment(controller, id, &k8s.KubeDeploymentProps{
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
							},
						},
					},
				},
			},
		},
	})

	return controller
}

type ControllerCfgProps struct {
	ControllerPort *float64
	NodeID         *float64
	ClusterID      *string
}

func NewControllerCfg(scope constructs.Construct, id *string, props *ControllerCfgProps) k8s.KubeConfigMap {
	return k8s.NewKubeConfigMap(scope, id, &k8s.KubeConfigMapProps{
		Immutable: jsii.Bool(true),
		Data: &map[string]*string{
			"KAFKA_NODE_ID":                                  jsii.String(strconv.Itoa(int(*props.NodeID))),
			"KAFKA_PROCESS_ROLES":                            jsii.String("controller"),
			"KAFKA_CONTROLLER_QUORUM_VOTERS":                 jsii.String(fmt.Sprintf("1@controller-1:%d,2@controller-2:%d,3@controller-3:%d", props.ControllerPort, props.ControllerPort, props.ControllerPort)),
			"KAFKA_INTER_BROKER_LISTENER_NAME":               jsii.String("PLAINTEXT"),
			"KAFKA_CONTROLLER_LISTENER_NAMES":                jsii.String("CONTROLLER"),
			"KAFKA_LISTENERS":                                jsii.String(fmt.Sprintf("CONTROLLER://:%d", props.ControllerPort)),
			"CLUSTER_ID":                                     props.ClusterID,
			"KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR":         jsii.String("3"),
			"KAFKA_GROUP_INITIAL_REBALANCE_DELAY_MS":         jsii.String("0"),
			"KAFKA_TRANSACTION_STATE_LOG_MIN_ISR":            jsii.String("2"),
			"KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR": jsii.String("3"),
			"KAFKA_LOG_DIRS":                                 jsii.String("/tmp/kraft-combined-logs"),
		},
	})
}
