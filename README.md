# gamon <!-- omit from toc -->

GitHub Account Manager

- [Prerequisites](#prerequisites)
- [Dependencies](#dependencies)
- [Installation](#installation)
    - [Latest Release](#latest-release)
      - [Latest Apple Silicon](#latest-apple-silicon)
      - [Latest Linux x86\_64](#latest-linux-x86_64)
      - [Other](#other)
    - [Specific Version](#specific-version)
- [Set up](#set-up)
- [Documentation](#documentation)

## Prerequisites

Before you install this software, make sure you have a compatible shell installed. We currently support `bash` and `zsh`, and plan to support `fish` in the future.

|   | MacOS | Linux |
|---|-------|-------|
| Shell | bash / zsh | bash / zsh |
| Future Support | fish | fish |

## Dependencies

This software depends on the following software:

<!-- omit from toc -->
### Build Dependencies

| Application | Link |
|-------------|------|
| Go          | [Link to Go](https://golang.org/) |

<!-- omit from toc -->
### Runtime Dependencies

| Application | Link |
|-------------|------|
| gh          | [Link to gh](https://github.com/cli/cli) |
| yq          | [Link to yq](https://github.com/mikefarah/yq) |

`gh` is GitHub’s official command line tool. It brings pull requests, issues, and other GitHub concepts to the terminal next to where you are already working with `git` and your code.

`yq` is a portable command-line YAML processor. It uses `jq`-like syntax but works with yaml files as well as json.

## Installation

<!-- omit from toc -->
### Homebrew

Homebrew will manage dependencies automatically.

If you're on macOS, you can use [Homebrew](https://brew.sh/):

```shell
brew tap peter-bread/gamon
brew install gamon
```

> Not tested yet, but this might work on Linux if you are using Linuxbrew.

<!-- omit from toc -->
### Pre-built Binaries

Ensure you have installed the [runtime dependencies](#runtime-dependencies).

We provide pre-built binaries for different operating systems and CPUs. You can download the appropriate binary for your system from the [releases page](https://github.com/peter-bread/gamon/releases).

You can download from the terminal:

#### Latest Release

To install the latest release, pick one of the following:

##### Latest Apple Silicon

```shell
curl -s https://api.github.com/repos/peter-bread/gamon/releases/latest \
| grep "browser_download_url.*darwin_arm64*" \
| cut -d : -f 2,3 \
| tr -d \" \
| wget --show-progress -qi -
```

##### Latest Linux x86_64

```shell
curl -s https://api.github.com/repos/peter-bread/gamon/releases/latest \
| grep "browser_download_url.*linux_amd64*" \
| cut -d : -f 2,3 \
| tr -d \" \
| wget --show-progress -qi -
```

##### Other

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

Once you've downloaded the binary, extract it:

```shell
tar xzf gamon_X.Y.Z_<OS>_<CPU>.tar.gz
```

To run globally, move the binary to a directory in your PATH and rename it `gam`:

```shell
mv gamon_X.Y.Z_<OS>_<CPU> /usr/local/bin/gam
```

If you have any issues with this step, see [Build from Source](#build-from-source) below for guidance.

<!-- omit from toc -->
### Build from Source

You can also build the software from source.

Ensure that you have both the [build and runtime dependencies](#dependencies) installed.

```shell
# Clone gamon repository
git clone https://github.com/peter-bread/gamon.git

# Naviage to the gamon directory
cd gamon

# Ensure you are on the main branch
git checkout main

# Build the tool
make install
```

This will create a binary in `./bin/gam` and attempt to copy it to `/usr/local/bin`.

If this fails due to lack of permissions, try running with sudo:

```shell
sudo make install
```

If for some reason `/bin/local/usr` is not in the `PATH`, then you can add it to your shell configuration file (`~/.bashrc` for bash or `~/.zshrc` for zsh) with:

```shell
echo 'export PATH="/usr/local/bin:$PATH"' >> ~/.bashrc
```

Then source the file for the changes to take effect:

```shell
source ~/.bashrc
```

<!-- ### Installation Script -->

## Set up

See [this guide](https://github.com/peter-bread/git-ssh-management) for a detailed set up guide.

> In future releases, I hope to automate much of the set up process.

You need to set `GAM_REPO_ROOT_DIR`, the path where all of **your** git repositories will be stored.

This is illustrated below:

```text
"$GAM_REPO_ROOT_DIR/"
|-- personal/
|   |-- myrepo123/
|-- work/
|   |-- project1/
|   |-- project2/
|
| ...
```

To set `GAM_REPO_ROOT_DIR`, add the following line to your `.zshrc` or `.bashrc` file:

```shell
export GAM_REPO_ROOT_DIR="/path/to/repo/root/directory"
```

Replace `"/path/to/repo/root/directory"` with the actual path to your repository root directory.

Then, to enable automatic account switching based on the repository you're working with, add the following line to your .zshrc or .bashrc file:

```shell
source <(gam script)
```

This will run the gam script command every time you start a new shell session, setting up the necessary environment for automatic account switching.

## Documentation

- [Updating](./docs/update.md)
- [Uninstalling](./docs/uninstall.md)
