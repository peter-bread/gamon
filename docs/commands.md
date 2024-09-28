# Commands

## Command Structure

```text
gam <command>
|-- hook
|-- run
|-- completion <shell> //builtin
```

### `gam <command>`

Root command. Must be followed by a subcommand.

### `gam hook`

Hooks into shell so `gam run` is executed on every `cd`.

Checks user's shell and prints the corresponding script.

In shell `.rc` file, include:

```bash
source <(gam hook)
```

Or:

```bash
eval "$(gam hook)"
```

### gam run

This is the binary that checks your directory and switches to the correct account.

This is invoked automatically by the script printed in `gam hook`.

### `gam completion <shell>`

_THIS IS A BUILTIN COMMAND._

Generates completion script for specified shell.

To load in current shell session:

```bash
source <(gam completion <shell>)
```

This could be added to `~/.zshrc` if not a homebrew package.

**_IF_** it is a homebrew package, then completions should have been set up automatically.

If this is not the case, you can run the following command once:

```bash
gam completion zsh > $(brew --prefix)/share/zsh/site-functions/_gam
```

## Future Commands

> [!WARNING]
> These are not implemented yet.
> I don't know if they ever will be.

### `gam init [<filepath>]`

Initialises file structure. Optional argument filepath.

The default directory name is `repos/`.

With no arguments: creates `repos/` in current directory and prints command
to set `$GAM_REPO_ROOT_DIR`.

With `.`: assumes current directory is `repos/` and prints command to set `$GAM_REPO_ROOT_DIR`.

With `<filepath>`: creates the specified directory. E.g:

```bash
# Creates my-repos/ if it doesn't already exist
gam init ~/Documents/my-repos
```

In this case, you can name the `repos/` directory whatever you'd like, in the
above example it is called `my-repos/`.

With an argument, creates repo structure at that path, e.g. if `filepath="~/repos/"`,
then it will create `~/repos/` and prints command to set `$GAM_REPO_ROOT_DIR`.

After doing this, it will print out user instructions:

```text
Add the following line to ~/.zshrc:

    export GAM_REPO_ROOT_DIR="$HOME/repos/"
```

The following are equivalent:

```bash
gam init
gam init repos
gam init ./repos
```

`GAM_REPO_ROOT_DIR` will be referenced in the account switching script.

Also copy embedded config file to `$HOME/.config/gamon/config` and create
`account_names` in `$HOME/.config/gamon/account_names`.

### `gam create-acc-dirs`

Generates directories for all accounts.

This command reads the account names file and creates a directory in `$GAM_REPO_ROOT_DIR`
for each account that doesn't already have one.

### `gam account <command>`

Manages account names. Must be followed by a subcommand.

### `gam account add <account name>`

Takes in a single argument: account name.

Adds account name to account names file (if it is not already there).

**MAYBE**: there will be a `--add-dir` flag to automatically create new directory
in `$GAM_REPO_ROOT_DIR`. (I might make this default behaviour and add a `--no-dir`
as an option and let the user adjust the default in config file). This will only
work if `gam init` has already been used. If `$GAM_REPO_ROOT_DIR` doesn't exist,
throw error.

**MAYBE**: there will be a `--ssh` flag, which will generate ssh keys that correspond
to the account name.

**MAYBE**: there will be a `--login` flag, which will call `gh auth login`
(this will only work if `--ssh` has been used as well).

### `gam account remove <account name>`

Takes in a single argument: account name.

Removes account name from account names file (if it is already there).

**MAYBE**: add flag `--remove-dir` which will also delete the directory and all
repos in it.

### `gam account edit`

Opens account names file using `$EDITOR` if set, else `nano`, so user can
manually add/remove accounts.

### `gam account list`

Print all currently added account names to the terminal.

### `gam account view <account name>`

Takes in a single argument: account name.

Lists all the repos in that account's directory.

Could print any other settings associated with that account.

### `gam config`

Opens config file using `$EDITOR` if set, else `nano`.

Allows user to adjust other settings (yet to be defined).
