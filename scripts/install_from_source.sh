#!/bin/bash

# Check if git is installed
if ! command -v git &>/dev/null; then
  echo "Error: git is not installed"
  exit 1
fi

# Check if make is installed
if ! command -v make &>/dev/null; then
  echo "Error: make is not installed"
  exit 1
fi

# Function to display help message
display_help() {
  echo "Usage: $0 [version]"
  echo
  echo "Install a specific version or the latest version of the software."
  echo
  echo "If version is specified, the script will attempt to install that version."
  echo "If version is not specified, the script will install the latest version."
  echo
  echo "Examples:"
  echo "  $0 1       Install latest version of major version 1"
  echo "  $0         Install the latest version"
  exit 1
}

# If the first argument is -h or --help, display the help message
if [[ $1 = "-h" || $1 = "--help" ]]; then
  display_help
fi

# If an argument is passed, check if it is a number
if [[ -n $1 && ! $1 =~ ^[0-9]+$ ]]; then
  echo "Error: Version must be a number"
  exit 1
fi

MAJOR_VERSION=$1

# Clone gamon repository
if ! git clone --depth 1 https://github.com/peter-bread/gamon.git; then
  echo "Error: Failed to clone repository"
  rm -rf gamon # Cleanup partially populated directory
  exit 1
fi

# Naviage to the gamon directory
cd gamon || exit 1

# Fetch all tags
git fetch --tags

# If MAJOR_VERSION is not empty, get the latest tag that matches the major version
# Otherwise, get the latest tag
if [[ -n $MAJOR_VERSION ]]; then
  LATEST_TAG=$(git tag -l "v${MAJOR_VERSION}.*" | sort -V | tail -n 1)
else
  LATEST_TAG=$(git describe --tags "$(git rev-list --tags --max-count=1)")
fi

# If LATEST_TAG is empty, print an error message and exit
if [[ -z $LATEST_TAG ]]; then
  echo "Error: No tags found that match the major version: v${MAJOR_VERSION}"
  exit 1
fi

# Get the latest tag that matches the major version
LATEST_TAG=$(git tag -l "v${MAJOR_VERSION}.*" | sort -V | tail -n 1)

# Checkout the latest tag
git checkout "$LATEST_TAG"

# Build and install
if ! make install; then
  echo "Error: Failed to install"
  exit 1
fi
