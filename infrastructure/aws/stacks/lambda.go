package stacks

import (
	iam "github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	lambda "github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	custom "github.com/aws/aws-cdk-go/awscdk/v2/customresources"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

const (
	InvokeAction       string = "invoke"
	InvokeLambdaAction string = "lambda:InvokeFunction"

	LambdaService string = "Lambda"

	PhysicalResourceId string = "id"

	EventInvocation string = "Event"
)

type LambdaProps struct {
	Architecture lambda.Architecture
	Code         lambda.Code
	Handler      string
	Runtime      lambda.Runtime
	Environment  *map[string]*string
	Policy       *[]iam.PolicyStatement
}

func NewLambda(scope constructs.Construct, id string, props *LambdaProps) lambda.Function {
	return lambda.NewFunction(scope, jsii.String(id), &lambda.FunctionProps{
		Architecture:  props.Architecture,
		Code:          props.Code,
		Handler:       jsii.String(props.Handler),
		Runtime:       props.Runtime,
		Environment:   props.Environment,
		InitialPolicy: props.Policy,
	})
}

type LambdaTriggerProps struct {
	LambdaFunc   lambda.Function
	UseLatestSdk bool
}

// CustomLambdaTrigger triggers a Lambda function using AWS SDK
func CustomLambdaTrigger(scope constructs.Construct, id string, props *LambdaTriggerProps) custom.AwsCustomResource {
	return custom.NewAwsCustomResource(scope, jsii.String(id), &custom.AwsCustomResourceProps{
		InstallLatestAwsSdk: jsii.Bool(props.UseLatestSdk),
		Policy: custom.AwsCustomResourcePolicy_FromStatements(&[]iam.PolicyStatement{
			iam.NewPolicyStatement(&iam.PolicyStatementProps{
				Actions: &[]*string{
					jsii.String(InvokeLambdaAction),
				},
				Effect: iam.Effect_ALLOW,
				Resources: &[]*string{
					props.LambdaFunc.FunctionArn(),
				},
			}),
		}),
		OnCreate: &custom.AwsSdkCall{
			Service:            jsii.String(LambdaService),
			Action:             jsii.String(InvokeAction),
			PhysicalResourceId: custom.PhysicalResourceId_Of(jsii.String(PhysicalResourceId)),
			Parameters: map[string]*string{
				"FunctionName":   props.LambdaFunc.FunctionName(),
				"InvocationType": jsii.String(EventInvocation),
			},
		},
		OnUpdate: &custom.AwsSdkCall{
			Service: jsii.String(LambdaService),
			Action:  jsii.String(InvokeAction),
			Parameters: map[string]*string{
				"FunctionName":   props.LambdaFunc.FunctionName(),
				"InvocationType": jsii.String(EventInvocation),
			},
		},
	})
}
