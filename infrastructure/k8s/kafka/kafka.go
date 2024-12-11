package kafka

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

const (
	KafkaImage string = "apache/kafka:latest"
	ClusterID  string = "4L6g3nShT-eMCtK--X86sw"

	ControllerOne   string = "controller-1"
	ControllerTwo   string = "controller-2"
	ControllerThree string = "controller-3"

	BrokerOne   string = "kafka-1"
	BrokerTwo   string = "kafka-2"
	BrokerThree string = "kafka-3"

	BrokerOnePort   float64 = 29092
	BrokerTwoPort   float64 = 39092
	BrokerThreePort float64 = 49092

	ControllerOneID   int = 1
	ControllerTwoID   int = 2
	ControllerThreeID int = 3

	BrokerOneID   int = 4
	BrokerTwoID   int = 5
	BrokerThreeID int = 6
)

type ClusterProps struct{}

func NewCluster(scope constructs.Construct, id *string, props *ClusterProps) constructs.Construct {
	kafka := constructs.NewConstruct(scope, id)

	NewController(kafka, jsii.String(ControllerOne), &ControllerProps{
		Image:     jsii.String(KafkaImage),
		NodeID:    jsii.Number(ControllerOneID),
		ClusterID: jsii.String(ClusterID),
	})

	NewController(kafka, jsii.String(ControllerTwo), &ControllerProps{
		Image:     jsii.String(KafkaImage),
		NodeID:    jsii.Number(ControllerTwoID),
		ClusterID: jsii.String(ClusterID),
	})

	NewController(kafka, jsii.String(ControllerThree), &ControllerProps{
		Image:     jsii.String(KafkaImage),
		NodeID:    jsii.Number(ControllerThreeID),
		ClusterID: jsii.String(ClusterID),
	})

	NewBroker(kafka, jsii.String(BrokerOne), &BrokerProps{
		Image:     jsii.String(KafkaImage),
		Port:      jsii.Number(BrokerOnePort),
		NodeID:    jsii.Number(BrokerOneID),
		ClusterID: jsii.String(ClusterID),
	})

	NewBroker(kafka, jsii.String(BrokerTwo), &BrokerProps{
		Image:     jsii.String(KafkaImage),
		Port:      jsii.Number(BrokerTwoPort),
		NodeID:    jsii.Number(BrokerTwoID),
		ClusterID: jsii.String(ClusterID),
	})

	NewBroker(kafka, jsii.String(BrokerThree), &BrokerProps{
		Image:     jsii.String(KafkaImage),
		Port:      jsii.Number(BrokerThreePort),
		NodeID:    jsii.Number(BrokerThreeID),
		ClusterID: jsii.String(ClusterID),
	})

	return kafka
}
