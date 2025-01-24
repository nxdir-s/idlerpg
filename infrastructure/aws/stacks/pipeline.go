package stacks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	codebuild "github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
	codecommit "github.com/aws/aws-cdk-go/awscdk/v2/awscodecommit"
	codedeploy "github.com/aws/aws-cdk-go/awscdk/v2/awscodedeploy"
	pipeline "github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	pipelineactions "github.com/aws/aws-cdk-go/awscdk/v2/awscodepipelineactions"
	ec2 "github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	ecr "github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	ecs "github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	elb "github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	iam "github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	lambda "github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

const (
	BuildArtifactName string = "BuildArtifact"
	BuildStageName    string = "Build"
	BuildActionName   string = "DockerBuildPush"
	BuildImageSpec    string = "buildspec.yaml"
	BuildImageTag     string = "latest"

	BuildLambdaAsset   string = "../scripts/trigger-build.mjs"
	BuildLambdaHandler string = "trigger-build.handler"

	SourceArtifactName string = "SourceArtifact"
	SourceStageName    string = "Source"
	SourceActionName   string = "CodeCommit"

	DeployStageName  string = "Deploy"
	DeployActionName string = "EcsFargateDeploy"
	DeployFilePath   string = "./deployment.zip"
	DeployBranch     string = "main"

	CodeBuildStartBuildAction string = "codebuild:StartBuild"
	PrivilegedBuildEnv        bool   = true
	UseLatestSdk              bool   = true
)

type PipelineNestedStackProps struct {
	awscdk.NestedStackProps
	PipelineName   string
	AppName        string
	FargateService ecs.FargateService
	TaskDefinition ecs.FargateTaskDefinition
	Repository     ecr.Repository
	Vpc            ec2.Vpc
	LoadBalancer   elb.ApplicationLoadBalancer
	Listener       elb.ApplicationListener
	TargetGroupB   elb.ApplicationTargetGroup
	TargetGroupG   elb.ApplicationTargetGroup
}

func NewDeployPipelineStack(scope constructs.Construct, id string, props *PipelineNestedStackProps) awscdk.Stack {
	var nsprops awscdk.NestedStackProps
	if props != nil {
		nsprops = props.NestedStackProps
	}
	stack := awscdk.NewNestedStack(scope, &id, &nsprops)

	// Creates an AWS CodeCommit repository
	codeRepo := NewRepository(stack, "CodeRepo", &RepositoryProps{
		Name:     props.AppName,
		FilePath: DeployFilePath,
		Branch:   DeployBranch,
	})

	// CodeBuild project that builds the initial Docker image when the stack is created
	buildImage := NewProject(stack, "BuildImage", &ProjectProps{
		BuildSpec: codebuild.BuildSpec_FromSourceFilename(jsii.String(BuildImageSpec)),
		Source: codebuild.Source_CodeCommit(&codebuild.CodeCommitSourceProps{
			Repository: codeRepo,
		}),
		CodeRepo:  codeRepo,
		ImageRepo: props.Repository,
		TaskDef:   props.TaskDefinition,
		BuildEnv: &codebuild.BuildEnvironment{
			Privileged: jsii.Bool(PrivilegedBuildEnv),
		},
		// Sets environment variables to use during the build
		Environment: &map[string]*codebuild.BuildEnvironmentVariable{
			"REGION": {
				Value: stack.Region(),
			},
			"AWS_ACCOUNT_ID": {
				Value: stack.Account(),
			},
			"IMAGE_TAG": {
				Value: jsii.String(BuildImageTag),
			},
			"IMAGE_REPO_NAME": {
				Value: props.Repository.RepositoryName(),
			},
			"REPOSITORY_URI": {
				Value: props.Repository.RepositoryUri(),
			},
			"TASK_ROLE_ARN": {
				Value: props.TaskDefinition.TaskRole().RoleArn(),
			},
			"EXECUTION_ROLE_ARN": {
				Value: props.TaskDefinition.TaskRole().RoleArn(),
			},
			"TASK_DEFINITION_ARN": {
				Value: props.TaskDefinition.TaskDefinitionArn(),
			},
		},
	})

	// Grants CodeBuild Project access to pull/push images from/to ECR repo
	props.Repository.GrantPullPush(buildImage)

	triggerCodeBuild := NewLambda(stack, "BuildLambda", &LambdaProps{
		Architecture: lambda.Architecture_ARM_64(),
		Code:         lambda.AssetCode_FromAsset(jsii.String(BuildLambdaAsset), nil),
		Handler:      BuildLambdaHandler,
		Runtime:      lambda.Runtime_NODEJS_18_X(),
		Environment: &map[string]*string{
			"REGION":                 stack.Region(),
			"CODEBUILD_PROJECT_NAME": buildImage.ProjectName(),
		},
		// Allows this Lambda function to trigger the buildImage CodeBuild project
		Policy: &[]iam.PolicyStatement{
			iam.NewPolicyStatement(&iam.PolicyStatementProps{
				Effect: iam.Effect_ALLOW,
				Actions: &[]*string{
					jsii.String(CodeBuildStartBuildAction),
				},
				Resources: &[]*string{
					buildImage.ProjectArn(),
				},
			}),
		},
	})

	triggerLambda := CustomLambdaTrigger(stack, "BuildLambdaTrigger", &LambdaTriggerProps{
		LambdaFunc:   triggerCodeBuild,
		UseLatestSdk: UseLatestSdk,
	})

	// Deploys the cluster VPC after the initial image build triggers
	props.Vpc.Node().AddDependency(triggerLambda)

	// Creates new pipeline artifacts
	sourceArtifact := pipeline.NewArtifact(jsii.String(SourceArtifactName))
	buildArtifact := pipeline.NewArtifact(jsii.String(BuildArtifactName))

	// Creates a new CodeDeploy Deployment Group
	deploymentGroup := NewDeploymentGroup(stack, "CodeDeployGroup", &DeploymentGroupProps{
		Service: props.FargateService,
		// Configurations for CodeDeploy Blue/Green deployments
		DeploymentConfig: &codedeploy.EcsBlueGreenDeploymentConfig{
			Listener:         props.Listener,
			BlueTargetGroup:  props.TargetGroupB,
			GreenTargetGroup: props.TargetGroupG,
		},
	})

	// Creates the source stage for CodePipeline
	sourceStage := SourceStage(sourceArtifact, codeRepo)

	// Creates the build stage for CodePipeline
	buildStage := BuildStage(buildArtifact, buildImage)

	// Creates the deploy stage for CodePipeline
	deployStage := DeployStage(buildArtifact, deploymentGroup)

	// Creates an AWS CodePipeline with source, build, and deploy stages
	pipeline.NewPipeline(stack, jsii.String(id), &pipeline.PipelineProps{
		PipelineName: jsii.String(props.PipelineName),
		Stages:       &[]*pipeline.StageProps{sourceStage, buildStage, deployStage},
	})

	return stack
}

type DeploymentGroupProps struct {
	Service          ecs.IBaseService
	DeploymentConfig *codedeploy.EcsBlueGreenDeploymentConfig
}

func NewDeploymentGroup(scope constructs.Construct, id string, props *DeploymentGroupProps) codedeploy.EcsDeploymentGroup {
	return codedeploy.NewEcsDeploymentGroup(scope, jsii.String(id), &codedeploy.EcsDeploymentGroupProps{
		Service:                   props.Service,
		BlueGreenDeploymentConfig: props.DeploymentConfig,
	})
}

type RepositoryProps struct {
	Name     string
	FilePath string
	Branch   string
}

func NewRepository(scope constructs.Construct, id string, props *RepositoryProps) codecommit.Repository {
	return codecommit.NewRepository(scope, jsii.String(id), &codecommit.RepositoryProps{
		RepositoryName: jsii.String(props.Name),
		Code: codecommit.Code_FromZipFile(
			jsii.String(props.FilePath),
			jsii.String(props.Branch),
		),
	})
}

type ProjectProps struct {
	BuildSpec   codebuild.BuildSpec
	Source      codebuild.ISource
	CodeRepo    codecommit.IRepository
	ImageRepo   ecr.IRepository
	TaskDef     ecs.FargateTaskDefinition
	BuildEnv    *codebuild.BuildEnvironment
	Environment *map[string]*codebuild.BuildEnvironmentVariable
}

func NewProject(scope constructs.Construct, id string, props *ProjectProps) codebuild.Project {
	return codebuild.NewProject(scope, jsii.String(id), &codebuild.ProjectProps{
		BuildSpec:            props.BuildSpec,
		Source:               props.Source,
		Environment:          props.BuildEnv,
		EnvironmentVariables: props.Environment,
	})
}

func SourceStage(source pipeline.Artifact, codeRepo codecommit.IRepository) *pipeline.StageProps {
	return &pipeline.StageProps{
		StageName: jsii.String(SourceStageName),
		Actions: &[]pipeline.IAction{
			pipelineactions.NewCodeCommitSourceAction(&pipelineactions.CodeCommitSourceActionProps{
				ActionName: jsii.String(SourceActionName),
				Branch:     jsii.String(DeployBranch),
				Output:     source,
				Repository: codeRepo,
			}),
		},
	}
}

func BuildStage(buildArtifact pipeline.Artifact, buildImage codebuild.IProject) *pipeline.StageProps {
	return &pipeline.StageProps{
		StageName: jsii.String(BuildStageName),
		Actions: &[]pipeline.IAction{
			pipelineactions.NewCodeBuildAction(&pipelineactions.CodeBuildActionProps{
				ActionName: jsii.String(BuildActionName),
				Input:      pipeline.NewArtifact(jsii.String(SourceArtifactName)),
				Project:    buildImage,
				Outputs: &[]pipeline.Artifact{
					buildArtifact,
				},
			}),
		},
	}
}

func DeployStage(buildArtifact pipeline.Artifact, deploymentGroup codedeploy.IEcsDeploymentGroup) *pipeline.StageProps {
	return &pipeline.StageProps{
		StageName: jsii.String(DeployStageName),
		Actions: &[]pipeline.IAction{
			pipelineactions.NewCodeDeployEcsDeployAction(&pipelineactions.CodeDeployEcsDeployActionProps{
				ActionName:                  jsii.String(DeployActionName),
				AppSpecTemplateInput:        buildArtifact,
				DeploymentGroup:             deploymentGroup,
				TaskDefinitionTemplateInput: buildArtifact,
			}),
		},
	}
}
