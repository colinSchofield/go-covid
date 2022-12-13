all: help

.PHONY: help
help:     		## Show this help
	@echo 'Usage: make [TARGET]'
	@echo 'Targets:'
	@grep '^[a-zA-Z]' $(MAKEFILE_LIST) | awk -F ':.*?## ' 'NF==2 {printf "\033[36m  %-25s\033[0m %s\n", $$1, $$2}'

.PHONY: clean
clean:              ## Removes any transient build artifacts
	@rm -f deployment/serverless/go-covid-api
	@cd deployment/serverless && serverless remove

.PHONY: build
build:              ## Compile Go binary for Rest API
	@cd go-src && env GOOS=linux go build -o ../deployment/serverless/go-covid-api .

.PHONY: deploy
deploy:	build		## Deploy the Lambda Rest API to AWS CloudFormation
	@cd deployment/serverless && serverless deploy