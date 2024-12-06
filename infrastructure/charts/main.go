package main

import (
	"example.com/charts/kafka"
	"example.com/charts/server"
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

	kafka.NewCluster(chart, jsii.String("kafka"), nil)

	server.NewGameServer(chart, jsii.String("gameserver"), &server.GameServerProps{
		Image:         jsii.String("idlerpg:latest"),
		Port:          jsii.Number(3000),
		ContainerPort: jsii.Number(3000),
	})

	return chart
}

func main() {
	app := cdk8s.NewApp(nil)

	NewRpgChart(app, "IdleRpg", nil)

	app.Synth()
}
