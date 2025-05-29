package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// getCurrentBranch returns the current git branch name
func getCurrentBranch() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

// getWorktreeBranch returns the worktree branch name for the current directory
func getWorktreeBranch() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	dir := filepath.Base(cwd)
	for _, agent := range []string{"forge", "axiom", "jarvis"} {
		if strings.HasSuffix(dir, fmt.Sprintf("-%s", agent)) {
			return fmt.Sprintf("%s-worktree", agent), nil
		}
	}

	return "", fmt.Errorf("not in an agent worktree directory")
}

// runWorktreeMakeImpl creates a new topic branch
func runWorktreeMakeImpl(topic string) error {
	// Get the worktree branch
	worktreeBranch, err := getWorktreeBranch()
	if err != nil {
		return err
	}

	// Get current branch
	currentBranch, err := getCurrentBranch()
	if err != nil {
		return fmt.Errorf("could not get current branch: %v", err)
	}

	// Check if we're on the worktree branch
	if currentBranch != worktreeBranch {
		PrintWarning("Currently on topic branch: %s", currentBranch)
		PrintInfo("Run 'agenter worktree next' to return to base branch first")
		return fmt.Errorf("already on a topic branch")
	}

	// Create and checkout new topic branch
	branchName := fmt.Sprintf("%s-%s", worktreeBranch, topic)
	cmd := exec.Command("git", "checkout", "-b", branchName)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("could not create topic branch: %s", string(output))
	}

	PrintSuccess("Created topic branch: %s", branchName)
	PrintInfo("Now working on topic: %s", topic)
	return nil
}

// runWorktreePushImpl pushes the current topic branch
func runWorktreePushImpl() error {
	// Get current branch
	currentBranch, err := getCurrentBranch()
	if err != nil {
		return fmt.Errorf("could not get current branch: %v", err)
	}

	// Check if it's a worktree branch
	worktreeBranch, _ := getWorktreeBranch()
	if currentBranch == worktreeBranch {
		return fmt.Errorf("no topic to push. Create a topic branch first with 'agenter worktree make <topic>'")
	}

	PrintInfo("Pushing topic branch: %s", currentBranch)

	// Push the branch
	cmd := exec.Command("git", "push", "-u", "origin", currentBranch)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("could not push: %s", string(output))
	}

	// Get the remote URL
	cmd = exec.Command("git", "remote", "get-url", "origin")
	remoteOutput, err := cmd.Output()
	if err != nil {
		PrintSuccess("Branch pushed successfully")
		return nil
	}

	// Parse GitHub URL and generate PR link
	remote := strings.TrimSpace(string(remoteOutput))
	if strings.Contains(remote, "github.com") {
		// Convert git@github.com:owner/repo.git to https://github.com/owner/repo
		prURL := remote
		prURL = strings.Replace(prURL, "git@github.com:", "https://github.com/", 1)
		prURL = strings.TrimSuffix(prURL, ".git")
		prURL = fmt.Sprintf("%s/pull/new/%s", prURL, currentBranch)

		PrintSuccess("Branch pushed successfully")
		fmt.Println()
		PrintBold("Create PR at:")
		fmt.Println(prURL)
	} else {
		PrintSuccess("Branch pushed successfully")
	}

	return nil
}

// runWorktreeNextImpl returns to base branch and optionally creates new topic
func runWorktreeNextImpl(newTopic string) error {
	// Get the worktree branch
	worktreeBranch, err := getWorktreeBranch()
	if err != nil {
		return err
	}

	// Get current branch
	currentBranch, err := getCurrentBranch()
	if err != nil {
		return fmt.Errorf("could not get current branch: %v", err)
	}

	// If on worktree branch and new topic provided, just create it
	if currentBranch == worktreeBranch && newTopic != "" {
		return runWorktreeMakeImpl(newTopic)
	}

	// Ensure changes are committed or stashed
	cmd := exec.Command("git", "status", "--porcelain")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("could not check git status: %v", err)
	}

	if len(output) > 0 {
		PrintWarning("You have uncommitted changes")
		PrintInfo("Please commit or stash changes before switching branches")
		return fmt.Errorf("unsaved changes")
	}

	// Return to worktree branch
	cmd = exec.Command("git", "checkout", worktreeBranch)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("could not switch to base branch: %s", string(output))
	}

	PrintSuccess("Returned to base branch: %s", worktreeBranch)

	// Update from main
	PrintInfo("Updating from main...")
	cmd = exec.Command("git", "pull", "origin", "main")
	if output, err := cmd.CombinedOutput(); err != nil {
		PrintWarning("Could not pull from main: %s", string(output))
	}

	// Create new topic if provided
	if newTopic != "" {
		fmt.Println()
		return runWorktreeMakeImpl(newTopic)
	}

	PrintInfo("Ready for next topic. Use 'agenter worktree make <topic>' to start")
	return nil
}

// runWorktreeListImpl lists all agent worktrees
func runWorktreeListImpl() error {
	cmd := exec.Command("git", "worktree", "list")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("could not list worktrees: %v", err)
	}

	PrintHeader("Agent Worktrees")

	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	foundAgent := false

	for scanner.Scan() {
		line := scanner.Text()
		for _, agent := range []string{"forge", "axiom", "jarvis"} {
			if strings.Contains(line, fmt.Sprintf("-%s", agent)) {
				parts := strings.Fields(line)
				if len(parts) >= 3 {
					path := parts[0]
					branch := strings.Trim(parts[2], "[]")
					fmt.Printf("  %s: %s [%s]\n", PrintAgent(agent), FormatPath(path), branch)
					foundAgent = true
				}
				break
			}
		}
	}

	if !foundAgent {
		PrintInfo("No agent worktrees found")
		PrintInfo("Run 'agenter setup <repository>' to create them")
	}

	return nil
}

// runWorktreeCreateImpl creates all agent worktrees
func runWorktreeCreateImpl() error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("could not get current directory: %v", err)
	}

	if !HasGitRepository(cwd) {
		return fmt.Errorf("not in a git repository")
	}

	// Use the setup implementation
	return runSetupImpl(cwd)
}
