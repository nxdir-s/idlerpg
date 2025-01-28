package main

import (
	"example.com/charts/consumers"
	"example.com/charts/imports/k8s"
	"example.com/charts/servers"
	"example.com/charts/storage"

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

	storageChart := cdk8s.NewChart(scope, jsii.String(id+"-storage"), &cprops)
	storageChart.AddDependency(namespaceChart)

	storage.NewPostgresCluster(storageChart, jsii.String("database"), &storage.PostgresProps{
		Namespace: namespace,
	})

	chart.AddDependency(namespaceChart)
	chart.AddDependency(storageChart)

	servers.NewGameServer(chart, jsii.String("gameserver"), &servers.GameServerProps{
		Namespace: namespace,
		Image:     jsii.String("idlerpg/gameserver:latest"),
	})

	servers.NewWebServer(chart, jsii.String("webserver"), &servers.WebServerProps{
		Namespace: namespace,
		Image:     jsii.String("idlerpg/webserver:latest"),
	})

	consumers.NewUserEvents(chart, jsii.String("userevents"), &consumers.UserEventsProps{
		Namespace: namespace,
		Image:     jsii.String("idlerpg/userevents:latest"),
	})

	return chart
}

func main() {
	app := cdk8s.NewApp(nil)

	NewRpgChart(app, "idlerpg", nil)

	app.Synth()
}
