GOCMD=go
GOBIN=$(GOPATH)/bin
BUILD_DIR=$(HOME)/oshbuilds
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=stripe
MODULE_NAME=github.com/appno/stripe
GOIMPORTS=goimports
GOFMT=$(GOCMD) fmt

all: build test

.PHONY: all format run test

build:
		$(KUBE_GEN)
		$(GOBUILD) -o $(GOBIN)/$(BINARY_NAME) -tags internal
run:
		$(GOBUILD) -o $(GOBIN)/$(BINARY_NAME)
		$(GOBIN)/$(BINARY_NAME)
clean:
		$(GOCLEAN)
		rm -f $(GOBIN)/$(BINARY_NAME)
test:
		$(GOTEST) -v ./...
format:
		$(OSH_FMT)
