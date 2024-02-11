.PHONY: build test clean release

# The name of your binary
BINARY_NAME=gam

# The output directory
OUTPUT_DIR=bin

# The Go binary path
GOBIN=$(shell which go)

# The Go build command
GOBUILD=$(GOBIN) build

# The Go test command
GOTEST=$(GOBIN) test

# The Go clean command
GOCLEAN=$(GOBIN) clean

build:
	mkdir -p $(OUTPUT_DIR)
	$(GOBUILD) -o $(OUTPUT_DIR)/$(BINARY_NAME) -v

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -rf $(OUTPUT_DIR)

build-binaries:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(OUTPUT_DIR)/$(BINARY_NAME)-darwin-amd64 -v
	GOOS=darwin GOARCH=arm64 $(GOBUILD) -o $(OUTPUT_DIR)/$(BINARY_NAME)-darwin-arm64 -v
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(OUTPUT_DIR)/$(BINARY_NAME)-linux-amd64 -v
	GOOS=linux GOARCH=arm64 $(GOBUILD) -o $(OUTPUT_DIR)/$(BINARY_NAME)-linux-arm64 -v

install: build
	@mkdir -p /usr/local/bin || true
	@cp $(OUTPUT_DIR)/$(BINARY_NAME) /usr/local/bin || echo "Failed to copy binary. Try running with sudo."

uninstall:
	@rm -f /usr/local/bin/$(BINARY_NAME) || echo "Failed to remove binary. Try running with sudo."