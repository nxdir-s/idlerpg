package main

import (
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

	// kafka.NewCluster(chart, jsii.String("kafka"), nil)

	servers.NewGameServer(chart, jsii.String("gameserver"), &servers.GameServerProps{
		Image: jsii.String("idlerpg:latest"),
	})

	return chart
}

type ObservabilityChartProps struct {
	cdk8s.ChartProps
}

func NewObservabilityChart(scope constructs.Construct, id string, props *ObservabilityChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)

	return chart
}

func main() {
	app := cdk8s.NewApp(nil)

	NewRpgChart(app, "idlerpg", nil)

	app.Synth()
}
