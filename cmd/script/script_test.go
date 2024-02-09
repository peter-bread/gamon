package script

import (
	"bytes"
	_ "embed"
	"io"
	"os"
	"strings"
	"testing"
)

//go:embed scripts/gam.sh
var script string

func TestScriptCmd_Run(t *testing.T) {
	tests := []struct {
		name     string
		shell    string
		expected string
	}{
		{
			name:     "zsh shell",
			shell:    "/bin/zsh",
			expected: script,
		},
		// {
		// 	name:     "bash shell",
		// 	shell:    "/bin/bash",
		// 	expected: script,
		// },
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
			var buf bytes.Buffer
			io.Copy(&buf, r)
			actualOutput := buf.String()

			actualOutput = strings.TrimSuffix(actualOutput, "\n")

			// Restore the original stdout
			os.Stdout = oldStdout

			if actualOutput != tt.expected {
				// t.Errorf("\nExpected: %s\nActual: %s", tt.expected, actualOutput)
				t.Errorf("\nExpected bytes: %v\nActual bytes: %v", []byte(tt.expected), []byte(actualOutput)) // analyse text exactly
			}
		})
	}
}
