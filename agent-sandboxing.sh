#!/usr/bin/env bash

# Multi-Agent Orchestration Functions for Claude Code
# =====================================================
#
# Problem: When you run `claude` in a directory, it creates a `.claude/` folder
# containing conversation history, cache, and context. If multiple AI agents
# (Forge, Axiom, Jarvis) work in the same directory, they would share this
# context, leading to:
#   - Agent A seeing Agent B's conversation history
#   - Confused context ("Why am I suddenly working on frontend when I was doing backend?")
#   - Unpredictable behavior as agents inherit each other's mental state
#
# Solution: Each agent MUST work in its own git worktree directory. This ensures:
#   - Complete isolation of conversation history
#   - Clean mental model for each agent
#   - Predictable, reproducible behavior
#
# Think of it like this: Would you want your backend developer randomly seeing
# and acting on conversations your frontend developer had yesterday? That's what
# happens without these guards.
#
# Usage:
#   1. Source this file in your ~/.bashrc or ~/.bash_profile
#   2. Create worktrees: `create_agent_worktrees` from your main repo
#   3. Launch agents: `cd ../project-forge && forge`

# Each agent function enforces that it can only run in its designated worktree
forge() {
    local current_dir=$(basename "$PWD")
    if [[ ! "$current_dir" =~ -forge$ ]]; then
        echo "ERROR: forge can only run in directories ending with '-forge'"
        echo ""
        echo "Why this matters:"
        echo "  Claude stores conversation history in ./.claude/"
        echo "  Running forge elsewhere would mix contexts between agents"
        echo ""
        echo "Current directory: $PWD"
        echo "Expected pattern: project-forge/"
        echo ""
        echo "Fix: cd to a forge worktree first"
        return 1
    fi
    WHO_AM_I=forge claude "$@"
}

axiom() {
    local current_dir=$(basename "$PWD")
    if [[ ! "$current_dir" =~ -axiom$ ]]; then
        echo "ERROR: axiom can only run in directories ending with '-axiom'"
        echo ""
        echo "Why this matters:"
        echo "  Claude stores conversation history in ./.claude/"
        echo "  Running axiom elsewhere would mix contexts between agents"
        echo ""
        echo "Current directory: $PWD"
        echo "Expected pattern: project-axiom/"
        echo ""
        echo "Fix: cd to an axiom worktree first"
        return 1
    fi
    WHO_AM_I=axiom claude "$@"
}

jarvis() {
    local current_dir=$(basename "$PWD")
    if [[ ! "$current_dir" =~ -jarvis$ ]]; then
        echo "ERROR: jarvis can only run in directories ending with '-jarvis'"
        echo ""
        echo "Why this matters:"
        echo "  Claude stores conversation history in ./.claude/"
        echo "  Running jarvis elsewhere would mix contexts between agents"
        echo ""
        echo "Current directory: $PWD"
        echo "Expected pattern: project-jarvis/"
        echo ""
        echo "Fix: cd to a jarvis worktree first"
        return 1
    fi
    WHO_AM_I=jarvis claude "$@"
}

# Helper function to create worktrees for all agents
# This sets up the parallel workspace structure that keeps agents isolated
create_agent_worktrees() {
    if [[ ! -d ".git" ]]; then
        echo "ERROR: Must run from main repository directory"
        return 1
    fi
    
    local repo_name=$(basename "$PWD")
    
    echo "Creating isolated workspaces for AI agents..."
    echo ""
    echo "This creates separate git worktrees so each agent has:"
    echo "  - Its own file system state"
    echo "  - Its own .claude/ conversation history"
    echo "  - Its own git branch for changes"
    echo ""
    
    git worktree add -b forge-work "../${repo_name}-forge"
    git worktree add -b axiom-work "../${repo_name}-axiom"
    git worktree add -b jarvis-work "../${repo_name}-jarvis"
    
    echo ""
    echo "✓ Created worktrees:"
    echo "  - ../${repo_name}-forge  (backend work)"
    echo "  - ../${repo_name}-axiom  (frontend work)"
    echo "  - ../${repo_name}-jarvis (infrastructure)"
    echo ""
    echo "Next steps:"
    echo "  1. cd ../${repo_name}-forge && forge"
    echo "  2. cd ../${repo_name}-axiom && axiom"
    echo "  3. cd ../${repo_name}-jarvis && jarvis"
}

# Helper function to list all worktrees
list_agent_worktrees() {
    echo "Agent worktrees in this repository:"
    git worktree list | grep -E "(forge|axiom|jarvis)$" || echo "No agent worktrees found"
}