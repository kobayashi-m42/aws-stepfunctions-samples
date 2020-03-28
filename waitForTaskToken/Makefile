ENV        := dev
STACK_NAME := $(ENV)-aws-stepfunctions-token-samples
APP_NAME   := $(ENV)-stepfunctions-wait-token
REGISTRY   := $(AWS_ACCOUNT_ID).dkr.ecr.ap-northeast-1.amazonaws.com

.PHONY: push
push:
	aws ecr get-login-password | docker login --username AWS --password-stdin $(REGISTRY)/$(APP_NAME)
	docker build -t $(REGISTRY)/$(APP_NAME):latest -f Dockerfile .
	docker push $(REGISTRY)/$(APP_NAME):latest

.PHONY: deploy
deploy:
	aws cloudformation deploy \
		--stack-name $(STACK_NAME) \
		--parameter-overrides `cat parameters.ini | tr '\n' ' '` \
		--template-file cloudformation/main.yaml

.PHONY: delete-stack
delete-stack:
	aws cloudformation delete-stack \
		--stack-name $(STACK_NAME)
