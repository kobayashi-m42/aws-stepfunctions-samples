ENV        := dev
STACK_NAME := $(ENV)-aws-stepfunctions-sync-samples
TASK_ROLE_ARN := ''
APP_IMAGE := ''

.PHONY: deploy
deploy:
	aws cloudformation deploy \
		--stack-name $(STACK_NAME) \
		--parameter-overrides TaskRoleArn=${TASK_ROLE_ARN} AppImage=${APP_IMAGE} \
		--template-file cloudformation/main.yaml \
		--capabilities CAPABILITY_IAM

delete-stack:
	aws cloudformation delete-stack \
		--stack-name $(STACK_NAME)
