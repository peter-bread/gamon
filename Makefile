.PHONY: build test clean uninstall

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

_start_install: build
	@mkdir -p ~/.gamon/bin
	@mv $(OUTPUT_DIR)/$(BINARY_NAME) ~/.gamon/bin/$(BINARY_NAME)
	@cp ./README.md ~/.gamon/README.md
	@cp ./LICENSE ~/.gamon/LICENSE

install: _start_install clean
	@echo "Installation completed successfully."
	@echo ""
	@echo "Please add the following line to your .bashrc or .zshrc:"
	@echo ""
	@echo '    export PATH="$$HOME/.gamon/bin:$$PATH"'

uninstall:
	@rm -f ~/.gamon/bin/$(BINARY_NAME) || echo "Failed to remove binary."