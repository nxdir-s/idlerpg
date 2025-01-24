package main

import (
	"example.com/charts/imports/k8s"

	"example.com/charts/servers"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type RpgChartProps struct {
	cdk8s.ChartProps
}

func NewRpgChart(scope constructs.Construct, id string, props *RpgChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)

	namespaceChart := cdk8s.NewChart(scope, jsii.String(id+"-namespace"), &cprops)

	// kafka.NewCluster(chart, jsii.String("kafka"), nil)

	namespace := k8s.NewKubeNamespace(namespaceChart, jsii.String(id), &k8s.KubeNamespaceProps{
		Metadata: &k8s.ObjectMeta{
			Name: jsii.String(id),
		},
	})

	servers.NewGameServer(chart, jsii.String("gameserver"), &servers.GameServerProps{
		Namespace: namespace,
		Image:     jsii.String("idlerpg/gameserver:latest"),
	})

	servers.NewWebServer(chart, jsii.String("webserver"), &servers.WebServerProps{
		Namespace: namespace,
		Image:     jsii.String("idlerpg/webserver:latest"),
	})

	servers.NewConsumer(chart, jsii.String("consumer"), &servers.ConsumerProps{
		Namespace: namespace,
		Image:     jsii.String("idlerpg/consumer:latest"),
	})

	return chart
}

func main() {
	app := cdk8s.NewApp(nil)

	NewRpgChart(app, "idlerpg", nil)

	app.Synth()
}
