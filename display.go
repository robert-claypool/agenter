package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

var (
	successColor = color.New(color.FgGreen)
	errorColor   = color.New(color.FgRed)
	warnColor    = color.New(color.FgYellow)
	infoColor    = color.New(color.FgCyan)
	boldColor    = color.New(color.Bold)
)

// PrintSuccess prints a success message with a checkmark
func PrintSuccess(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	successColor.Printf("✓ %s\n", msg)
}

// PrintError prints an error message with an X
func PrintError(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	errorColor.Printf("✗ %s\n", msg)
}

// PrintWarning prints a warning message
func PrintWarning(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	warnColor.Printf("⚠️  %s\n", msg)
}

// PrintInfo prints an info message
func PrintInfo(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	infoColor.Printf("ℹ %s\n", msg)
}

// PrintBold prints text in bold
func PrintBold(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	boldColor.Println(msg)
}

// PrintHeader prints a section header
func PrintHeader(text string) {
	fmt.Println()
	boldColor.Printf("=== %s ===\n", text)
	fmt.Println()
}

// PrintStep prints a step in a process
func PrintStep(step int, total int, text string) {
	fmt.Printf("[%d/%d] %s\n", step, total, text)
}

// PrintAgent prints an agent name with appropriate styling
func PrintAgent(agent string) string {
	switch agent {
	case "forge":
		return color.RedString(agent)
	case "axiom":
		return color.BlueString(agent)
	case "jarvis":
		return color.GreenString(agent)
	default:
		return agent
	}
}

// PrintCommand prints a command that will be executed
func PrintCommand(cmd string) {
	color.New(color.FgHiBlack).Printf("$ %s\n", cmd)
}

// FormatPath shortens a path for display
func FormatPath(path string) string {
	home := os.Getenv("HOME")
	if home != "" && home != "/" && strings.HasPrefix(path, home) {
		// Handle the case where HOME has trailing slash
		home = strings.TrimSuffix(home, "/")
		after := strings.TrimPrefix(path, home)
		if after == "" {
			return "~"
		}
		if !strings.HasPrefix(after, "/") {
			// Path matched home but isn't a subdirectory
			return path
		}
		return "~" + after
	}
	return path
}
