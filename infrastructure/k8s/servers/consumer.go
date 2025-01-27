package servers

import (
	"example.com/charts/imports/k8s"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

const (
	ConsumerImage    string = "idlerpg/consumer:latest"
	ConsumerPort     int    = 5000
	ConsumerReplicas int    = 1
)

type ConsumerProps struct {
	Namespace k8s.KubeNamespace
	Image     *string
	Replicas  *float64
	Port      *float64
}

func NewConsumer(scope constructs.Construct, id *string, props *ConsumerProps) constructs.Construct {
	server := constructs.NewConstruct(scope, id)

	replicas := props.Replicas
	if replicas == nil {
		replicas = jsii.Number(ConsumerReplicas)
	}

	port := props.Port
	if port == nil {
		port = jsii.Number(ConsumerPort)
	}

	// label := map[string]*string{"app": cdk8s.Names_ToLabelValue(server, nil)}
	labels := map[string]*string{"name": id}

	configMap := NewConsumerCfg(server, jsii.String(*id+"-cm"), &ConsumerCfgProps{
		Namespace: props.Namespace,
	})

	service := k8s.NewKubeService(server, jsii.String(*id+"-srv"), &k8s.KubeServiceProps{
		Metadata: &k8s.ObjectMeta{
			Name:      id,
			Namespace: props.Namespace.Name(),
		},
		Spec: &k8s.ServiceSpec{
			Type:     jsii.String("LoadBalancer"),
			Selector: &labels,
			Ports: &[]*k8s.ServicePort{
				{
					Protocol:   jsii.String("TCP"),
					Port:       port,
					TargetPort: k8s.IntOrString_FromNumber(port),
				},
			},
		},
	})

	deployment := k8s.NewKubeStatefulSet(server, id, &k8s.KubeStatefulSetProps{
		Metadata: &k8s.ObjectMeta{
			Labels:    &labels,
			Name:      id,
			Namespace: props.Namespace.Name(),
		},
		Spec: &k8s.StatefulSetSpec{
			MinReadySeconds: jsii.Number(3),
			ServiceName:     service.Name(),
			Replicas:        replicas,
			Selector: &k8s.LabelSelector{
				MatchLabels: &labels,
			},
			UpdateStrategy: &k8s.StatefulSetUpdateStrategy{
				RollingUpdate: &k8s.RollingUpdateStatefulSetStrategy{
					MaxUnavailable: k8s.IntOrString_FromString(jsii.String("25%")),
				},
				Type: jsii.String("RollingUpdate"),
			},
			Template: &k8s.PodTemplateSpec{
				Metadata: &k8s.ObjectMeta{
					Labels: &map[string]*string{
						"name":                  labels["name"],
						"network/kafka-network": jsii.String("true"),
					},
					Namespace: props.Namespace.Name(),
				},
				Spec: &k8s.PodSpec{
					SecurityContext: &k8s.PodSecurityContext{
						FsGroup: jsii.Number(1000),
					},
					EnableServiceLinks: jsii.Bool(false),
					RestartPolicy:      jsii.String("Always"),
					Containers: &[]*k8s.Container{
						{
							Image:           jsii.String(ConsumerImage),
							ImagePullPolicy: jsii.String("Always"),
							Name:            id,
							Ports: &[]*k8s.ContainerPort{
								{
									ContainerPort: port,
								},
							},
							Env: &[]*k8s.EnvVar{
								{
									Name: jsii.String("BROKERS"),
									ValueFrom: &k8s.EnvVarSource{
										SecretKeyRef: &k8s.SecretKeySelector{
											Name: jsii.String("redpanda-brokers"),
											Key:  jsii.String("brokers"),
										},
									},
								},
								{
									Name: jsii.String("REDPANDA_SASL_USERNAME"),
									ValueFrom: &k8s.EnvVarSource{
										SecretKeyRef: &k8s.SecretKeySelector{
											Name: jsii.String("redpanda-auth"),
											Key:  jsii.String("username"),
										},
									},
								},
								{
									Name: jsii.String("REDPANDA_SASL_PASSWORD"),
									ValueFrom: &k8s.EnvVarSource{
										SecretKeyRef: &k8s.SecretKeySelector{
											Name: jsii.String("redpanda-auth"),
											Key:  jsii.String("password"),
										},
									},
								},
								{
									Name: jsii.String("PROFILE_URL"),
									ValueFrom: &k8s.EnvVarSource{
										SecretKeyRef: &k8s.SecretKeySelector{
											Name: jsii.String("profiling"),
											Key:  jsii.String("endpoint"),
										},
									},
								},
								{
									Name: jsii.String("GCLOUD_USER"),
									ValueFrom: &k8s.EnvVarSource{
										SecretKeyRef: &k8s.SecretKeySelector{
											Name: jsii.String("grafana-auth"),
											Key:  jsii.String("username"),
										},
									},
								},
								{
									Name: jsii.String("GCLOUD_PASSWORD"),
									ValueFrom: &k8s.EnvVarSource{
										SecretKeyRef: &k8s.SecretKeySelector{
											Name: jsii.String("grafana-auth"),
											Key:  jsii.String("password"),
										},
									},
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
									"cpu":    k8s.Quantity_FromString(jsii.String("100m")),
									"memory": k8s.Quantity_FromString(jsii.String("32Mi")),
								},
							},
						},
					},
				},
			},
		},
	})

	deployment.AddDependency(service)
	deployment.AddDependency(props.Namespace)

	return server
}

type ConsumerCfgProps struct {
	Namespace k8s.KubeNamespace
	Port      *float64
}

func NewConsumerCfg(scope constructs.Construct, id *string, props *ConsumerCfgProps) k8s.KubeConfigMap {
	return k8s.NewKubeConfigMap(scope, id, &k8s.KubeConfigMapProps{
		Metadata: &k8s.ObjectMeta{
			Name:      id,
			Namespace: props.Namespace.Name(),
		},
		Immutable: jsii.Bool(false),
		Data: &map[string]*string{
			"OTEL_EXPORTER_OTLP_TRACES_INSECURE": jsii.String("true"),
			"OTEL_EXPORTER_OTLP_ENDPOINT":        jsii.String("grafana-k8s-monitoring-alloy.default.svc.cluster.local:4317"),
			"OTEL_SERVICE_NAME":                  jsii.String("consumer"),
			"CONSUMER_GROUP_NAME":                jsii.String("consumer-dev"),
		},
	})
}
