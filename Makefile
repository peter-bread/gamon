.PHONY: build test clean release

# The name of your binary
BINARY_NAME=gam

# The output directory
OUTPUT_DIR=bin

# The Go path
GOPATH=$(shell go env GOPATH)

# The Go build command
GOBUILD=go build

# The Go test command
GOTEST=go test

# The Go clean command
GOCLEAN=go clean

build:
	mkdir -p $(OUTPUT_DIR)
	$(GOBUILD) -o $(OUTPUT_DIR)/$(BINARY_NAME) -v

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -rf $(OUTPUT_DIR)

release:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(OUTPUT_DIR)/$(BINARY_NAME)-darwin-amd64 -v
	GOOS=darwin GOARCH=arm64 $(GOBUILD) -o $(OUTPUT_DIR)/$(BINARY_NAME)-darwin-arm64 -v
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(OUTPUT_DIR)/$(BINARY_NAME)-linux-amd64 -v
	GOOS=linux GOARCH=arm64 $(GOBUILD) -o $(OUTPUT_DIR)/$(BINARY_NAME)-linux-arm64 -v