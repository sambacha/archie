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
COVERAGE_RESULTS=coverage.out

.PHONY: all build clean install uninstall fmt simplify check run lint test coverage

all: check install

$(TARGET): $(SRC)
	@go build $(LDFLAGS) -o $(TARGET)

build: $(TARGET)
	@true

clean:
	@rm -f $(TARGET)

install:
	@go install $(LDFLAGS)

uninstall: clean
	@rm -f $$(which ${TARGET})

fmt:
	@gofmt -l -w $(SRC)

simplify:
	@gofmt -s -l -w $(SRC)

check:
	@echo "[Formatting]"
	@test -z $(shell gofmt -l main.go | tee /dev/stderr) || echo "[WARN] Fix formatting issues with 'make fmt'"
	@echo "[Vetting]"
	@for d in $$(go list ./... | grep -v /vendor/); do go vet $${d}; done
	@echo "[Linting]"
	@for d in $$(go list ./... | grep -v /vendor/); do golint $${d}; done

test:
	@go test $(TEST_FLAGS) ./...

coverage: TEST_FLAGS+= -coverprofile=$(COVERAGE_RESULTS)
coverage: test
	@go tool cover -html=coverage.out

run: install
	@$(TARGET)

print-%  : ; @echo $* = $($*)