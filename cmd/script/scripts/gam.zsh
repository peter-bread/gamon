#!/bin/zsh

gh_auth_switch_on_pwd() {

    current_account=$(gh api /user | jq -r .login)

    # TODO read from account names file
    # account_names= ...

    # check directories using the GAM_REPO_ROOT_DIR environment variable
    for account_name in "${account_names[@]}"; do
        if [[ "$PWD" == "$GAM_REPO_ROOT_DIR/$account_name"* && "$current_account" != "$account_name" ]]; then
            gh auth switch --user "$account_name"
        fi
    done

}

# add hook so function is called whenever cd is used
autoload -U add-zsh-hook
add-zsh-hook chpwd gh_auth_switch_on_pwd
