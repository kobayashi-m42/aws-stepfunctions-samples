ENV        := dev
STACK_NAME := $(ENV)-aws-stepfunctions-sync-samples

.PHONY: deploy
deploy:
	aws cloudformation deploy \
		--stack-name $(STACK_NAME) \
		--parameter-overrides `cat parameters.ini | tr '\n' ' '` \
		--template-file cloudformation/main.yaml

delete-stack:
	aws cloudformation delete-stack \
		--stack-name $(STACK_NAME)
