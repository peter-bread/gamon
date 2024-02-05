/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package init

import (
	"bufio"
	"fmt"
	"log"
	"os"
	fp "path/filepath"
	"strings"

	"github.com/peter-bread/gamon/v2/cmd"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [filepath]",
	Short: "Creates repository root directory.",
	Long: `Creates repository root directory.

This is used by the other commands to locate the repository root directory.
If no path is provided, the command will ask for a name for the new directory.
If the environment variable $HOME is found in the path, it will be replaced with $HOME.
The absolute path will be printed to the console, and a line to add to the .bashrc or .zshrc file will be printed.
This line sets the environment variable GAM_REPO_ROOT_DIR to the absolute path of the directory.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		// Default path
		filepath := ""

		// If path is not provided
		if len(args) == 0 {
			// Ask for path
			fmt.Print("\nEnter a name for the new directory: ")
			reader := bufio.NewReader(os.Stdin)

			// Read input
			var err error
			filepath, err = reader.ReadString('\n')

			// Check for errors
			if err != nil {
				log.Fatalf("Failed to read input: %v", err)
			}
		} else {
			filepath = args[0]
		}

		filepath = processFilepath(filepath)

		// Create directory
		err := os.MkdirAll(filepath, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create directory: %v", err)
		} else {
			fmt.Printf("\nDirectory created: %s\n", filepath)
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

		// 		// Default path
		// 		filepath := ""

		// 		// If path is provided
		// 		if len(args) > 0 {

		// 			// Trim whitespace
		// 			filepath = strings.TrimSpace(args[0])

		// 			// Replace multiple spaces with single space
		// 			filepath = strings.Join(strings.Fields(filepath), " ")

		// 			// Replace backslashes with forward slashes
		// 			filepath = strings.ReplaceAll(filepath, " / ", "/")

		// 			// Check for empty string
		// 			if filepath == "" {
		// 				log.Fatalf("Directory name cannot be empty")
		// 			}

		// 			// Check for special characters
		// 			if strings.ContainsAny(filepath, `\|:;*?"<>`) {
		// 				log.Fatalf("Directory name cannot contain special characters")
		// 			}

		// 			// Check for trailing dot (except for ".")
		// 			if strings.HasSuffix(filepath, ".") && filepath != "." {
		// 				log.Fatalf("Directory name cannot end with a dot")
		// 			}

		// 		} else {
		// 			// Ask for path
		// 			fmt.Print("\nEnter a name for the new directory: ")
		// 			reader := bufio.NewReader(os.Stdin)

		// 			// Read input
		// 			var err error
		// 			filepath, err = reader.ReadString('\n')

		// 			// Check for errors
		// 			if err != nil {
		// 				log.Fatalf("Failed to read input: %v", err)
		// 			}

		// 			// Trim whitespace
		// 			filepath = strings.TrimSpace(filepath)

		// 			// Replace multiple spaces with single space
		// 			filepath = strings.Join(strings.Fields(filepath), " ")

		// 			// Replace backslashes with forward slashes
		// 			filepath = strings.ReplaceAll(filepath, " / ", "/")

		// 			// Check for empty string
		// 			if filepath == "" {
		// 				log.Fatalf("Directory name cannot be empty")
		// 			}

		// 			// Check for special characters
		// 			if strings.ContainsAny(filepath, `\|:;*?"<>`) {
		// 				log.Fatalf("Directory name cannot contain special characters")
		// 			}

		// 			// Check for trailing dot (except for ".")
		// 			if strings.HasSuffix(filepath, ".") && filepath != "." {
		// 				log.Fatalf("Directory name cannot end with a dot")
		// 			}

		// 		}

		// 		// Prepend "./" to the path if it doesn't already start with it
		// 		if !strings.HasPrefix(filepath, "./") {
		// 			filepath = "./" + filepath
		// 		}

		// 		// Remove trailing "/" from the path if it exists
		// 		filepath = strings.TrimSuffix(filepath, "/")

		// 		// Create directory
		// 		err := os.MkdirAll(filepath, os.ModePerm)
		// 		if err != nil {
		// 			log.Fatalf("Failed to create directory: %v", err)
		// 		} else {
		// 			fmt.Printf("\nDirectory created: %s\n", filepath)
		// 		}

		// 		// Get the absolute path
		// 		absPath, err := fp.Abs(filepath)
		// 		if err != nil {
		// 			log.Fatalf("Failed to get absolute path: %v", err)
		// 		} else {
		// 			fmt.Printf("Absolute path extracted: %s\n", absPath)

		// 			// Get home directory
		// 			homedir, err := os.UserHomeDir()
		// 			if err != nil {
		// 				log.Fatalf("Failed to get home directory: %v", err)
		// 			} else {
		// 				fmt.Printf("Home directory extracted: %s\n", homedir)
		// 			}

		// 			// Replace home directory with $HOME in the path
		// 			absPath = strings.Replace(absPath, homedir, "$HOME", 1)
		// 		}

		// 		// Print shell command to set environment variable
		// 		fmt.Printf("Add the following line to your .bashrc or .zshrc file:\n\n    export GAM_REPO_ROOT_DIR=\"%s\"\n", absPath)
	},
}

func processFilepath(filepath string) string {
	// Trim whitespace
	filepath = strings.TrimSpace(filepath)

	// Replace multiple spaces with single space
	filepath = strings.Join(strings.Fields(filepath), " ")

	// Remove spaces around slashes
	filepath = strings.ReplaceAll(filepath, " / ", "/")

	// TODO replace consecutive slashes with single slash

	// Check for empty string
	if filepath == "" {
		log.Fatalf("Directory name cannot be empty")
	}

	// Check for special characters
	if strings.ContainsAny(filepath, `\|:;*?"<>`) {
		log.Fatalf("Directory name cannot contain special characters")
	}

	// Check for trailing dot (except for ".")
	if strings.HasSuffix(filepath, ".") && filepath != "." {
		log.Fatalf("Directory name cannot end with a dot")
	}

	// Prepend "./" to the path if it doesn't already start with it
	if !strings.HasPrefix(filepath, "./") {
		filepath = "./" + filepath
	}

	// TODO remove repeated ./ from the path

	// Remove trailing "/" from the path if it exists
	filepath = strings.TrimSuffix(filepath, "/")

	return filepath

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
