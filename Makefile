# Set an output prefix, which is the local directory if not specified
PREFIX?=$(shell pwd)
BUILDTAGS=

.PHONY:  all fmt vet lint build test
.DEFAULT: default

all: build fmt lint test cover vet

build:
	@echo "+ $@"
	@go build -tags "$(BUILDTAGS) cgo" .

fmt:
	@echo "+ $@"
	@gofmt -s -l . | grep -v vendor | tee /dev/stderr

lint:
	@echo "+ $@"
	@golint ./... | grep -v vendor | tee /dev/stderr

test: fmt lint vet
	@echo "+ $@"
	@go test -v -tags "$(BUILDTAGS) cgo" $(shell go list ./... | grep -v vendor)

cover:
	@echo "+ $@"
	@go test -cover $(shell go list ./... | grep -v vendor)
vet:
	@echo "+ $@"
	@go vet $(shell go list ./... | grep -v vendor)