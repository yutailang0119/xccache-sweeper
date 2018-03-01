NAME     := go-xccache-sweeper
VERSION  := $(shell git describe --tags --abbrev=0)
REVISION := $(shell git rev-parse --short HEAD)

GOPATH   := $(shell go env GOPATH)
SRCS     := $(shell find . -type f -name '*.go')
LDFLAGS  := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\" -extldflags \"-static\""

PREFIX?= /usr/local

## Setup
setup:
	#go get github.com/Masterminds/glide

## Install dependencies
deps: setup
	#glide install

## Update dependencies
update: setup
	#glide update
# Build
build: setup
	go build $(LDFLAGS)
	
## Install application
install: build
	mkdir -p "$(PREFIX)/bin"
	cp -f "./$(NAME)" "$(PREFIX)/bin/$(NAME)"

## Test
test:
	go test -cover -v `glide novendor`

## Clean
clean:
	rm -rf bin/*
	rm -rf vendor/*

.PHONY: setup deps update build install test clean

