/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package init

import (
	"fmt"

	"github.com/peter-bread/gamon/v2/cmd"
	"github.com/spf13/cobra"
)

// TODO provide long description

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [<filepath>]",
	Short: "Builds repository file structure in specified path.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
		// TODO define command

		if len(args) == 0 {
			// if no argument is passed, generate in home directory
			return
		} else {
			// if filepath is provided, generate in the given path
			return
		}

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
