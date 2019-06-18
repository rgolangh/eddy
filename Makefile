# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GODEP=dep

PREFIX=.

VERSION?=$(shell git describe --tags --always --match "v[0-9]*" | awk -F'-' '{print substr($$1,2) }')
RELEASE?=$(shell git describe --tags --always --match "v[0-9]*" | awk -F'-' '{if ($$2 != "") {print $$2 "." $$3} else {print 1}}')
VERSION_RELEASE=$(VERSION)$(if $(RELEASE),-$(RELEASE))

COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

COMMON_ENV=CGO_ENABLED=0 GOOS=linux GOARCH=amd64
COMMON_GO_BUILD_FLAGS=-ldflags '-extldflags "-static"'

TARBALL=eddy-$(VERSION_RELEASE).tar.gz

all: clean deps build test

binaries = \
	eddy

$(binaries): vet
	$(COMMON_ENV) $(GOBUILD) \
    	$(COMMON_GO_BUILD_FLAGS) \
    	-o $(PREFIX)/$@ \
    	-v $@.go

vet:
	go vet ./...

build: $(binaries)

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	git clean -df

deps:
	dep ensure --update

tarball: $(TARBALL)

$(TARBALL):
	/bin/git archive --format=tar.gz HEAD > $(TARBALL)

.PHONY: all tarball test build build-containers push-containers apb_build apb_docker_push apb_push
