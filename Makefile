SHELL := /bin/bash

# The name of the executable (default is current directory name)
TARGET := $(shell echo $${PWD\#\#*/})
.DEFAULT_GOAL: $(TARGET)

# These will be provided to the target
VERSION := 1.0.0
BUILD := `git rev-parse HEAD`

# Use linker flags to provide version/build settings to the target
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

# go source files, ignore vendor directory
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

# Testing flags
TEST_FLAGS=-v

# Format flags
FMT_FLAGS=-l -e -s

.PHONY: fmt vet lint run test coverage docs

fmt:
	@gofmt -w $(FMT_FLAGS) $(SRC)

check:
	@gofmt -d $(FMT_FLAGS) $(SRC)

vet:
	@for d in $$(go list ./... | grep -v /vendor/); do go vet $${d}; done

lint:
	@for d in $$(go list ./... | grep -v /vendor/); do golint $${d}; done

coverage: TEST_FLAGS+= -covermode=count -coverprofile=coverage.out
coverage: test
	$$GOPATH/bin/goveralls -coverprofile=coverage.out -service=circle-ci -repotoken=$$COVERALLS_TOKEN

test:
	go test $(TEST_FLAGS) ./...

download:
	@echo Download go.mod dependencies
	@go mod download

install-tools: download
	@echo Installing tools from tools.go
	@cat tools/tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

docs:
	@echo Creating output directory
	@rm -rf ./docs/bin
	@echo Copying content
	@cp -r ./docs/content ./docs/bin
	@echo Generating content
	@mkdir -p ./docs/bin/cli
	@go run ./docs/main.go ./docs/bin/cli


print-%  : ; @echo $* = $($*)
