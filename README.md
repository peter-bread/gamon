# gamon <!-- omit from toc -->

GitHub Account Manager

- [Prerequisites](#prerequisites)
- [Dependencies](#dependencies)
- [Installation](#installation)
  - [Homebrew](#homebrew)
  - [Pre-built Binaries](#pre-built-binaries)
  - [Build from Source](#build-from-source)
- [Set up](#set-up)

## Prerequisites

Before you install this software, make sure you have a compatible shell installed. We currently support `bash` and `zsh`, and plan to support `fish` in the future.

|   | MacOS | Linux |
|---|-------|-------|
| Shell | bash / zsh | bash / zsh |
| Future Support | fish | fish |

## Dependencies

This software depends on the following software:

### Software <!-- omit from toc -->

| Application | Link |
|-------------|------|
| gh          | [Link to gh](https://github.com/cli/cli) |
| yq          | [Link to yq](https://github.com/mikefarah/yq) |

`gh` is GitHub’s official command line tool. It brings pull requests, issues, and other GitHub concepts to the terminal next to where you are already working with `git` and your code.

`yq` is a portable command-line YAML processor. It uses `jq`-like syntax but works with yaml files as well as json.

## Installation

### Homebrew

If you're on macOS, you can use [Homebrew](https://brew.sh/):

```shell
brew tap peter-bread/gamon
brew install gamon
```

> Not tested yet, but this might work on Linux if you are using Linuxbrew.

### Pre-built Binaries

We provide pre-built binaries for different operating systems and CPUs. You can download the appropriate binary for your system from the [releases page](https://github.com/peter-bread/gamon/releases).

The filename will be in the format `gamon_x.y.z_<OS>_<CPU>.tar.gz`, where:

- `x.y.z` is the version number.
- `<OS>` should be replaced with your operating system (e.g., `linux`, `darwin` for macOS).
- `<CPU>` should be replaced with your CPU architecture (e.g., `amd64`, `arm64`, `386`).

Once you've downloaded the binary, extract it:

```shell
tar xzf gamon_x.y.z_<OS>_<CPU>.tar.gz
```

To run globally, move the binary to a directory in your PATH:

```shell
mv gamon_x.y.z_<OS>_<CPU> /usr/local/bin
```

If you have any issues with this step, see [Build from Source](#build-from-source) below for guidance.

### Build from Source

You can also build the software from source.

Ensure that you have `go` installed.

```shell
git clone https://github.com/peter-bread/gamon.git
cd gamon
make build
```

This will create a binary in `./bin/gam`.

To make this available globally:

```shell
mv ./bin/gam /usr/local/bin
```

If you encounter permission issues:

```shell
sudo mv ./bin/gam /usr/local/bin
```

If `/usr/local/bin` does not exist, you can create it with:

```shell
sudo mkdir -p /usr/local/bin
```

If for some reason `/bin/local/usr` is not in the `PATH`, then you can add it to your shell configuration file (`~/.bashrc` for bash or `~/.zshrc` for zsh) with:

```shell
echo 'export PATH="/usr/local/bin:$PATH"' >> ~/.bashrc
```

Then source the file for the changes to take effect:

```shell
source ~/.bashrc
```

After installation, you can delete the cloned repository if you wish:

```shell
cd ..
rm -rf gamon
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
