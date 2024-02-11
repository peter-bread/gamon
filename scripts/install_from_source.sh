#!/bin/bash

# Clone gamon repository
git clone https://github.com/peter-bread/gamon.git

# Naviage to the gamon directory
cd gamon

# Ensure you are on the main branch
git checkout main

# Build the tool
make build

# Create `/usr/local/bin` if it doesn't already exist
sudo mkdir -p /usr/local/bin

# Copy binary from repository to `/usr/local/bin`
sudo mv ./bin/gam /usr/local/bin
