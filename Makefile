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
