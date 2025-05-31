package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// TestAgenterE2E runs through all Agenter functionality with a real repository
func TestAgenterE2E(t *testing.T) {
	// Skip in CI environments
	if os.Getenv("CI") != "" {
		t.Skip("Skipping e2e test in CI")
	}

	// Create temp directory - clean up any existing one first
	tempDir := filepath.Join(".", "temp")
	
	// Remove any existing temp directory from previous runs
	os.RemoveAll(tempDir)
	
	err := os.MkdirAll(tempDir, 0755)
	if err != nil {
		t.Fatalf("Could not create temp dir: %v", err)
	}
	
	// Ensure cleanup happens even if test fails
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			t.Logf("Warning: Could not remove temp directory: %v", err)
		}
	}()

	// Use a small, popular repository - "hello" by GitHub (simple example repo)
	repoURL := "https://github.com/octocat/Hello-World.git"
	repoPath := filepath.Join(tempDir, "Hello-World")

	t.Logf("Test directory: %s", tempDir)

	// Clone the repository
	t.Run("Clone repository", func(t *testing.T) {
		cmd := exec.Command("git", "clone", repoURL, repoPath)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Could not clone repo: %v\nOutput: %s", err, output)
		}
		t.Logf("Cloned repository to %s", repoPath)
	})

	// Build agenter binary
	binaryPath := filepath.Join(tempDir, "agenter")
	absBinaryPath, _ := filepath.Abs(binaryPath)
	t.Run("Build agenter", func(t *testing.T) {
		cmd := exec.Command("go", "build", "-o", absBinaryPath, ".")
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Could not build agenter: %v\nOutput: %s", err, output)
		}
		t.Logf("Built agenter binary at %s", absBinaryPath)
	})

	// Helper to run agenter commands
	runAgenter := func(args ...string) (string, error) {
		cmd := exec.Command(absBinaryPath, args...)
		output, err := cmd.CombinedOutput()
		return string(output), err
	}

	// Test check command
	t.Run("Check requirements", func(t *testing.T) {
		output, err := runAgenter("check")
		if err != nil {
			// Check might fail if Claude is not installed, which is OK for testing
			t.Logf("Check output: %s", output)
			if strings.Contains(output, "Git is installed") {
				t.Log("Git check passed")
			}
			if strings.Contains(output, "Claude Code not found") {
				t.Log("Claude not installed (expected in test environment)")
			}
		} else {
			t.Logf("All checks passed: %s", output)
		}
	})

	// Test setup command
	t.Run("Setup repository", func(t *testing.T) {
		output, err := runAgenter("setup", repoPath)
		if err != nil {
			t.Fatalf("Setup failed: %v\nOutput: %s", err, output)
		}
		t.Logf("Setup output: %s", output)

		// Verify worktrees were created
		parentDir := filepath.Dir(repoPath)
		for _, agent := range []string{"forge", "axiom", "jarvis"} {
			worktreePath := filepath.Join(parentDir, fmt.Sprintf("Hello-World-%s", agent))
			if _, err := os.Stat(worktreePath); os.IsNotExist(err) {
				t.Errorf("Worktree for %s not created at %s", agent, worktreePath)
			} else {
				t.Logf("Verified worktree for %s exists", agent)
			}
		}
	})

	// Test worktree list command
	t.Run("List worktrees", func(t *testing.T) {
		// Change to one of the worktrees
		forgeDir := filepath.Join(filepath.Dir(repoPath), "Hello-World-forge")
		originalDir, _ := os.Getwd()
		err := os.Chdir(forgeDir)
		if err != nil {
			t.Fatalf("Could not change to forge directory: %v", err)
		}
		defer os.Chdir(originalDir)

		output, err := runAgenter("worktree", "list")
		if err != nil {
			t.Fatalf("Worktree list failed: %v\nOutput: %s", err, output)
		}
		t.Logf("Worktree list output: %s", output)

		// Verify output contains all agents
		for _, agent := range []string{"forge", "axiom", "jarvis"} {
			if !strings.Contains(output, agent) {
				t.Errorf("Worktree list missing %s", agent)
			}
		}
	})

	// Test worktree make command
	t.Run("Create topic branch", func(t *testing.T) {
		forgeDir := filepath.Join(filepath.Dir(repoPath), "Hello-World-forge")
		originalDir, _ := os.Getwd()
		err := os.Chdir(forgeDir)
		if err != nil {
			t.Fatalf("Could not change to forge directory: %v", err)
		}
		defer os.Chdir(originalDir)

		output, err := runAgenter("worktree", "make", "test-feature")
		if err != nil {
			t.Fatalf("Worktree make failed: %v\nOutput: %s", err, output)
		}
		t.Logf("Created topic branch: %s", output)

		// Verify we're on the new branch
		cmd := exec.Command("git", "branch", "--show-current")
		branchOutput, err := cmd.Output()
		if err != nil {
			t.Fatalf("Could not get current branch: %v", err)
		}
		currentBranch := strings.TrimSpace(string(branchOutput))
		expectedBranch := "forge-worktree-test-feature"
		if currentBranch != expectedBranch {
			t.Errorf("Expected branch %s, got %s", expectedBranch, currentBranch)
		} else {
			t.Logf("Verified on branch %s", currentBranch)
		}
	})

	// Test worktree next command
	t.Run("Return to base branch", func(t *testing.T) {
		forgeDir := filepath.Join(filepath.Dir(repoPath), "Hello-World-forge")
		originalDir, _ := os.Getwd()
		err := os.Chdir(forgeDir)
		if err != nil {
			t.Fatalf("Could not change to forge directory: %v", err)
		}
		defer os.Chdir(originalDir)

		output, err := runAgenter("worktree", "next")
		if err != nil {
			t.Fatalf("Worktree next failed: %v\nOutput: %s", err, output)
		}
		t.Logf("Returned to base: %s", output)

		// Verify we're back on the worktree branch
		cmd := exec.Command("git", "branch", "--show-current")
		branchOutput, err := cmd.Output()
		if err != nil {
			t.Fatalf("Could not get current branch: %v", err)
		}
		currentBranch := strings.TrimSpace(string(branchOutput))
		expectedBranch := "forge-worktree"
		if currentBranch != expectedBranch {
			t.Errorf("Expected branch %s, got %s", expectedBranch, currentBranch)
		} else {
			t.Logf("Verified on branch %s", currentBranch)
		}
	})

	// Test launch command (will fail without Claude, but tests the flow)
	t.Run("Test launch validation", func(t *testing.T) {
		// Try to launch from wrong directory
		output, err := runAgenter("launch", "forge")
		if err == nil {
			t.Error("Expected launch to fail from non-agent directory")
		} else {
			t.Logf("Launch correctly failed from wrong directory: %s", output)
		}

		// Try to launch from correct directory
		forgeDir := filepath.Join(filepath.Dir(repoPath), "Hello-World-forge")
		originalDir, _ := os.Getwd()
		err = os.Chdir(forgeDir)
		if err != nil {
			t.Fatalf("Could not change to forge directory: %v", err)
		}
		defer os.Chdir(originalDir)

		// Check if Claude is installed to avoid hanging the test
		checkOutput, _ := runAgenter("check")
		if strings.Contains(checkOutput, "Claude Code is installed") {
			t.Log("Claude is installed - skipping actual launch to avoid hanging test")
		} else {
			output, err = runAgenter("launch", "forge")
			// This will fail if Claude is not installed, which is expected
			if err != nil && strings.Contains(output, "claude not found") {
				t.Log("Launch correctly detected missing Claude")
			} else if err != nil {
				t.Logf("Launch failed (expected if Claude not installed): %s", output)
			}
		}
	})

	t.Log("E2E test completed successfully")
}