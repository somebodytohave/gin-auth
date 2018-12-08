# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOGET=$(GOCMD) get
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOLINT=$(GOCMD)lint
BINARY_NAME=mybinary
BINARY_UNIX=$(BINARY_NAME)_unix

.PHONY: build get clean tool lint help


all: build

build:
	$(GOBUILD) -v .

get:
	$(GOGET) ./... -v

tool:
	go tool vet . |& grep -v vendor \
	gofmt -w .

lint:
	$(GOLINT) ./...

clean:
	rm -rf gin-auth
	$(GOCLEAN) -i .

help:
	@echo "make: compile packages and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make lint: golint ./..."
	@echo "make clean: remove object files and cached files"