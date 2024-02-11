# Updating

## Homebrew

If you installed `gamon` using Homebrew, you can update it by first updating Homebrew itself and then upgrading the `gamon` package:

```shell
brew update
brew upgrade gamon
```

> Ensure that the Homebrew tap `peter-bread/gamon` is still active.

## Pre-built Binaries

To update `gamon` when you've installed it using a pre-built binary, you'll need to manually replace the old binary with the new one. To do this you simply need to reinstall the tool, either with the [installation script](../README.md#pre-built-binaries) or [manually](./manual_installation.md#pre-built-binaries).

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
