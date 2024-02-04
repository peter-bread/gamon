package script

import (
	_ "embed"
	"os"
	"strings"
	"testing"
)

//go:embed scripts/gam.zsh
var zshScript string

//go:embed scripts/gam.bash
var bashScript string

func TestScriptCmd_Run(t *testing.T) {
	tests := []struct {
		name     string
		shell    string
		expected string
	}{
		{
			name:     "zsh shell",
			shell:    "/bin/zsh",
			expected: zshScript,
		},
		{
			name:     "bash shell",
			shell:    "/bin/bash",
			expected: bashScript,
		},
		// TODO: Add test case for fish shell when implemented
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture the current stdout
			oldStdout := os.Stdout

			// Create a pipe to capture the output
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Set the SHELL environment variable
			os.Setenv("SHELL", tt.shell)

			// Run the scriptCmd
			scriptCmd.Run(scriptCmd, []string{})

			// Close the write end of the pipe
			w.Close()

			// Capture the output
			output := make([]byte, 1024)
			n, _ := r.Read(output)
			actualOutput := string(output[:n])

			actualOutput = strings.TrimSuffix(actualOutput, "\n")

			// Restore the original stdout
			os.Stdout = oldStdout

			if actualOutput != tt.expected {
				t.Errorf("\nExpected: %s\nActual: %s", tt.expected, actualOutput)
				t.Errorf("\nExpected bytes: %v\nActual bytes: %v", []byte(tt.expected), []byte(actualOutput)) // analyse text exactly
			}
		})
	}
}
