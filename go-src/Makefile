all: help

help:     ## Show this help.
	@echo 'Usage: make [TARGET]'
	@echo 'Targets:'
	@grep '^[a-zA-Z]' $(MAKEFILE_LIST) | awk -F ':.*?## ' 'NF==2 {printf "\033[36m  %-25s\033[0m %s\n", $$1, $$2}'

test:	  		## Test and Code Coverage 
	go test ./... -cover

build:	  		## Build Docker image.
	docker build -t go-web .

run:			## Run Go application (available on port 8080).
	docker run -itd --name go-web -p 8080:8080 go-web

shell:    ## Shell into Docker image.
	docker run -it go-web /bin/sh

clean:			## Removes any transient build artifacts.
	docker image rm go-web:latest -f
	rm main

tag:			## Tag the current version as 'latest'
	docker tag go-web:latest colinschofield/go-web:latest

push:			## push the tagged version to the Docker Hub
	docker push colinschofield/go-web:latest

refresh:		## Initiate a rolling update of the Kubernetes managed cluster, reloading the current image from the Docker Hub
	kubectl rollout restart deployment
	kubectl get events
	kubectl get pods -w -o wide

update:  build tag push refresh			## Push and load the current version to the Kubernetes managed cluster