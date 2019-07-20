SHELL := /bin/sh

GO := go

.POSIX:
.PHONY: clean all deps build test

all: deps build test

build:
	${GO} build -v ./...

deps:
	go mod download
	go mod vendor

clean:
	rm -rf vendor
	${GO} clean -v ./...

test: deps build
	${GO} test -v ./...
