package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetWorktreeBranchReturnsCorrectBranchName(t *testing.T) {
	// getWorktreeBranch looks at current directory name to determine
	// which agent we are. If we're in "myproject-forge", it returns
	// "forge-worktree" as the branch name.

	originalDir, _ := os.Getwd()
	defer os.Chdir(originalDir)

	tmpDir := t.TempDir()

	// Test valid agent directories
	validTests := []struct {
		directory  string
		wantBranch string
	}{
		{"project-forge", "forge-worktree"},
		{"myapp-axiom", "axiom-worktree"},
		{"test-jarvis", "jarvis-worktree"},
	}

	for _, test := range validTests {
		t.Run("valid/"+test.directory, func(t *testing.T) {
			dir := filepath.Join(tmpDir, test.directory)
			os.Mkdir(dir, 0755)
			os.Chdir(dir)

			branch, err := getWorktreeBranch()

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if branch != test.wantBranch {
				t.Errorf("got branch %q, want %q", branch, test.wantBranch)
			}
		})
	}

	// Test invalid directories (should return error)
	invalidTests := []struct {
		directory string
		reason    string
	}{
		{"project", "no agent suffix"},
		{"forge", "missing dash before agent name"},
		{"project-unknown", "unknown is not a valid agent"},
	}

	for _, test := range invalidTests {
		t.Run("invalid/"+test.directory, func(t *testing.T) {
			dir := filepath.Join(tmpDir, test.directory)
			os.Mkdir(dir, 0755)
			os.Chdir(dir)

			_, err := getWorktreeBranch()

			if err == nil {
				t.Errorf("expected error for %s (%s), but got none",
					test.directory, test.reason)
			}
		})
	}
}
