package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	verbose bool
	debug   bool
)

var rootCmd = &cobra.Command{
	Use:     "agenter",
	Short:   "Multi-agent orchestration for Claude",
	Long:    "Agenter helps you run multiple Claude Code instances in parallel with isolated contexts using git worktrees.",
	Version: Version,
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Interactive first-time setup",
	Long:  "Initialize agenter by checking prerequisites and setting up Claude configuration.",
	Run:   runInit,
}

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Validate prerequisites",
	Long:  "Check that all required tools are installed and configured correctly.",
	Run:   runCheck,
}

var setupCmd = &cobra.Command{
	Use:   "setup <repository>",
	Short: "Create worktrees for a repository",
	Long:  "Set up a repository with agent worktrees for Forge, Axiom, and Jarvis.",
	Args:  cobra.ExactArgs(1),
	Run:   runSetup,
}

var launchCmd = &cobra.Command{
	Use:   "launch <agent>",
	Short: "Launch agent with sandboxing",
	Long:  "Launch Claude Code as a specific agent (forge, axiom, or jarvis) with environment isolation.",
	Args:  cobra.ExactArgs(1),
	Run:   runLaunch,
}

var worktreeCmd = &cobra.Command{
	Use:   "worktree",
	Short: "Git worktree management",
	Long:  "Manage git worktrees for agent development workflows.",
}

var worktreeMakeCmd = &cobra.Command{
	Use:   "make <topic>",
	Short: "Create topic branch",
	Long:  "Create a new topic branch from the current agent's base branch.",
	Args:  cobra.ExactArgs(1),
	Run:   runWorktreeMake,
}

var worktreePushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push topic, get PR URL",
	Long:  "Push the current topic branch and display the PR creation URL.",
	Run:   runWorktreePush,
}

var worktreeNextCmd = &cobra.Command{
	Use:   "next [topic]",
	Short: "Return to base, start new topic",
	Long:  "Return to the agent's base branch and optionally start a new topic.",
	Args:  cobra.MaximumNArgs(1),
	Run:   runWorktreeNext,
}

var worktreeListCmd = &cobra.Command{
	Use:   "list",
	Short: "List agent worktrees",
	Long:  "Display all agent worktrees for the current repository.",
	Run:   runWorktreeList,
}

var worktreeCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create agent worktrees",
	Long:  "Create worktrees for all agents (forge, axiom, jarvis) in the current repository.",
	Run:   runWorktreeCreate,
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show configured projects",
	Long:  "Display all projects configured with agenter.",
	Run:   runList,
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Health check for all agents",
	Long:  "Check the status and health of all configured agents.",
	Run:   runStatus,
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Enable debug logging")

	// Add subcommands
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(checkCmd)
	rootCmd.AddCommand(setupCmd)
	rootCmd.AddCommand(launchCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(statusCmd)

	// Add worktree subcommands
	worktreeCmd.AddCommand(worktreeMakeCmd)
	worktreeCmd.AddCommand(worktreePushCmd)
	worktreeCmd.AddCommand(worktreeNextCmd)
	worktreeCmd.AddCommand(worktreeListCmd)
	worktreeCmd.AddCommand(worktreeCreateCmd)
	rootCmd.AddCommand(worktreeCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// Command implementations
func runInit(cmd *cobra.Command, args []string) {
	InitLogger(debug)
	if err := runInitImpl(); err != nil {
		os.Exit(1)
	}
}

func runCheck(cmd *cobra.Command, args []string) {
	InitLogger(debug)
	if err := runCheckImpl(); err != nil {
		os.Exit(1)
	}
}

func runSetup(cmd *cobra.Command, args []string) {
	InitLogger(debug)
	if err := runSetupImpl(args[0]); err != nil {
		PrintError("Setup failed: %v", err)
		os.Exit(1)
	}
}

func runLaunch(cmd *cobra.Command, args []string) {
	InitLogger(debug)
	if err := runLaunchImpl(args[0]); err != nil {
		PrintError("Launch failed: %v", err)
		os.Exit(1)
	}
}

func runWorktreeMake(cmd *cobra.Command, args []string) {
	InitLogger(debug)
	if err := runWorktreeMakeImpl(args[0]); err != nil {
		PrintError("Failed to create topic: %v", err)
		os.Exit(1)
	}
}

func runWorktreePush(cmd *cobra.Command, args []string) {
	InitLogger(debug)
	if err := runWorktreePushImpl(); err != nil {
		PrintError("Failed to push: %v", err)
		os.Exit(1)
	}
}

func runWorktreeNext(cmd *cobra.Command, args []string) {
	InitLogger(debug)
	newTopic := ""
	if len(args) > 0 {
		newTopic = args[0]
	}
	if err := runWorktreeNextImpl(newTopic); err != nil {
		PrintError("Failed: %v", err)
		os.Exit(1)
	}
}

func runWorktreeList(cmd *cobra.Command, args []string) {
	InitLogger(debug)
	if err := runWorktreeListImpl(); err != nil {
		PrintError("Failed to list worktrees: %v", err)
		os.Exit(1)
	}
}

func runWorktreeCreate(cmd *cobra.Command, args []string) {
	InitLogger(debug)
	if err := runWorktreeCreateImpl(); err != nil {
		PrintError("Failed to create worktrees: %v", err)
		os.Exit(1)
	}
}

func runList(cmd *cobra.Command, args []string) {
	InitLogger(debug)
	PrintInfo("Project listing not yet implemented")
	PrintInfo("This will show all projects configured with agenter")
}

func runStatus(cmd *cobra.Command, args []string) {
	InitLogger(debug)
	PrintInfo("Status command not yet implemented")
	PrintInfo("This will show health status of all agents")
}
