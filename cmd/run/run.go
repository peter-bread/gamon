/*
Copyright © 2024 Peter Sheehan <github.com/peter-bread>
*/

package run

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/peter-bread/gamon/v2/cmd"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// Struct to represent the YAML structure
type HostConfig struct {
	Users map[string]struct{} `yaml:"users"`
	User  string              `yaml:"user"`
}

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Change to correct GH CLI account based on directory",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Check and validate environment variables
		gamRepoRootDir, err := getEnv("GAM_REPO_ROOT_DIR")
		if err != nil {
			log.Fatal(err)
		}

		// Ensure that GAM_REPO_ROOT_DIR is valid
		if err := ensureDirectory(gamRepoRootDir); err != nil {
			log.Fatal(err)
		}

		// Check if 'gh' is installed
		if err := checkCommandExists("gh"); err != nil {
			log.Fatal(err)
		}

		// Resolve GitHub config directory
		configDir := getGitHubConfigDir()
		if err := ensureDirectory(configDir); err != nil {
			log.Fatal(err)
		}

		// Load accounts from the config file
		accountsInfo, err := loadAccountsInfo(filepath.Join(configDir, "hosts.yml"))
		if err != nil {
			log.Fatal(err)
		}

		// Get the current directory
		currentDir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		// Switch accounts if needed
		if err := switchAccountIfNeeded(gamRepoRootDir, currentDir, accountsInfo); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scriptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scriptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// getEnv checks and returns the value of the specified environment variable.
func getEnv(varName string) (string, error) {
	value := os.Getenv(varName)
	if value == "" {
		return "", fmt.Errorf("%s is not set. Please set it and try again.", varName)
	}
	return strings.TrimSuffix(value, "/"), nil
}

// ensureDirectory checks if a directory exists.
func ensureDirectory(dir string) error {
	if stat, err := os.Stat(dir); err != nil || !stat.IsDir() {
		return fmt.Errorf("error: %s could not be found or is not a directory", dir)
	}
	return nil
}

// checkCommandExists ensures a command exists on the system.
func checkCommandExists(cmd string) error {
	if _, err := exec.LookPath(cmd); err != nil {
		return fmt.Errorf("error: %s is not installed", cmd)
	}
	return nil
}

// getGitHubConfigDir returns the GitHub config directory.
func getGitHubConfigDir() string {
	if ghConfigDir := os.Getenv("GH_CONFIG_DIR"); ghConfigDir != "" {
		return strings.TrimSuffix(ghConfigDir, "/")
	}
	return filepath.Join(os.Getenv("HOME"), ".config", "gh")
}

// loadAccountsInfo reads and parses the hosts.yml file.
func loadAccountsInfo(hostsFilePath string) (*HostConfig, error) {
	file, err := os.ReadFile(hostsFilePath)
	if err != nil {
		return nil, fmt.Errorf("error: could not find accounts in hosts.yml")
	}

	var config map[string]HostConfig
	if err := yaml.Unmarshal(file, &config); err != nil {
		return nil, fmt.Errorf("error: failed to parse hosts.yml")
	}

	githubConfig, exists := config["github.com"]
	if !exists || githubConfig.User == "" {
		return nil, fmt.Errorf("error: could not find current account in hosts.yml")
	}

	return &githubConfig, nil
}

// switchAccountIfNeeded checks the current directory and switches the GitHub account if needed.
func switchAccountIfNeeded(gamRepoRootDir, currentDir string, config *HostConfig) error {
	for accountName := range config.Users {
		accountDir := filepath.Join(gamRepoRootDir, accountName)

		// Check if the current directory is within the user's repo
		if strings.HasPrefix(currentDir, accountDir) && config.User != accountName {

			// Capture the output from the 'gh' command
			cmd := exec.Command("gh", "auth", "switch", "--user", accountName)

			// Force gh to output colored text even when not running in an interactive terminal
			cmd.Env = append(os.Environ(), "GH_FORCE_TTY=1")

			output, err := cmd.CombinedOutput()
			if err != nil {
				return fmt.Errorf("error: could not switch to account %s: %v\nCommand output: %s", accountName, err, output)
			}

			// Print the command's output
			fmt.Printf("%s", output)
			break
		}

		// Create the directory if it doesn't exist
		// if err := os.MkdirAll(accountDir, 0755); err != nil {
		// 	return fmt.Errorf("error: could not create directory %s", accountDir)
		// }
	}
	return nil
}
