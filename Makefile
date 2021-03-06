GIT_VERSION = $(shell git describe --tags --dirty)
VERSION ?= $(GIT_VERSION)

GO ?= go
GOVERSIONS ?= go1.9 go1.10 go1.11
OS := $(shell uname)
SHELL := /bin/bash

BASE = $(GOPATH)/src/github.com/mahakamcloud/mahakam

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: cli
cli: fmt vet ## Build mahakam cli binary
	./hack/build-bin.sh mahakam cmd/mahakam

.PHONY: server-darwin
server-darwin: fmt vet ## Build mahakam server binary
	./hack/build-bin.sh mahakam-server cmd/mahakam_server darwin

.PHONY: server-linux
server-linux: fmt vet ## Build mahakam server binary
	./hack/build-bin.sh mahakam-server cmd/mahakam_server linux

.PHONY: server-image
server-image: server-linux
	./build/run.sh mahakam-server

.PHONY: fmt
fmt: ## Run go fmt against code
	go fmt ./pkg/... ./cmd/...

.PHONY: vet
vet: ## Run go vet against code
	go vet ./pkg/... ./cmd/...

.PHONY: test
test: ## run tests
	@echo running tests...
	$(GO) test -v $(shell go list -v ./... | grep -v /vendor/ | grep -v integration | grep -v /playground )

.PHONY: dev
dev: generate-server server-linux ## run dev server and consul
	@echo running dev server and consul...
	docker-compose -f docker-compose.dev.yaml up --build

.PHONY: dev-server
dev-server: ## run dev server
	@echo running dev server...
	$(GO) run $(BASE)/cmd/mahakam_server/main.go --config ./config.dev.yaml

.PHONY: dev-store
dev-store: ## run dev store with consul backend
	@echo running dev store...
	docker run -d --name=dev-consul -e CONSUL_BIND_INTERFACE=eth0 -p 8500:8500 -p 8600:8600 consul

.PHONY: dev-docker
dev-docker: ## run dev docker that has terraform libvirt plugin and golang
	@echo running dev docker container...
	docker run -it --rm -p 9000:9000 -v $(HOME)/.ssh:/root/.ssh -v $(PWD):/root/go/src/github.com/mahakamcloud/mahakam -v $(HOME)/.aws:/root/.aws -w /root/go/src/github.com/mahakamcloud/mahakam devrunner:latest /bin/bash

.PHONY: staging-docker
staging-docker: ## run staging docker that has terraform libvirt plugin and golang with privileged access
	@echo running staging docker container...
	docker run -it --rm --network host --privileged -v $(PWD):/root/go/src/github.com/mahakamcloud/mahakam -v $(HOME)/.aws:/root/.aws -w /root/go/src/github.com/mahakamcloud/mahakam devrunner:latest /bin/bash

.PHONY: generate
generate: generate-server generate-client ## Generate both swagger server and client

.PHONY: generate-server
generate-server: ## Generate swagger server
	cd $(BASE)/pkg/api/v1 && swagger generate server \
		-A mahakam -f $(BASE)/swagger/mahakam.yaml --exclude-main --skip-validation

.PHONY: generate-client
generate-client: ## Generate swagger client
	cd $(BASE)/pkg/api/v1 && swagger generate client \
		-A mahakam -f $(BASE)/swagger/mahakam.yaml -c client --skip-validation

.PHONY: build-ci-runner
build-ci-runner:
	docker build -f Dockerfile.ci -t builder:latest .

.PHONY: build-dev-runner
build-dev-runner:
	docker build -f Dockerfile.dev -t devrunner:latest .
