package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Verifies GitHub CLI exists and has valid authentication token.
// Returns error if gh is missing or user hasn't run 'gh auth login'.
// We need authenticated gh for creating PRs and managing issues.
func IsGitHubCLIAuthenticated() error {
	cmd := exec.Command("gh", "auth", "status")
	output, err := cmd.CombinedOutput()
	if err != nil {
		if strings.Contains(string(output), "not logged in") {
			return fmt.Errorf("GitHub CLI is not authenticated. Run 'gh auth login' first")
		}
		return fmt.Errorf("GitHub CLI is not installed. Install it from https://cli.github.com")
	}
	return nil
}

// FindClaudePath locates the claude binary. First tries PATH, then common installation location.
func FindClaudePath() (string, error) {
	// First try PATH
	if path, err := exec.LookPath("claude"); err == nil {
		return path, nil
	}
	
	// Try standard macOS installation location
	claudePath := filepath.Join(os.Getenv("HOME"), ".claude", "local", "claude")
	if _, err := os.Stat(claudePath); err == nil {
		return claudePath, nil
	}
	
	return "", fmt.Errorf("claude not found")
}

// Checks for claude binary. First tries PATH, then common installation location.
// Claude Code stores conversation history in .claude/ directories,
// which is why we need workspace isolation for each agent.
func IsClaudeInstalled() error {
	path, err := FindClaudePath()
	if err != nil {
		return fmt.Errorf("Claude Code is not installed. Install it from https://claude.ai/code")
	}
	LogDebug("Found claude at: %s", path)
	return nil
}

// Verifies git is available. Worktrees require Git 2.5+ (2015),
// but checking version compatibility adds complexity for little benefit
// since Git 2.5 is 9 years old.
func IsGitInstalled() error {
	cmd := exec.Command("git", "--version")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("Git is not installed")
	}

	LogDebug("Git version: %s", strings.TrimSpace(string(output)))
	return nil
}

// Detects if path contains a git repository by looking for .git.
// In worktrees, .git is a file pointing to the main repo's git directory.
func HasGitRepository(path string) bool {
	if path == "" {
		return false
	}
	gitPath := filepath.Join(path, ".git")
	info, err := os.Stat(gitPath)
	if err != nil {
		return false
	}
	return info.IsDir() || info.Mode().IsRegular()
}

// Ensures agent name is one of our three known agents.
// We use fixed agent names to maintain consistent workspace isolation
// and prevent accidental context mixing.
func IsKnownAgentName(agent string) error {
	validAgents := []string{"forge", "axiom", "jarvis"}
	for _, valid := range validAgents {
		if agent == valid {
			return nil
		}
	}
	return fmt.Errorf("unknown agent name: %s (must be forge, axiom, or jarvis)", agent)
}

// Prevents agents from running in wrong directories by checking directory suffix.
// Critical for maintaining conversation isolation - if forge runs in axiom's
// directory, it inherits axiom's conversation history and context.
func IsInAgentWorkspace(agent string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %v", err)
	}

	dir := filepath.Base(cwd)
	expectedSuffix := fmt.Sprintf("-%s", agent)

	if !strings.HasSuffix(dir, expectedSuffix) {
		return fmt.Errorf("%s can only run in directories ending with '%s'", agent, expectedSuffix)
	}

	return nil
}
