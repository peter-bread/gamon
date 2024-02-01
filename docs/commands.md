# Commands

## Command Structure

```text
gam <command>
|-- init [<filepath>]
|-- script
|-- account <command>
|   |-- add <account name> [flags]
|   |-- remove <account name>
|   |-- edit
|   |-- list
|   |-- view <account name>
|-- config
|-- completion <shell> //builtin
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

### `gam account <command>`

Manages account names. Must be followed by a sucommand.

### `gam account add <account name>`

Takes in a single argument: account name.

Adds account name to account names file (if it is not already there).

**MAYBE**: there will be a `--add-dir` flag to automatically create new directory in `$GAM_REPO_ROOT_DIR`. (I might make this default behaviour and add a `--no-dir` as an option and let the user adjust the default in config file)

**MAYBE**: there will be a `--ssh` flag, which will generate ssh keys that correspond to the account name.

### `gam account remove <account name>`

Takes in a single argument: account name.

Removes account name from account names file (if it is already there).

**MAYBE**: add flag `--remove-dir` which will also delete the directory and all repos in it.

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

### `gam completion <shell>`

*THIS IS A BUILTIN COMMAND.*

Generates completion script for specified shell.

To load in current shell session:

```bash
source <(gam completion <shell>)
```

This could be added to `~/.zshrc` if not a homebrew package.

***IF*** this is a homebrew package, run the following command once:

```bash
gam completion zsh > $(brew --prefix)/share/zsh/site-functions/_gam
```
