package stacks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	ec2 "github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	ecr "github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	ecs "github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	elb "github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type FargateServiceProps struct {
	ServiceName          string
	DesiredCount         int
	Port                 int
	TaskDefinition       ecs.FargateTaskDefinition
	Vpc                  ec2.Vpc
	Cluster              ecs.ICluster
	DeploymentController *ecs.DeploymentController
}

func NewFargateService(scope constructs.Construct, id string, props *FargateServiceProps) ecs.FargateService {
	return ecs.NewFargateService(scope, jsii.String(id), &ecs.FargateServiceProps{
		DesiredCount:         jsii.Number(props.DesiredCount),
		ServiceName:          jsii.String(props.ServiceName),
		TaskDefinition:       props.TaskDefinition,
		Cluster:              props.Cluster,
		DeploymentController: props.DeploymentController,
	})
}

type FargateTaskDefinitionProps struct {
	ImageRepo       ecr.IRepository
	ImageTag        string
	ContainerName   string
	Port            int
	LogDriver       ecs.LogDriver
	Cpu             int
	PidMode         ecs.PidMode
	RuntimePlatform *ecs.RuntimePlatform
}

func NewFargateTaskDefinition(scope constructs.Construct, id string, props *FargateTaskDefinitionProps) ecs.FargateTaskDefinition {
	fargateTaskDef := ecs.NewFargateTaskDefinition(scope, jsii.String(id), &ecs.FargateTaskDefinitionProps{
		Cpu:             jsii.Number(props.Cpu),
		PidMode:         props.PidMode,
		RuntimePlatform: props.RuntimePlatform,
	})
	fargateTaskDef.AddContainer(jsii.String(props.ContainerName), &ecs.ContainerDefinitionOptions{
		ContainerName: jsii.String(props.ContainerName),
		Image:         ecs.ContainerImage_FromEcrRepository(props.ImageRepo, jsii.String(props.ImageTag)),
		PortMappings: &[]*ecs.PortMapping{
			{
				ContainerPort: jsii.Number(props.Port),
			},
		},
		Logging: props.LogDriver,
	})

	return fargateTaskDef
}

type EcrRepositoryProps struct {
	RemovalPolicy awscdk.RemovalPolicy
}

func NewEcrRepository(scope constructs.Construct, id string, props *EcrRepositoryProps) ecr.Repository {
	return ecr.NewRepository(scope, jsii.String(id), &ecr.RepositoryProps{
		RemovalPolicy: props.RemovalPolicy,
	})
}

type EcsClusterProps struct {
	Vpc                            ec2.Vpc
	EnableFargateCapacityProviders bool
}

func NewEcsCluster(scope constructs.Construct, id string, props *EcsClusterProps) ecs.ICluster {
	return ecs.NewCluster(scope, jsii.String(id), &ecs.ClusterProps{
		EnableFargateCapacityProviders: jsii.Bool(props.EnableFargateCapacityProviders),
		Vpc:                            props.Vpc,
	})
}

type TargetGroupProps struct {
	Name       string
	TargetType elb.TargetType
	Vpc        ec2.Vpc
	Port       int
}

func NewTargetGroup(scope constructs.Construct, id string, props *TargetGroupProps) elb.ApplicationTargetGroup {
	return elb.NewApplicationTargetGroup(scope, jsii.String(id), &elb.ApplicationTargetGroupProps{
		TargetGroupName: jsii.String(props.Name),
		TargetType:      props.TargetType,
		Vpc:             props.Vpc,
		Port:            jsii.Number(props.Port),
	})
}

type VpcProps struct {
	IpAddresses string
}

func NewVpc(scope constructs.Construct, id string, props *VpcProps) ec2.Vpc {
	return ec2.NewVpc(scope, jsii.String(id), &ec2.VpcProps{
		IpAddresses: ec2.IpAddresses_Cidr(
			jsii.String(props.IpAddresses),
		),
	})
}

type AlbProps struct {
	Vpc            ec2.Vpc
	SecurityGroups ec2.ISecurityGroup
	InternetFacing bool
}

func NewAlb(scope constructs.Construct, id string, props *AlbProps) elb.ApplicationLoadBalancer {
	return elb.NewApplicationLoadBalancer(scope, jsii.String(id), &elb.ApplicationLoadBalancerProps{
		Vpc:            props.Vpc,
		InternetFacing: jsii.Bool(props.InternetFacing),
		SecurityGroup:  props.SecurityGroups,
	})
}

type SecurityGroupProps struct {
	Vpc           ec2.Vpc
	Port          int
	AllowOutbound bool
	RemoteRule    bool
	Description   string
}

func NewAlbSecurityGroup(scope constructs.Construct, id string, props *SecurityGroupProps) ec2.SecurityGroup {
	albSg := ec2.NewSecurityGroup(scope, jsii.String(id), &ec2.SecurityGroupProps{
		Vpc:              props.Vpc,
		AllowAllOutbound: jsii.Bool(props.AllowOutbound),
	})
	albSg.AddIngressRule(ec2.Peer_AnyIpv4(), ec2.Port_Tcp(jsii.Number(props.Port)), jsii.String(props.Description), jsii.Bool(props.RemoteRule))

	return albSg
}

type LogDriverProps struct {
	StreamPrefix string
	Mode         ecs.AwsLogDriverMode
}

func NewLogDriver(scope constructs.Construct, props *LogDriverProps) ecs.LogDriver {
	return ecs.NewAwsLogDriver(&ecs.AwsLogDriverProps{
		StreamPrefix: jsii.String(props.StreamPrefix),
		Mode:         props.Mode,
	})
}
