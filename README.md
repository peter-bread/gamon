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

## Command Structure

```text
gam <command>
|-- init [<filepath>]
|-- account <command>
|   |-- add <account name>
|   |-- remove <account name>
|   |-- edit
|   |-- list
|   |-- view <account name>
|-- config
```

### `gam <command>`

Root command. Must be followed by a subcommand.

### `gam init [<filepath>]`

Initialises file structure. Optional argument filepath.

Without an argument, creates repo structure in (home or current ??? directory).

With an argument, creates repo structure at that path (is the path where `repos/` is created or is it where `work/` and `personal/` are created ??? leaning towars `repos/` I think ???).

### `gam account <command>`

Manages account names. Must be followed by a sucommand.

### `gam account add <account name>`

Takes in a single argument: account name.

Adds account name to account names file (if it is not already there).

### `gam account remove <account name>`

Takes in a single argument: account name.

Removes account name from account names file (if it is already there).

### `gam account edit`

Opens account names file using `$EDITOR` if set, else `nano`, so user can manually add/remove accounts.

### `gam account list`

Print all currently added account names to the terminal.

### `gam account view <account name>`

Takes in a single argument: account name.

Lists all the repos in that account's directory.

Could print any other settings associated with that account.

### `gam config`

Opens config file using `$EDITOR` if set, else `nano`.

Allows user to adjust other settings (yet to be defined).
