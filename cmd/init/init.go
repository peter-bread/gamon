/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package init

import (
	"fmt"
	"log"
	"os"
	fp "path/filepath"
	"strings"

	"github.com/peter-bread/gamon/v2/cmd"
	"github.com/spf13/cobra"
)

// TODO provide long description

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [filepath]",
	Short: "Creates repository root directory.",
	Long: `Cretes a new directory.

This is used by the other commands to locate the repository root directory.
If no path is provided, the command will ask for a name for the new directory.
If the environment variable $HOME is found in the path, it will be replaced with $HOME.
The absolute path will be printed to the console, and a line to add to the .bashrc or .zshrc file will be printed.
This line sets the environment variable GAM_REPO_ROOT_DIR to the absolute path of the directory.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		// Default path
		filepath := ""

		// If path is provided
		if len(args) > 0 {
			filepath = args[0]
		} else {
			// Ask for path
			fmt.Print("\nEnter a name for the new directory: ")
			_, err := fmt.Scanln(&filepath)
			if err != nil {
				log.Fatalf("Failed to read input: %v", err)
			}
			filepath = "./" + filepath
		}

		// Create directory
		err := os.MkdirAll(filepath, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create directory: %v", err)
		} else {
			fmt.Printf("Directory created: %s\n", filepath)
		}

		// Get the absolute path
		absPath, err := fp.Abs(filepath)
		if err != nil {
			log.Fatalf("Failed to get absolute path: %v", err)
		} else {
			fmt.Printf("Absolute path extracted: %s\n", absPath)

			// Get home directory
			homedir, err := os.UserHomeDir()
			if err != nil {
				log.Fatalf("Failed to get home directory: %v", err)
			} else {
				fmt.Printf("Home directory extracted: %s\n", homedir)
			}

			// Replace home directory with $HOME in the path
			absPath = strings.Replace(absPath, homedir, "$HOME", 1)
		}

		// Print shell command to set environment variable
		fmt.Printf("Add the following line to your .bashrc or .zshrc file:\n\n    export GAM_REPO_ROOT_DIR=\"%s\"\n", absPath)

	},
}

func init() {
	cmd.RootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
