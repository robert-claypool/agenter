package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIsKnownAgentNameAcceptsValidAgents(t *testing.T) {
	tests := []struct {
		name    string
		agent   string
		wantErr bool
	}{
		{"accepts forge", "forge", false},
		{"accepts axiom", "axiom", false},
		{"accepts jarvis", "jarvis", false},
		{"rejects unknown agent", "invalid", true},
		{"rejects empty string", "", true},
		{"rejects uppercase FORGE", "FORGE", true},
		{"rejects mixed case Forge", "Forge", true},
		{"rejects forge with spaces", " forge ", true},
		{"rejects partial match", "forg", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := IsKnownAgentName(tt.agent)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsKnownAgentName(%q) error = %v, wantErr %v", tt.agent, err, tt.wantErr)
			}
		})
	}
}

func TestIsInAgentWorkspaceEnforcesDirectoryNaming(t *testing.T) {
	// Tests change working directory which affects other tests
	originalDir, _ := os.Getwd()
	defer os.Chdir(originalDir)

	tmpDir := t.TempDir()
	testCases := []struct {
		dirName string
		dirs    map[string]bool // agent -> should pass
	}{
		{
			"project-forge",
			map[string]bool{"forge": true, "axiom": false, "jarvis": false},
		},
		{
			"myapp-axiom",
			map[string]bool{"forge": false, "axiom": true, "jarvis": false},
		},
		{
			"test-jarvis",
			map[string]bool{"forge": false, "axiom": false, "jarvis": true},
		},
		{
			"project", // no agent suffix
			map[string]bool{"forge": false, "axiom": false, "jarvis": false},
		},
		{
			"forge", // missing dash
			map[string]bool{"forge": false, "axiom": false, "jarvis": false},
		},
		{
			"project-forger", // wrong suffix
			map[string]bool{"forge": false, "axiom": false, "jarvis": false},
		},
	}

	for _, tc := range testCases {
		testDir := filepath.Join(tmpDir, tc.dirName)
		os.Mkdir(testDir, 0755)

		for agent, shouldPass := range tc.dirs {
			t.Run(tc.dirName+"/"+agent, func(t *testing.T) {
				os.Chdir(testDir)
				err := IsInAgentWorkspace(agent)

				if shouldPass && err != nil {
					t.Errorf("IsInAgentWorkspace(%q) in %q: unexpected error: %v",
						agent, tc.dirName, err)
				}
				if !shouldPass && err == nil {
					t.Errorf("IsInAgentWorkspace(%q) in %q: expected error but got none",
						agent, tc.dirName)
				}
			})
		}
	}
}

func TestHasGitRepositoryDetectsGitDirectories(t *testing.T) {
	tmpDir := t.TempDir()

	gitRepo := filepath.Join(tmpDir, "git-repo")
	os.Mkdir(gitRepo, 0755)
	os.Mkdir(filepath.Join(gitRepo, ".git"), 0755)

	worktree := filepath.Join(tmpDir, "worktree")
	os.Mkdir(worktree, 0755)
	os.WriteFile(filepath.Join(worktree, ".git"), []byte("gitdir: ../git-repo/.git"), 0644)

	nonGit := filepath.Join(tmpDir, "non-git")
	os.Mkdir(nonGit, 0755)

	tests := []struct {
		name     string
		path     string
		expected bool
	}{
		{"detects .git directory", gitRepo, true},
		{"detects .git file in worktree", worktree, true},
		{"returns false for non-git directory", nonGit, false},
		{"returns false for non-existent path", filepath.Join(tmpDir, "does-not-exist"), false},
		{"returns false for empty path", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HasGitRepository(tt.path)
			if result != tt.expected {
				t.Errorf("HasGitRepository(%q) = %v, want %v", tt.path, result, tt.expected)
			}
		})
	}
}

// Skip to avoid CI failures when tools aren't installed
func TestIsClaudeInstalledSkipped(t *testing.T) {
	t.Skip("Requires claude binary")
}

func TestIsGitInstalledSkipped(t *testing.T) {
	t.Skip("Requires git binary")
}

func TestIsGitHubCLIAuthenticatedSkipped(t *testing.T) {
	t.Skip("Requires gh binary")
}
