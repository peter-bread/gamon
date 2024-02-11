# gamon <!-- omit from toc -->

GitHub Account Manager

- [Prerequisites](#prerequisites)
- [Dependencies](#dependencies)
- [Documentation](#documentation)
- [Installation](#installation)
- [Set up](#set-up)

## Prerequisites

Before you install this software, make sure you have a compatible shell installed. We currently support `bash` and `zsh`, and plan to support `fish` in the future.

|                | MacOS      | Linux      |
| -------------- | ---------- | ---------- |
| Shell          | bash / zsh | bash / zsh |
| Future Support | fish       | fish       |

## Dependencies

This software depends on the following software:

<!-- omit from toc -->
### Build Dependencies

| Application | Link                              |
| ----------- | --------------------------------- |
| Go          | [Link to Go](https://golang.org/) |

<!-- omit from toc -->
### Runtime Dependencies

| Application | Link                                          |
| ----------- | --------------------------------------------- |
| gh          | [Link to gh](https://github.com/cli/cli)      |
| yq          | [Link to yq](https://github.com/mikefarah/yq) |

`gh` is GitHub’s official command line tool. It brings pull requests, issues, and other GitHub concepts to the terminal next to where you are already working with `git` and your code.

`yq` is a portable command-line YAML processor. It uses `jq`-like syntax but works with yaml files as well as json.

## Documentation

- [Manual Installation](./docs/manual_installation.md)
- [Updating](./docs/update.md)
- [Uninstalling](./docs/uninstall.md)

## Installation

Below are the two simplest ways to install the gamon binary, `gam`.

Check out [manual installation](./docs/manual_installation.md) for more detailed installation instructions.

<!-- omit from toc -->
### Homebrew

If you're on macOS, you can use [Homebrew](https://brew.sh/):

```shell
brew tap peter-bread/gamon
brew install gamon
```

Homebrew will manage dependencies automatically.

> Not tested yet, but this might work on Linux if you are using Linuxbrew.

<!-- omit from toc -->
### Pre-built Binaries

Ensure you have installed the [runtime dependencies](#runtime-dependencies).

You can use the [installation script](./scripts/install_binary.sh) to download, extract and install the binary.

This script downloads the pre-built binary, moves it to `/usr/local/bin`, and performs other necessary setup tasks.

> Before using this script, please ensure you review and understand the operations it performs. This is to ensure that it aligns with your intended use and doesn't cause unintended effects.

You will be prompted for your password as the script does require sudo to:

1. Create `/usr/local/bin` if it doesn't already exist.
2. Copy the binary to `/usr/local/bin`.

For the latest release:

```shell
curl -fsSL https://raw.githubusercontent.com/peter-bread/gamon/main/scripts/install.sh | bash
```

For a specific version:

```shell
curl -fsSL https://raw.githubusercontent.com/peter-bread/gamon/main/scripts/install.sh | bash -s -- X.Y.Z
```

where `X.Y.Z` is the version you wish to install.

See [manual installation](./docs/manual_installation.md) for more installation options.

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
