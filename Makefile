# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOLINT=golangci-lint run
BIN_DIR=bin

# Binary name
BINARY_NAME=playgrnd

# Docker parameters
DOCKER_BUILD_CMD=docker build
DOCKER_IMAGE_NAME=enkya-playground

all:  help

help:
	@echo "make lint - Lint the project"
	@echo "make build - Build the project"
	@echo "make test - Test the project"
	@echo "make clean - Clean the project"
	@echo "make docker-build - Build Docker image"

lint:
	$(GOLINT) ./...

lint-fix:
	$(GOLINT) --fix ./...

build:
	$(GOBUILD) -o $(BIN_DIR)/$(BINARY_NAME) -v

run: clean build
	./$(BIN_DIR)/$(BINARY_NAME)

test:
	$(GOTEST) -v ./...

clean:
	rm -f $(BIN_DIR)/$(BINARY_NAME)

docker-build:
	$(DOCKER_BUILD_CMD) -t $(DOCKER_IMAGE_NAME) .

docker-run:
	docker run -p 8080:8080 $(DOCKER_IMAGE_NAME)

docker-push:
	docker push $(DOCKER_IMAGE_NAME)

docker-clean:
	docker rmi $(DOCKER_IMAGE_NAME)

docker-all: docker-build docker-run

.PHONY: all help lint build test clean docker-build docker-run docker-push docker-clean docker-all

.lint-setup:
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b $(shell go env GOPATH)/bin v1.27.0

.lint-setup-mac:
	brew install golangci/tap/golangci-lint

.docker-setup-mac:
	brew install docker

.docker-setup-linux:
	sudo apt-get install docker-ce docker-ce-cli containerd.io

.setup-live-reload:
	go get github.com/cosmtrek/air@latest
	go install github.com/cosmtrek/air@latest

.setup-python-virtualenv:
	python3 -m venv /pie/.venv
	source /pie/.venv/bin/activate

.setup-python-dependencies:
	pip install -r pie/requirements.txt

.setup-python: .setup-python-virtualenv .setup-python-dependencies

.setup: .lint-setup .setup-live-reload .setup-python
