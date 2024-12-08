package server

import (
	"example.com/charts/imports/k8s"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

const (
	ServerPort     int = 3000
	ServerReplicas int = 1
)

type GameServerProps struct {
	Image         *string
	Replicas      *float64
	Port          *float64
	ContainerPort *float64
}

func NewGameServer(scope constructs.Construct, id *string, props *GameServerProps) constructs.Construct {
	server := constructs.NewConstruct(scope, id)

	replicas := props.Replicas
	if replicas == nil {
		replicas = jsii.Number(ServerReplicas)
	}

	port := props.Port
	if port == nil {
		port = jsii.Number(ServerPort)
	}

	containerPort := props.ContainerPort
	if containerPort == nil {
		containerPort = jsii.Number(ServerPort)
	}

	label := map[string]*string{"app": cdk8s.Names_ToLabelValue(server, nil)}

	configMap := NewGSConfig(server, jsii.String(*id+"-cm"), nil)

	service := k8s.NewKubeService(server, jsii.String(*id+"-srv"), &k8s.KubeServiceProps{
		Metadata: &k8s.ObjectMeta{
			Name: id,
		},
		Spec: &k8s.ServiceSpec{
			Type:     jsii.String("LoadBalancer"),
			Selector: &label,
			Ports: &[]*k8s.ServicePort{
				{
					Protocol:   jsii.String("TCP"),
					Port:       port,
					TargetPort: k8s.IntOrString_FromNumber(containerPort),
				},
			},
		},
	})

	deployment := k8s.NewKubeStatefulSet(server, id, &k8s.KubeStatefulSetProps{
		Metadata: &k8s.ObjectMeta{
			Name: id,
		},
		Spec: &k8s.StatefulSetSpec{
			ServiceName: service.Name(),
			Replicas:    replicas,
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
							Image:           props.Image,
							ImagePullPolicy: jsii.String("Never"),
							Name:            id,
							Ports: &[]*k8s.ContainerPort{
								{
									ContainerPort: port,
								},
							},
							EnvFrom: &[]*k8s.EnvFromSource{
								{
									ConfigMapRef: &k8s.ConfigMapEnvSource{
										Name: configMap.Name(),
									},
								},
							},
							Resources: &k8s.ResourceRequirements{
								Requests: &map[string]k8s.Quantity{
									"cpu":    k8s.Quantity_FromString(jsii.String("2000m")),
									"memory": k8s.Quantity_FromString(jsii.String("2Gi")),
								},
								Limits: &map[string]k8s.Quantity{
									"cpu":    k8s.Quantity_FromString(jsii.String("3000m")),
									"memory": k8s.Quantity_FromString(jsii.String("4Gi")),
								},
							},
						},
					},
				},
			},
		},
	})

	deployment.AddDependency(service)

	return server
}

type GSConfigProps struct{}

func NewGSConfig(scope constructs.Construct, id *string, props *GSConfigProps) k8s.KubeConfigMap {
	return k8s.NewKubeConfigMap(scope, id, &k8s.KubeConfigMapProps{
		Immutable: jsii.Bool(true),
		Data: &map[string]*string{
			"BROKERS": jsii.String("kafka-1:19092,kafka-2:19092,kafka-3:19092"),
		},
	})
}
