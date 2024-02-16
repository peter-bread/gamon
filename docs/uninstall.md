# Uninstall

## Homebrew

If you installed `gamon` using Homebrew, you can uninstall it using the `brew uninstall` command:

```shell
brew uninstall gamon
```

Optionally, you can remove the tap:

```shell
brew untap peter-bread/gamon
```

## Pre-built Binaries

If you installed `gamon` using a pre-built binary, you can uninstall it by simply removing the binary:

```shell
rm -r ~/.gamon/bin
```

You can also delete the whole directory that it was originally extracted to, `~/.gamon`:

```shell
rm -r ~/.gamon
```

## Built from Source

If you built `gamon` from source, you can uninstall it by removing the binary:

```shell
# Navigate to the gamon directory
cd path/to/gamon/repository

# Run the uninstall command
make uninstall
```

You can also remove the source code if you no longer need it:

```shell
cd ..
rm -rf path/to/gamon/repository
```

Replace `path/to/gam/repository` with the actual path to the `gamon` repository on your system.
