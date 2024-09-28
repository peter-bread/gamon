# This script is intended to be sourced from a .bashrc or .zshrc file.
# It uses features that are specific to bash and zsh, and may not work correctly in other shells.

# get current shell
current_shell=$(ps -p $$ -ocomm=)

cmd="$(gam run)"

# calls function to check and/or switch github account on every cd
if [[ $current_shell == *"zsh"* ]]; then
  autoload -U add-zsh-hook
  add-zsh-hook chpwd "$cmd"
elif [[ $current_shell == *"bash"* ]]; then
  cd() {
    builtin cd "$@" || exit
    "$cmd"
  }
else
  echo "Error: Unsupported shell. Only zsh and bash are supported." >&2
  return 1
fi
