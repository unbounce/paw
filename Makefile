.DEFAULT_GOAL := help
.PHONY: help build test require-profile deploy-iam-stack deploy-code-stack prepare-code-stack deploy-code check-iam-stack check-code-stack

pwd := $(shell pwd)

git_hash := $(shell git rev-parse --short head)

project.name := paw
project.repo := github.com/unbounce/$(project.name)

build.dir := $(pwd)/.build
build.filename := main
build.file := $(build.dir)/$(build.filename)

dist.dir := $(pwd)/.dist
dist.filename := artifact.zip
dist.file := $(dist.dir)/$(dist.filename)

aws.region := us-east-1

iam.template.file := $(pwd)/iam.cft
iam.export.name := "paw:iam:role:arn"

code.template.file := $(pwd)/code.cft
lambda.export.name := "paw:lambda:function:arn"

cfn.tags := Key=project,Value=$(project.name) Key=lifetime,Value=long Key=repository,Value=$(project.repo)

help: ## Shows this message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Compiles the code for the Linux target (AWS Lambda)
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main.VERSION=$(git_hash)" -o $(build.file) $(project.repo)

test: ## Runs tests locally
	go test $(project.repo)

dist: ## Packages the file for upload to AWS Lambda
	mkdir -p $(dist.dir)
	cd $(build.dir) && zip $(dist.file) $(build.filename)

clean: ## Removes any project-generated files/directories
	rm -rf $(build.dir)
	rm -rf $(dist.dir)

require-profile:
ifndef AWS_PROFILE
	@echo "Please specify an AWS_PROFILE to continue."; exit 1
endif

deploy-iam-stack: require-profile ## Deploys the IAM resource stack via CFN
	aws cloudformation create-stack --stack-name $(project.name)-iam --region $(aws.region) --template-body file://$(iam.template.file) --enable-termination-protection --tags $(cfn.tags) --profile $(AWS_PROFILE) --capabilities CAPABILITY_IAM

deploy-code-stack: require-profile ## Deploys the code resource stack via CFN
	aws cloudformation create-stack --stack-name $(project.name)-code --region $(aws.region) --template-body file://$(code.template.file) --enable-termination-protection --tags $(cfn.tags) --profile $(AWS_PROFILE) --parameters ParameterKey=IamRoleArn,ParameterValue=$(shell aws cloudformation list-exports --region $(aws.region) --query 'Exports[?Name == `$(iam.export.name)`].Value' --profile $(AWS_PROFILE) --output text)

prepare-code-stack: require-profile ## Prepares the code stack to run the code
	aws lambda update-function-configuration --function-name $(shell aws cloudformation list-exports --region $(aws.region) --query 'Exports[?Name == `$(lambda.export.name)`].Value' --profile $(AWS_PROFILE) --output text) --region $(aws.region) --handler main --runtime go1.x --profile $(AWS_PROFILE)

deploy-code: require-profile ## Deploys the code to the Lambda function
	aws lambda update-function-code --function-name $(shell aws cloudformation list-exports --region $(aws.region) --query 'Exports[?Name == `$(lambda.export.name)`].Value' --profile $(AWS_PROFILE) --output text) --region $(aws.region) --profile $(AWS_PROFILE) --zip-file fileb://$(dist.file)

check-iam-stack: require-profile ## Displays the last status from the IAM stack
	aws cloudformation describe-stack-events --stack-name $(project.name)-iam --region $(aws.region) --profile $(AWS_PROFILE) --query 'StackEvents[0]'

check-code-stack: require-profile ## Displays the last status from the code stack
	aws cloudformation describe-stack-events --stack-name $(project.name)-code --region $(aws.region) --profile $(AWS_PROFILE) --query 'StackEvents[0]'

init-ssm: require-profile
ifndef SLACK_WEBHOOK_URL
	@echo "Please specify a SLACK_WEBHOOK_URL to continue."; exit 1
endif
	aws ssm put-parameter --name '/$(project.name)/slack/incoming-webhook/url' --description 'Incoming webhook URL for Slack' --type String --value '$(SLACK_WEBHOOK_URL)' --region $(aws.region) --profile $(AWS_PROFILE)

