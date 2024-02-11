# Updating

## Homebrew

```shell
brew update
brew upgrade gamon
```

## Pre-built Binaries

```shell
# Remove the old binary
rm /usr/local/bin/gam

# Download the new binary
curl -LO https://github.com/peter-bread/gamon/releases/download/vX.Y.Z/gamon_X.Y.Z_<OS>_<CPU>.tar.gz

# Extract the binary
tar xzf gamon_X.Y.Z_<OS>_<CPU>.tar.gz

# Move the binary to a directory in your PATH and rename to `gam`
mv gamon_X.Y.Z_<OS>_<CPU> /usr/local/bin/gam
```

Replace `X.Y.Z`, `<OS>`, and `<CPU>` with the version number, your operating system, and your CPU architecture, respectively.

## Built from Source

If you built gamon from source, you can pull the latest changes and rebuild the tool:

```shell
# Navigate to the gamon directory
cd path/to/gamon/repository

# Ensure you are on the main branch
git checkout main

# Pull the latest changes
git pull

# Rebuild the tool
make install
```

If this fails due to lack of permissions, try running with sudo:

```shell
sudo make install
```
