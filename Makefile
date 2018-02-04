NAME     := go-xccache-sweeper
VERSION  := $(shell git --tags --abbrev=0)
REVISION := $(shell git rev-parse --short HEAD)

SRCS    := $(shell find . -type f -name '*.go')
LDFLAGS := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\" -extldflags \"-static\""

## Setup
setup:
	#go get github.com/Masterminds/glide

## Install dependencies
deps: setup
	#glide install

## Update dependencies
update: setup
	#glide update

## Install application
install:
	go install $(LDFLAGS)

.PHONY: setup deps update install
