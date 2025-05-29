package main

import (
	"os"
	"strings"
	"testing"
)

func TestFormatPathReplacesHomeDirectory(t *testing.T) {
	// Save original HOME to restore
	originalHome := os.Getenv("HOME")
	defer os.Setenv("HOME", originalHome)

	tests := []struct {
		name     string
		home     string
		path     string
		expected string
	}{
		{
			"replaces home with tilde",
			"/Users/test",
			"/Users/test/project",
			"~/project",
		},
		{
			"handles exact home path",
			"/Users/test",
			"/Users/test",
			"~",
		},
		{
			"ignores non-home paths",
			"/Users/test",
			"/tmp/project",
			"/tmp/project",
		},
		{
			"handles empty home",
			"",
			"/Users/test/project",
			"/Users/test/project",
		},
		{
			"handles trailing slash in home",
			"/Users/test/",
			"/Users/test/project",
			"~/project",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("HOME", tt.home)
			result := FormatPath(tt.path)
			if result != tt.expected {
				t.Errorf("FormatPath(%q) with HOME=%q = %q, want %q",
					tt.path, tt.home, result, tt.expected)
			}
		})
	}
}

func TestPrintAgentReturnsColoredNames(t *testing.T) {
	// Testing color output is tricky because fatih/color auto-detects
	// terminal support. In tests it usually disables colors.
	// We mainly verify the function doesn't panic and returns something.

	agents := []string{"forge", "axiom", "jarvis", "unknown"}

	for _, agent := range agents {
		t.Run(agent, func(t *testing.T) {
			result := PrintAgent(agent)
			// At minimum, result should contain the agent name
			if !strings.Contains(result, agent) {
				t.Errorf("PrintAgent(%q) = %q, expected to contain agent name",
					agent, result)
			}
		})
	}
}
