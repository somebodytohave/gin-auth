# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOMOD=$(GOCMD) mod
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOLINT=$(GOCMD)lint
BINARY_NAME=mybinary
BINARY_UNIX=$(BINARY_NAME)_unix

.PHONY: all build pull download tool lint clean help

all: build

build:
	$(GOBUILD) -v .

pull:
	git pull
	cp app.ini conf/app.ini

download:
	$(GOMOD) download

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
	@echo "make pull: pull project from github and cp app.ini"
	@echo "make download: download all go.mod packages"
	@echo "make tool: run specified go tool"
	@echo "make lint: golint ./..."
	@echo "make clean: remove object files and cached files"