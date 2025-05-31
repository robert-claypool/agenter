package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// runCheckImpl implements the check command
func runCheckImpl() error {
	PrintHeader("Checking Tools")

	// Check Claude Code
	PrintStep(1, 4, "Checking Claude Code installation...")
	if err := IsClaudeInstalled(); err != nil {
		PrintError("Claude Code not found: %v", err)
		return err
	}
	PrintSuccess("Claude Code is installed")

	// Check Git
	PrintStep(2, 4, "Checking Git installation...")
	if err := IsGitInstalled(); err != nil {
		PrintError("Git not found: %v", err)
		return err
	}
	PrintSuccess("Git is installed with worktree support")

	// Check GitHub CLI
	PrintStep(3, 4, "Checking GitHub CLI...")
	if err := IsGitHubCLIAuthenticated(); err != nil {
		PrintError("GitHub CLI issue: %v", err)
		return err
	}
	PrintSuccess("GitHub CLI is installed and authenticated")

	// Check current directory
	PrintStep(4, 4, "Checking current directory...")
	cwd, _ := os.Getwd()
	if HasGitRepository(cwd) {
		PrintSuccess("Current directory is a git repository")
	} else {
		PrintInfo("Current directory is not a git repository")
	}

	fmt.Println()
	PrintSuccess("All checks passed!")
	return nil
}

// runInitImpl implements the init command
func runInitImpl() error {
	PrintBold("Welcome to Agenter v%s - Multi-Agent Claude Orchestration", Version)
	fmt.Println()

	// Run checks first
	if err := runCheckImpl(); err != nil {
		fmt.Println()
		PrintError("Some tools missing. Please install them and try again.")
		return err
	}

	// Check for CLAUDE.md
	fmt.Println()
	PrintStep(1, 2, "Checking Claude configuration...")

	claudeConfigPath := filepath.Join(os.Getenv("HOME"), ".claude", "CLAUDE.md")
	claudeMdPath := filepath.Join(filepath.Dir(os.Args[0]), "CLAUDE.md")

	if _, err := os.Stat(claudeConfigPath); os.IsNotExist(err) {
		PrintInfo("Claude configuration not found at %s", claudeConfigPath)

		// Check if CLAUDE.md exists in agenter directory
		if _, err := os.Stat(claudeMdPath); os.IsNotExist(err) {
			PrintWarning("CLAUDE.md not found in agenter directory")
			PrintInfo("Run bootstrap.sh to set up Claude configuration")
		} else {
			PrintInfo("Run bootstrap.sh to link CLAUDE.md to Claude's config directory")
		}
	} else {
		PrintSuccess("Claude configuration found at %s", claudeConfigPath)
	}

	// Provide next steps
	fmt.Println()
	PrintStep(2, 2, "Next steps...")
	PrintInfo("1. Run 'agenter setup <repository>' to configure a project")
	PrintInfo("2. Launch agents with 'agenter launch <agent>'")
	PrintInfo("3. Use 'agenter worktree' commands for git workflow")

	fmt.Println()
	PrintSuccess("Setup complete!")
	return nil
}

// runSetupImpl implements the setup command
func runSetupImpl(repoPath string) error {
	// Expand path
	if strings.HasPrefix(repoPath, "~") {
		home := os.Getenv("HOME")
		repoPath = filepath.Join(home, repoPath[2:])
	}

	absPath, err := filepath.Abs(repoPath)
	if err != nil {
		return fmt.Errorf("bad path: %v", err)
	}

	// Check if it's a git repository
	if !HasGitRepository(absPath) {
		return fmt.Errorf("%s is not a git repository", absPath)
	}

	PrintHeader(fmt.Sprintf("Setting up %s for multi-agent development", filepath.Base(absPath)))

	// Get the repository name
	repoName := filepath.Base(absPath)
	parentDir := filepath.Dir(absPath)

	// Create worktrees for each agent
	agents := []string{"forge", "axiom", "jarvis"}
	for i, agent := range agents {
		PrintStep(i+1, len(agents), fmt.Sprintf("Creating %s worktree...", agent))

		worktreePath := filepath.Join(parentDir, fmt.Sprintf("%s-%s", repoName, agent))
		branchName := fmt.Sprintf("%s-worktree", agent)

		// Check if worktree already exists
		if _, err := os.Stat(worktreePath); !os.IsNotExist(err) {
			PrintWarning("Worktree %s already exists", worktreePath)
			continue
		}

		// Create worktree
		cmd := exec.Command("git", "-C", absPath, "worktree", "add", "-b", branchName, worktreePath)
		if output, err := cmd.CombinedOutput(); err != nil {
			PrintError("Could not create worktree: %s", string(output))
			return err
		}

		PrintSuccess("Created %s", FormatPath(worktreePath))
	}

	// Print launch instructions
	fmt.Println()
	PrintBold("Ready! Launch agents with:")
	for _, agent := range agents {
		worktreePath := filepath.Join(parentDir, fmt.Sprintf("%s-%s", repoName, agent))
		fmt.Printf("  cd %s && agenter launch %s\n", FormatPath(worktreePath), agent)
	}

	return nil
}

// runLaunchImpl runs the launch command
func runLaunchImpl(agent string) error {
	// Validate agent name
	if err := IsKnownAgentName(agent); err != nil {
		return err
	}

	// Validate directory
	if err := IsInAgentWorkspace(agent); err != nil {
		return err
	}

	// Set environment variable
	os.Setenv("WHO_AM_I", agent)

	PrintSuccess("Launching Claude as %s...", PrintAgent(agent))

	// Find claude binary
	claudePath, err := FindClaudePath()
	if err != nil {
		return fmt.Errorf("claude not found: %v", err)
	}
	
	// Launch Claude Code
	cmd := exec.Command(claudePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
