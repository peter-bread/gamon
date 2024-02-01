# Commands

## Command Structure

```text
gam <command>
|-- init [<filepath>]
|-- script
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

Without an argument, creates repo structure in the home directory.

With an argument, creates repo structure at that path, e.g. if `filepath="~/repos/"`, then it will create `~/repos/` and any subfolders.

>**Make sure `~` is handled correctly.**

After doing this, it will print out user instructions:

```text
Add the following line to ~/.zshrc:

    export GAM_REPO_ROOT_DIR="$HOME/repos/"
```

`GAM_REPO_ROOT_DIR` will be referenced in the account switching script.

Also copy embedded config file to `$HOME/.config/gamon/config` and create `account_names` in `$HOME/.config/gamon/account_names`.

### `gam script`

Prints script necessary for account switching.

Checks user's shell and prints the corresponding script.

In shell `.rc` file, include:

```bash
source <(gam script)
```

Or:

```bash
eval "$(gam script)"
```

<!-- Should add one of the following to shell configuration file:

```bash
# ~/.zshrc
eval "$(gam script zsh)"

# ~/.bashrc
eval "$(gam script bash)"
``` -->

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
