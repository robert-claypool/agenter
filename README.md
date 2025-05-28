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

# Launch agents
forge   # WHO_AM_I=forge claude
axiom   # WHO_AM_I=axiom claude
jarvis  # WHO_AM_I=jarvis claude
```
