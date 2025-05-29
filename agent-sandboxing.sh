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

_launch_agent() {
    local agent_name="$1"
    shift  # Remove agent name from args
    
    local current_dir=$(basename "$PWD")
    if [[ ! "$current_dir" =~ -${agent_name}$ ]]; then
        echo "ERROR: ${agent_name} can only run in directories ending with '-${agent_name}'"
        echo "  Current: $PWD"
        echo "  Why: Claude stores history in ./.claude/ - mixing contexts breaks agents"
        echo "  Fix: cd to a ${agent_name} worktree first"
        return 1
    fi
    WHO_AM_I=$agent_name claude "$@"
}

# Your three AI team members - each gets their own isolated workspace
# We enforce that each agent can only work in their designated worktree
forge() { _launch_agent forge "$@"; }
axiom() { _launch_agent axiom "$@"; }
jarvis() { _launch_agent jarvis "$@"; }

# Set up git worktrees - separate directories where each agent can work independently
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
    
    # Create worktrees, checking if they already exist
    for agent in forge axiom jarvis; do
        if [[ -d "../${repo_name}-${agent}" ]]; then
            echo "✓ Worktree already exists: ../${repo_name}-${agent}"
        else
            git worktree add -b "${agent}-worktree" "../${repo_name}-${agent}" || {
                # If branch already exists, try without -b flag
                git worktree add "../${repo_name}-${agent}" "${agent}-worktree" || {
                    echo "ERROR: Failed to create worktree for ${agent}"
                }
            }
        fi
    done
    
    echo ""
    echo "✓ Created worktrees:"
    echo "  - ../${repo_name}-forge"
    echo "  - ../${repo_name}-axiom"
    echo "  - ../${repo_name}-jarvis"
    echo ""
    echo "Next steps:"
    echo "  1. cd ../${repo_name}-forge && forge"
    echo "  2. cd ../${repo_name}-axiom && axiom"
    echo "  3. cd ../${repo_name}-jarvis && jarvis"
}

# Helper: list all worktrees
list_agent_worktrees() {
    echo "Agent worktrees in this repository:"
    git worktree list | grep -E "(forge|axiom|jarvis)$" || echo "No agent worktrees found"
}

# ===================================================================
# Worktree Topic Management
# ===================================================================
#
# Each agent works in their own git worktree but still needs to create
# topic branches for specific tasks. These functions manage that workflow.
#
# The pattern mirrors how developers work: create a branch, make changes,
# push for review, then move to the next task. The difference is that
# each agent's branches originate from their own worktree branch rather
# than from main, maintaining isolation between agents.
#
# The workflow:
#   1. worktree_make_topic <name>  - Create a topic branch
#   2. worktree_push_topic         - Push the branch for PR
#   3. worktree_next_topic [name]  - Return to base and optionally start next
#
# This handles the git operations so agents can focus on their work
# rather than remembering command sequences.

# Helper: Extract agent name from current directory
_get_agent_from_dir() {
    local dir=$(basename "$PWD")
    if [[ "$dir" =~ -(forge|axiom|jarvis)$ ]]; then
        echo "${BASH_REMATCH[1]}"
    else
        echo ""
    fi
}

# Helper: Ensure we're in an agent worktree directory
_require_agent_worktree() {
    local agent=$(_get_agent_from_dir)
    if [[ -z "$agent" ]]; then
        echo "ERROR: Must run from an agent worktree (*-forge, *-axiom, *-jarvis)"
        return 1
    fi
    echo "$agent"
}

# Create a topic branch for the current agent's work
worktree_make_topic() {
    local topic_name="$1"
    [[ -z "$topic_name" ]] && { echo "Usage: worktree_make_topic <topic-name>"; return 1; }
    
    local agent=$(_require_agent_worktree) || return 1
    local worktree_branch="${agent}-worktree"
    
    echo "Creating topic branch '$topic_name' for agent $agent..."
    
    # Ensure on worktree branch
    local current_branch=$(git branch --show-current)
    if [[ "$current_branch" != "$worktree_branch" ]]; then
        git checkout "$worktree_branch" || return 1
    fi
    
    # Update and create topic
    git pull origin main --no-edit || echo "WARNING: Could not pull from main"
    git checkout -b "$topic_name" || {
        echo "ERROR: Failed to create '$topic_name' (may already exist)"
        return 1
    }
    
    echo "✓ Now working on topic: $topic_name"
    echo "  Next: worktree_push_topic"
}

# Push the current topic branch to origin
worktree_push_topic() {
    local agent=$(_require_agent_worktree) || return 1
    local current_branch=$(git branch --show-current)
    
    # Check if we're on a worktree branch (not a topic)
    if [[ "$current_branch" =~ -worktree$ ]]; then
        echo "ERROR: No topic to push. Create with: worktree_make_topic <name>"
        return 1
    fi
    
    echo "Pushing topic branch '$current_branch'..."
    git push -u origin "$current_branch" || return 1
    
    echo "✓ Pushed successfully!"
    
    # Try to construct GitHub PR URL
    local remote_url=$(git remote get-url origin 2>/dev/null)
    if [[ "$remote_url" =~ github.com[:/]([^/]+)/([^/]+)(\.git)?$ ]]; then
        echo "Create PR: https://github.com/${BASH_REMATCH[1]}/${BASH_REMATCH[2]%.git}/compare/$current_branch?expand=1"
    fi
    
    echo "Next: worktree_next_topic [new-topic-name]"
}

# Return to worktree branch and optionally start a new topic
worktree_next_topic() {
    local next_topic_name="$1"
    local agent=$(_require_agent_worktree) || return 1
    local worktree_branch="${agent}-worktree"
    local current_branch=$(git branch --show-current)
    
    # If already on worktree branch and no next topic
    if [[ "$current_branch" == "$worktree_branch" ]] && [[ -z "$next_topic_name" ]]; then
        git pull origin main --no-edit || echo "WARNING: Could not pull from main"
        echo "✓ Ready for next topic. Run: worktree_make_topic <name>"
        return 0
    fi
    
    # Return to worktree branch if needed
    if [[ "$current_branch" != "$worktree_branch" ]]; then
        git checkout "$worktree_branch" || return 1
    fi
    
    # Update from main
    git pull origin main --no-edit || echo "WARNING: Could not pull from main"
    
    # If next topic provided, create it
    if [[ -n "$next_topic_name" ]]; then
        worktree_make_topic "$next_topic_name"
    else
        echo "✓ Ready for next topic. Run: worktree_make_topic <name>"
    fi
}
