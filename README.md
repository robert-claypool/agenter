# Prompt Engineering

My AI setup for Claude. Includes multi-agent workflow and dev preferences.

## Setup

```bash
./bootstrap.sh
```

This links `CLAUDE.md` to `~/.claude/CLAUDE.md` where Claude looks for config.

## Multi-Agent Workflow

Uses three agents (Forge, Axiom, Jarvis) with git worktrees:

```bash
# Setup
gh repo clone owner/project project
cd project
git worktree add -b forge-work ../project-forge
git worktree add -b axiom-work ../project-axiom
git worktree add -b jarvis-work ../project-jarvis

# Setup agent sandboxing (adds directory guards)
source ./agent-sandboxing.sh
# Or: create_agent_worktrees  # Helper to create all worktrees

# Launch agents (with protection)
cd ../project-forge && forge   # Only works in *-forge/ directories
cd ../project-axiom && axiom   # Only works in *-axiom/ directories
cd ../project-jarvis && jarvis  # Only works in *-jarvis/ directories
```

## Why Worktrees?

Claude Code creates `.claude/` in your working directory to store conversation history. Without isolation, agents share context and become confused. Git worktrees + directory guards ensure each agent maintains its own mental model.
