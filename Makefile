NAME = $(notdir $(PWD))

VERSION = $(shell printf "%s.%s" \
		$$(git rev-list --count HEAD) \
		$$(git rev-parse --short HEAD))

BRANCH = $(shell git rev-parse --abbrev-ref HEAD)

generate:
	@echo :: getting generator
	go get -v -d
	go get -d github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.4.2

	@echo :: generating code

	oapi-codegen -package v1 -generate chi-server,types api/v1/v1.yml > api/v1/v1.gen.go
	mkdir -p api/client/v1
	oapi-codegen -package v1 -generate client,types api/v1/v1.yml > api/client/v1/client.gen.go

build:  $(OUTPUT)
	CGO_ENABLED=0 GOOS=linux go build -o bin/app \
		-ldflags "-X main.version=$(VERSION)" \
		-gcflags "-trimpath $(GOPATH)/src"

all: generate build

$(OUTPUT):
	mkdir -p $(OUTPUT)
