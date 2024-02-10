# gamon

GitHub Account Manager

## Prerequisites

|   | MacOS | Linux |
|---|-------|-------|
| Shell | bash / zsh | bash / zsh |
| Future Support | fish | fish |

## Dependencies

### Applications

| Application | Link |
|-------------|------|
| gh          | [Link to gh](https://github.com/cli/cli) |
| yq          | [Link to jq](https://github.com/mikefarah/yq) |

### Go Packages

| Package | Link |
|---------|------|
| cobra   | [Link to cobra](https://github.com/spf13/cobra) |

## Installation

## Set up

We need to set `GAM_REPO_ROOT_DIR`, the path where all git repos will be stored on the device.

```text
"$GAM_REPO_ROOT_DIR/"
|-- work/
|   |-- project1/
|   |-- project2/
|-- personal/
|   |-- myrepo123/
|
| ...
```

It should be set in `.zshrc` or `.bashrc` using the following:

```shell
export GAM_REPO_ROOT_DIR="/path/to/repo/root/directory"
```
