<!-- omit from toc -->
# gamon

![GitHub Release](https://img.shields.io/github/v/release/peter-bread/gamon?style=for-the-badge&color=ff00a1) ![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/peter-bread/gamon/go.yml?branch=dev&style=for-the-badge&label=build%20and%20test%3A%20dev) ![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/peter-bread/gamon/go.yml?branch=main&style=for-the-badge&label=build%20and%20test%3A%20main) ![GitHub commits since latest release](https://img.shields.io/github/commits-since/peter-bread/gamon/latest?style=for-the-badge) ![GitHub License](https://img.shields.io/github/license/peter-bread/gamon?style=for-the-badge&color=009200&link=github.com%2Fpeter-bread%2Fgamon%2Fblob%2Fmain%2FLICENSE)

---

Gamon is a command line tool that helps with managing multiple GitHub accounts on one machine.

Right now, it's main functionality is switching the active `gh` account based on your current working directory.

<!-- omit from toc -->
## Table of Contents

- [Prerequisites](#prerequisites)
- [Dependencies](#dependencies)
- [Documentation](#documentation)
- [Installation](#installation)
- [Set up](#set-up)
- [Usage](#usage)
- [License](#license)

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

This script downloads the pre-built binary, moves it to `~/.gamon/bin`, and performs other necessary setup tasks.

> Before using this script, please ensure you review and understand the operations it performs. This is to ensure that it aligns with your intended use and doesn't cause unintended effects.

You will be prompted to add `~/.gamon/bin` to PATH in your `.bashrc` or `zshrc`.

For the latest release:

```shell
curl -fsSL https://raw.githubusercontent.com/peter-bread/gamon/main/scripts/install_binary.sh | bash
```

For a specific version:

```shell
curl -fsSL https://raw.githubusercontent.com/peter-bread/gamon/main/scripts/install_binary.sh | bash -s -- X
```

where `X` is the major version you wish to install.

> **Note:** The versioning scheme used is `X.Y.Z`. When you specify `X` in the command above, the script will install the latest release within that major version.

See [manual installation](./docs/manual_installation.md) for more installation options.

## Set up

See [this guide](https://github.com/peter-bread/git-ssh-management) for a setting up the required file structure.

> In future releases, I hope to automate much of the set up process.
>
> *Or change it entirely...*

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

## Usage

After settimg up `gam` as described in the [set up](#set-up) section, you can start using it to manage your GitHub accounts.

`gam` uses the `GAM_REPO_ROOT_DIR` environment variable to determine which GitHub account to use based on the current repository. When you navigate to a directory within `GAM_REPO_ROOT_DIR` and run a GitHub CLI command, `gam` will automatically switch to the account associated with that directory.

Remember to source the `gam script` in each new shell session to enable automatic account switching:

```shell
source <(gam script)
```

## License

This project is licensed under the terms of the MIT license. See the [LICENSE](./LICENSE) file for the full license text.
