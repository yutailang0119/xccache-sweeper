NAME     := xccache-sweeper
VERSION  := $(shell git describe --tags --abbrev=0)
REVISION := $(shell git rev-parse --short HEAD)

GO = go
GOPATH := $(shell go env GOPATH)

VERBOSE_FLAG = $(if $(VERBOSE),-v)

BUILD_FLAGS = -ldflags "\
	      -X \"main.Version=$(VERSION)\" \
	      -X \"main.Revision=$(REVISION)\" \
	      "

PREFIX?= /usr/local

build: deps
	$(GO) build $(VERBOSE_FLAG) $(BUILD_FLAGS)

test: testdeps
	$(GO) test $(VERBOSE_FLAG) $($(GO) list ./... | grep -v '^github.com/yutailang0119/xccache-sweeper/vendor/')

deps:
	$(GO) get -d $(VERBOSE_FLAG)

testdeps:
	$(GO) get -d -t $(VERBOSE_FLAG)

go-install: deps
	$(GO) install $(VERBOSE_FLAG) $(BUILD_FLAGS)

install: deps
	mkdir -p "$(PREFIX)/bin"
	cp -f "./$(NAME)" "$(PREFIX)/bin/$(NAME)"

bump-minor:
	git diff --quiet && git diff --cached --quiet
	new_version=$$(gobump minor -w -r -v) && \
	test -n "$$new_version" && \
	git commit -a -m "bump version to $$new_version" && \
	git tag v$$new_version

.PHONY: build test deps testdeps install hoge
