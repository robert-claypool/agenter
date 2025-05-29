# Agenter

Multi-agent orchestration for Claude. Safe parallel AI development with context isolation.

## Installation

### From Source

```bash
git clone https://github.com/robert-claypool/agenter.git
cd agenter
make build
sudo make install
```

### Direct Download (Coming Soon)

Binary releases will be available on the GitHub releases page.

## Quick Start

```bash
# First-time setup
agenter init

# Set up a repository for multi-agent development
agenter setup ~/git/myproject

# Launch agents in separate terminals
cd ~/git/myproject-forge && agenter launch forge
cd ~/git/myproject-axiom && agenter launch axiom
cd ~/git/myproject-jarvis && agenter launch jarvis
```

## Commands

### Core Commands

- `agenter init` - Interactive first-time setup
- `agenter check` - Validate prerequisites  
- `agenter setup <repo>` - Create agent worktrees
- `agenter launch <agent>` - Launch Claude as an agent
- `agenter list` - Show configured projects
- `agenter status` - Health check for all agents

### Worktree Commands

- `agenter worktree make <topic>` - Create topic branch
- `agenter worktree push` - Push branch and get PR URL
- `agenter worktree next [topic]` - Return to base, optionally start new topic
- `agenter worktree list` - List agent worktrees
- `agenter worktree create` - Create worktrees in current repo

## Multi-Agent Workflow

Agenter uses three agents (Forge, Axiom, Jarvis) with git worktrees:

```bash
# Setup
gh repo clone owner/project project
cd project
agenter setup .

# Launch agents (with protection)
cd ../project-forge && agenter launch forge   # Only works in *-forge/ directories
cd ../project-axiom && agenter launch axiom   # Only works in *-axiom/ directories
cd ../project-jarvis && agenter launch jarvis # Only works in *-jarvis/ directories
```

## Why Worktrees?

Claude Code creates `.claude/` in your working directory to store conversation history. Without isolation, agents share context and become confused. Git worktrees + directory guards ensure each agent maintains its own mental model.

## How to Use

See our guide [WORKFLOW.md](WORKFLOW.md) for orchestrating agents.

## Development

```bash
# Build
make build

# Run tests
make test

# Install locally
make install

# See all commands
make help
```

## Requirements

- Claude Code
- Git (with worktree support)
- GitHub CLI (`gh`) - authenticated
