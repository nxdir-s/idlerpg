package kafka

import (
	"example.com/charts/imports/k8s"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

const (
	NetworkName  string = "kafka-network"
	NetworkLabel string = "network/" + NetworkName
)

type PolicyProps struct{}

func NewNetworkPolicy(scope constructs.Construct, id *string, props *PolicyProps) k8s.KubeNetworkPolicy {
	label := map[string]*string{
		NetworkLabel: jsii.String("true"),
	}

	return k8s.NewKubeNetworkPolicy(scope, id, &k8s.KubeNetworkPolicyProps{
		Metadata: &k8s.ObjectMeta{
			Name: jsii.String(NetworkName),
		},
		Spec: &k8s.NetworkPolicySpec{
			Ingress: &[]*k8s.NetworkPolicyIngressRule{
				{
					From: &[]*k8s.NetworkPolicyPeer{
						{
							PodSelector: &k8s.LabelSelector{
								MatchLabels: &label,
							},
						},
					},
				},
			},
			PodSelector: &k8s.LabelSelector{
				MatchLabels: &label,
			},
		},
	})
}
