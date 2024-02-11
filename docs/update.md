# Updating

## Homebrew

If you installed `gamon` using Homebrew, you can update it by first updating Homebrew itself and then upgrading the `gamon` package:

```shell
brew update
brew upgrade gamon
```

## Pre-built Binaries

To update `gamon` when you've installed it using a pre-built binary, you'll need to manually replace the old binary with the new one. Here's how you can do it:

```shell
# Remove the old binary
# You may need sudo for this
rm /usr/local/bin/gam

# Download the new binary
curl -LO https://github.com/peter-bread/gamon/releases/download/vX.Y.Z/gamon_X.Y.Z_<OS>_<CPU>.tar.gz

# Extract the binary
tar xzf gamon_X.Y.Z_<OS>_<CPU>.tar.gz

# Move the binary to a directory in your PATH and rename to `gam`
# You may need sudo for this
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
