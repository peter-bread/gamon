#!/bin/bash

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
  echo "  $0 1.2.3   Install version 1.2.3"
  echo "  $0         Install the latest version"
  exit 1
}

# Function to clean up partially downloaded or extracted files
cleanup() {
  if [ $? -ne 0 ]; then
    echo "Installation failed. Cleaning up..."
    rm -rf ~/.gamon
  fi
}

# If the script exits for any reason, run the cleanup function
trap cleanup EXIT

# If the first argument is -h or --help, display the help message
if [ "$1" = "-h" ] || [ "$1" = "--help" ]; then
  display_help
fi

# If a version is specified, validate the version number format
if [ -n "$1" ]; then
  if [[ ! $1 =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    echo "Error: Invalid version number format. Expected format is X.Y.Z"
    exit 1
  fi
fi

# Check if curl, wget, and tar are installed
if ! command -v curl >/dev/null; then
  echo "Error: curl is not installed."
  exit 1
fi

if ! command -v wget >/dev/null; then
  echo "Error: wget is not installed."
  exit 1
fi

if ! command -v tar >/dev/null; then
  echo "Error: tar is not installed."
  exit 1
fi

# Get OS and CPU information
OS=$(uname -s)
CPU=$(uname -m)

# Convert OS and CPU information to match the format used in the download URLs
case "$OS" in
"Darwin") OS="darwin" ;;
"Linux") OS="linux" ;;
*)
  echo "Unsupported OS: $OS"
  exit 1
  ;;
esac

case "$CPU" in
"x86_64") CPU="amd64" ;;
"i686" | "i386") CPU="386" ;;
"arm64") CPU="arm64" ;;
*)
  echo "Unsupported CPU: $CPU"
  exit 1
  ;;
esac

# If a version is specified as a command-line argument, use that version.
# Otherwise, fetch the latest release.
if [ -n "$1" ]; then
  VERSION=$1
  URL=$(curl -s https://api.github.com/repos/peter-bread/gamon/releases/tags/v$VERSION |
    grep "browser_download_url.*${OS}_${CPU}*" |
    cut -d : -f 2,3 |
    tr -d \")
else
  URL=$(curl -s https://api.github.com/repos/peter-bread/gamon/releases/latest |
    grep "browser_download_url.*${OS}_${CPU}*" |
    cut -d : -f 2,3 |
    tr -d \")
fi

# Check if URL is found
if [ -z "$URL" ]; then
  echo "No download URL found for OS: ${OS}, CPU: ${CPU}"
  exit 1
fi

# Create the application directories if they don't exist
mkdir -p ~/.gamon

# Download the tarball and extract it to the application directory
if ! wget --show-progress -q $URL -O - | tar xzf - -C ~/.gamon/; then
  echo "Error: Failed to download or extract tarball."
  exit 1
fi

# Ensure /usr/local/bin exists
sudo mkdir -p /usr/local/bin

# Move the binary to /usr/local/bin
sudo mv ~/.gamon/gam /usr/local/bin/gam

echo "Installation completed successfully."
