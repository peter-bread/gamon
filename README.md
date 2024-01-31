# gamon

GitHub Account Manager

## Prerequisites

### OS

MacOS / Linux

### Shell

zsh / bash (maybe add fish support in future)

## Dependencies

### Applications

- gh
- jq

### Go Packages

- cobra
- promptui ???

## Installation

## Set up

We need to set `GAM_REPO_ROOT_DIR`, the path where all git repos will be stored on the device:

>I think this variable will be set by the `init` command, and if I add a command to move it, then it can be updated there as well.
>
> OR
>
>The `init` command will tell the user what needs to be put in their `.rc` file.

```text
"$GAM_REPO_ROOT_DIR/"
|-- work
|   |-- project1
|   |-- project2
|-- personal
|   |-- myrepo123
|
| ...
```
