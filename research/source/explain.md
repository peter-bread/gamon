# Sourcing embedded file

The logic for printing a script from an embedding file is in [source.go](./source.go).

If `gam script` outputs the contents of `script.zsh`, and I want to source this script, I add the following to `~/.zshrc`:

```bash
source <(gam script)
```

Equivalently:

```bash
. <(gam script)
```

Alternatively, use `eval` instead of `source`:

```bash
eval "$(gam script)"
```
