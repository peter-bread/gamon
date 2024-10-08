/*
Copyright © 2024 Peter Sheehan <github.com/peter-bread>
*/
package script

import (
	"embed"
	"fmt"
	"io/fs"
	"os"

	"github.com/peter-bread/gamon/v2/cmd"
	"github.com/spf13/cobra"
)

// embed scripts directory so all can be accessed
//
//go:embed hooks/*
var content embed.FS

// hookCmd represents the hook command
var hookCmd = &cobra.Command{
	Use:   "hook",
	Short: "Generates a script that contains the functions for account switching",
	Long: `Generates a script that contains the functions for account switching.
    
Include the following in your shell configuration file:

    source <(gam hook)`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		var scriptPath string

		// get user's shell
		switch shell := os.Getenv("SHELL"); shell {
		case "/bin/zsh", "/usr/bin/zsh", "/bin/bash", "/usr/bin/bash":
			scriptPath = "hooks/hook.sh"
		case "/bin/fish", "/usr/bin/fish":
			// TODO add case for fish
			fmt.Println("Fish shell not supported yet")
			return
		default:
			fmt.Println("Unsupported shell: ", shell)
			return
		}

		// read the script that corresponds to the shell
		scriptContent, err := fs.ReadFile(content, scriptPath)
		// handle error reading script
		if err != nil {
			fmt.Println("Error: ", err)
			return
			// TODO consider making this command return an error
		}

		// print script
		fmt.Println(string(scriptContent))
	},
}

func init() {
	cmd.RootCmd.AddCommand(hookCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scriptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scriptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
