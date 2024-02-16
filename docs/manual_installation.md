# Manual Installation

## Pre-built Binaries

We provide pre-built binaries for different operating systems and CPUs. You can download the appropriate binary for your system from the [releases page](https://github.com/peter-bread/gamon/releases).

### Download the `.tar.gz`

<!-- omit from toc -->
#### Latest Release

To install the latest release, pick one of the following.

Latest Apple Silicon:

```shell
curl -s https://api.github.com/repos/peter-bread/gamon/releases/latest \
| grep "browser_download_url.*darwin_arm64*" \
| cut -d : -f 2,3 \
| tr -d \" \
| wget --show-progress -qi -
```

Latest Linux x86_64:

```shell
curl -s https://api.github.com/repos/peter-bread/gamon/releases/latest \
| grep "browser_download_url.*linux_amd64*" \
| cut -d : -f 2,3 \
| tr -d \" \
| wget --show-progress -qi -
```

Other:

```shell
curl -s https://api.github.com/repos/peter-bread/gamon/releases/latest \
| grep "browser_download_url.*{OS}_{CPU}*" \
| cut -d : -f 2,3 \
| tr -d \" \
| wget --show-progress -qi -
```

<!-- NOTE: if I release deb or RPM packages in the future, I will need to update the grep command to reflect that -->
<!-- `grep "browser_download_url.*linux_amd64.tar.gz" -->
<!-- `grep "browser_download_url.*linux_amd64.deb" -->
<!-- `grep "browser_download_url.*linux_amd64.rpm" -->

<!-- omit from toc -->
#### Specific Version

Alternatievly, see below to download a specific version from a terminal.

The filename will be in the format `gamon_X.Y.Z_{OS}_{CPU}.tar.gz`, where:

- `X.Y.Z` is the version number.
- `{OS}` should be replaced with your operating system (e.g., `linux`, `darwin` for macOS).
- `{CPU}` should be replaced with your CPU architecture (e.g., `amd64`, `arm64`, `386`).

Download the correct binary:

```shell
curl -LO https://github.com/peter-bread/gamon/releases/download/vX.Y.Z/gamon_X.Y.Z_{OS}_{CPU}.tar.gz
```

Alteratively, download from the [releases page](https://github.com/peter-bread/gamon/releases).

### Extract and Install

Once you've downloaded the binary, extract it into the `~/.gamon` directory:

```shell
mkdir -p ~/.gamon && tar xzf gamon_X.Y.Z_{OS}_{CPU}.tar.gz -C ~/.gamon
```

To run globally, move the binary to a directory in your PATH:

```shell
mv ~/.gamon/gam ~/.gamon/bin
```

Finally, you can delete the `.tar.gz` file:

```shell
rm gamon_X.Y.Z_{OS}_{CPU}.tar.gz
```

Remember to replace `X.Y.Z`, `{OS}` and `{CPU}` with the version number, your operating system, and your CPU architecture, respectively.

## Build from Source

You can also build the software from source.

Ensure that you have both the [build and runtime dependencies](../README.md#dependencies) installed.

### Installation Script

```shell
curl -fsSL https://raw.githubusercontent.com/peter-bread/gamon/main/scripts/install_from_source.sh | bash
```

> Before using this script, please ensure you review and understand the operations it performs. This is to ensure that it aligns with your intended use and doesn't cause unintended effects.

### Manual Terminal Commands

```shell
# Clone gamon repository
git clone https://github.com/peter-bread/gamon.git

# Naviage to the gamon directory
cd gamon

# Ensure you are on the main branch
git checkout main

# Build and install the tool
make install
```

You will be prompted to add `~/.gamon/bin` to PATH in `.bashrc` or `.zshrc`.
