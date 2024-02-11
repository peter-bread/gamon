.PHONY: build test clean release

# The name of your binary
BINARY_NAME=gam

# The output directory
OUTPUT_DIR=bin

# The Go binary path
export GOBIN=$(shell which go)

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

uninstall:
	@rm -f /usr/local/bin/$(BINARY_NAME) || echo "Failed to remove binary. Try running with sudo."