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
	swag init
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
	@echo "make download: download all packages from go.mod"
	@echo "make tool: run specified go tool"
	@echo "make lint: golint ./..."
	@echo "make clean: remove object files and cached files"


{
  "server":"74.82.202.168",             #你的 ss 服务器 ip
  "server_port":8388,                #你的 ss 服务器端口
  "local_address": "127.0.0.1",   #本地ip
  "local_port":1080,                 #本地端口
  "password":"password",          #连接 ss 密码
  "timeout":300,                  #等待超时
  "method":"aes-256-cfb",         #加密方式
  "workers": 1                    #工作线程数
}