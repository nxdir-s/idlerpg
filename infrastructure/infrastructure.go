package main

import (
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	ecs "github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	elb "github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/nxdir-s/idlerpg/infrastructure/stacks"
)

const (
	_ int = iota
	Dev
	Prod
)

const (
	AppName string = "idlerpg-web"

	OtelServiceDev  string = AppName + "-dev"
	OtelServiceProd string = AppName

	FargateServiceName              string = AppName + "-fargate-service"
	FargateDesiredCount             int    = 1
	FargateCapacityProvidersEnabled bool   = true
	FargateCpu                      int    = 4096

	LatestImage string = "latest"

	BlueTGName  string = "alb-blue-tg"
	GreenTGName string = "alb-green-tg"

	VpcIpAddresses string = "10.45.0.0/16"
	TcpPort        int    = 80

	AlbOpen   bool = false
	AlbPublic bool = true

	SgDescription      string = "Allows access on port 80/http"
	AllowOutboundRules bool   = true

	DefaultRemoteRule bool = false

	AccountEnvKey string = "CDK_DEFAULT_ACCOUNT"
	RegionEnvKey  string = "CDK_DEFAULT_REGION"

	PipelineName string = AppName + "-pipeline"
)

type InfrastructureStackProps struct {
	props   awscdk.StackProps
	name    string
	enabled bool
}

func NewInfrastructureStack(scope constructs.Construct, id string, props *InfrastructureStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.props
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	vpc := stacks.NewVpc(stack, id+"Vpc", &stacks.VpcProps{
		IpAddresses: VpcIpAddresses,
	})

	// Creates a new blue Target Group that routes traffic from the public Application Load Balancer (ALB) to the
	// registered targets within the Target Group e.g. (EC2 instances, IP addresses, Lambda functions)
	// https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-target-groups.html
	targetGroupBlue := stacks.NewTargetGroup(stack, id+"TGB", &stacks.TargetGroupProps{
		Name:       BlueTGName,
		TargetType: elb.TargetType_IP,
		Vpc:        vpc,
		Port:       TcpPort,
	})

	targetGroupGreen := stacks.NewTargetGroup(stack, id+"TGG", &stacks.TargetGroupProps{
		Name:       GreenTGName,
		TargetType: elb.TargetType_IP,
		Vpc:        vpc,
		Port:       TcpPort,
	})

	// Creates an Elastic Container Registry (ECR) image repository
	imageRepo := stacks.NewEcrRepository(stack, id+"Image", &stacks.EcrRepositoryProps{
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	})

	ecsCluster := stacks.NewEcsCluster(stack, id+"Cluster", &stacks.EcsClusterProps{
		Vpc:                            vpc,
		EnableFargateCapacityProviders: FargateCapacityProvidersEnabled,
	})

	logDriver := stacks.NewLogDriver(stack, &stacks.LogDriverProps{
		StreamPrefix: id + "Container",
		Mode:         ecs.AwsLogDriverMode_NON_BLOCKING,
	})

	taskDef := stacks.NewFargateTaskDefinition(stack, id+"Def", &stacks.FargateTaskDefinitionProps{
		ImageRepo:     imageRepo,
		ImageTag:      LatestImage,
		ContainerName: id + "Container",
		Port:          TcpPort,
		LogDriver:     logDriver,
		Cpu:           FargateCpu,
		PidMode:       ecs.PidMode_TASK,
		RuntimePlatform: &ecs.RuntimePlatform{
			OperatingSystemFamily: ecs.OperatingSystemFamily_LINUX(),
			CpuArchitecture:       ecs.CpuArchitecture_ARM64(),
		},
	})

	fargateService := stacks.NewFargateService(stack, id+"Serv", &stacks.FargateServiceProps{
		ServiceName:    FargateServiceName,
		DesiredCount:   FargateDesiredCount,
		TaskDefinition: taskDef,
		Vpc:            vpc,
		Cluster:        ecsCluster,
		// Sets CodeDeploy as the deployment controller
		DeploymentController: &ecs.DeploymentController{
			Type: ecs.DeploymentControllerType_CODE_DEPLOY,
		},
	})

	// Adds the ECS Fargate service to the ALB target group
	fargateService.AttachToApplicationTargetGroup(targetGroupBlue)

	// Creates a Security Group for the Application Load Balancer (ALB)
	albSg := stacks.NewAlbSecurityGroup(stack, id+"SG", &stacks.SecurityGroupProps{
		Vpc:           vpc,
		Port:          TcpPort,
		AllowOutbound: AllowOutboundRules,
		RemoteRule:    DefaultRemoteRule,
		Description:   SgDescription,
	})

	publicAlb := stacks.NewAlb(stack, id+"Alb", &stacks.AlbProps{
		Vpc:            vpc,
		SecurityGroups: albSg,
		InternetFacing: AlbPublic,
	})

	// Adds a listener on port 80 to the ALB
	albListener := publicAlb.AddListener(jsii.String(id+"Listener"), &elb.BaseApplicationListenerProps{
		Port: jsii.Number(TcpPort),
		Open: jsii.Bool(AlbOpen),
		DefaultTargetGroups: &[]elb.IApplicationTargetGroup{
			targetGroupBlue,
		},
	})

	stacks.NewDeployPipelineStack(stack, id+"Pipeline", &stacks.PipelineNestedStackProps{
		PipelineName:   PipelineName,
		AppName:        props.name,
		FargateService: fargateService,
		TaskDefinition: taskDef,
		Repository:     imageRepo,
		Vpc:            vpc,
		LoadBalancer:   publicAlb,
		Listener:       albListener,
		TargetGroupB:   targetGroupBlue,
		TargetGroupG:   targetGroupGreen,
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewInfrastructureStack(app, "IdleRpgWeb", &InfrastructureStackProps{
		props: awscdk.StackProps{
			Env: env(),
		},
		name:    AppName,
		enabled: true,
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return &awscdk.Environment{
		Account: jsii.String(os.Getenv(AccountEnvKey)),
		Region:  jsii.String(os.Getenv(RegionEnvKey)),
	}

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}

func envVars(region int) *map[string]*string {
	baseDev := map[string]*string{
		"OTEL_SERVICE_NAME": jsii.String(OtelServiceDev),
	}

	baseProd := map[string]*string{
		"OTEL_SERVICE_NAME": jsii.String(OtelServiceProd),
	}

	switch region {
	case Dev:
		// maps.Copy(baseDev, map[string]*string{
		//     "EAST_ENV": jsii.String("example")
		// })

		return &baseDev
	case Prod:
		// maps.Copy(baseProd, map[string]*string{
		//     "PROD_ENV": jsii.String("example")
		// })

		return &baseProd
	default:
		return &map[string]*string{}
	}
}
