# Set the shell to bash always
SHELL := /bin/bash

build-example: 
	docker build . -f example/Dockerfile -t hasheddan/veneer-example:latest

kind-example: build-example
	docker tag hasheddan/veneer-example:latest hasheddan/veneer-example:local
	$(KIND) load docker-image hasheddan/veneer-example:local
	$(KUBECTL) apply -f example/job.yaml

# ====================================================================================
# Tools

KIND=$(shell which kind)
KUBECTL=$(shell which kubectl)

.PHONY: build-example kind-example