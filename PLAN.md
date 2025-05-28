# Agenter: Multi-Agent Claude Orchestration CLI

## Vision

Transform the prompt-engineering repository into "agenter" - a professional CLI tool that makes multi-agent AI development fast, safe, and delightful. What starts as scripts and configuration becomes a product that orchestrates multiple Claude instances working in parallel without context contamination.

## Core Problem

When running multiple Claude Code instances, they share `.claude/` directories containing conversation history. This causes agents to inherit each other's context, leading to confusion and unpredictable behavior. Our solution uses git worktrees and directory guards to sandbox each agent.

## Product Evolution

### Phase 1: MVP (Current Scripts → Basic CLI)
- Port existing bash functions to Go CLI
- Hard-coded 3 agents: Forge, Axiom, Jarvis  
- Commands: `init`, `check`, `setup`, `launch`
- Single binary distribution
- Enhanced bootstrap process

### Phase 2: Customization
- User-defined agent names
- Variable number of agents (2-6)
- Agent profiles/personalities
- Project-specific configurations
- `agenter.yaml` config files

### Phase 3: Advanced Features
- Agent communication protocols
- Task distribution strategies
- Progress visualization
- Health monitoring
- Integration with task management (GitHub Issues)

## Technical Design

### CLI Structure (using Cobra)
```
agenter
├── init          # Interactive first-time setup
├── check         # Validate prerequisites
├── setup <repo>  # Create worktrees for a repository
├── launch <agent> # Launch agent with sandboxing
├── list          # Show configured projects
├── status        # Health check for all agents
└── version       # Version info
```

### User Experience

```bash
# First-time setup
$ agenter init
Welcome to Agenter v0.1.0 - Multi-Agent Claude Orchestration

Checking prerequisites...
✓ Claude Code installed (v1.95.1)
✓ Git version 2.35.0 (worktrees supported)
✓ GitHub CLI installed

Would you like to install our optimized CLAUDE.md configuration? [Y/n] y
✓ Configuration installed to ~/.claude/CLAUDE.md

Setup complete! Next: run 'agenter setup <repository>' to configure a project.

# Project setup
$ agenter setup ~/git/myproject
Setting up myproject for multi-agent development...

Creating agent worktrees:
✓ Created myproject-forge (branch: forge-work)
✓ Created myproject-axiom (branch: axiom-work)  
✓ Created myproject-jarvis (branch: jarvis-work)

Ready! Launch agents with:
  cd ~/git/myproject-forge && agenter launch forge
  cd ~/git/myproject-axiom && agenter launch axiom
  cd ~/git/myproject-jarvis && agenter launch jarvis

# Daily use
$ cd ~/git/myproject-axiom
$ agenter launch axiom
✓ Directory verified: myproject-axiom
✓ Launching Claude as Axiom...
[Claude Code starts with WHO_AM_I=axiom]
```

### Configuration Storage

Store project configurations in `~/.agenter/`:
```
~/.agenter/
├── config.yaml         # Global settings
└── projects/
    └── myproject.yaml  # Project-specific config
```

### Error Handling

Follow gabel's pattern:
- Clear, actionable error messages
- Suggest fixes for common problems
- Never panic, always graceful exits
- Colored output for clarity

### Distribution

1. **Direct download**: Releases on GitHub
2. **Homebrew**: `brew install robert-claypool/tap/agenter`
3. **Go install**: `go install github.com/robert-claypool/agenter@latest`

## Implementation Steps

1. **Repository preparation**
   - Rename prompt-engineering → agenter
   - Restructure for Go project
   - Keep existing scripts in `legacy/`

2. **MVP Development**
   - Set up Go module with Cobra
   - Implement basic commands
   - Port bash logic to Go
   - Add tests

3. **Polish**
   - Colored output
   - Progress indicators  
   - Interactive prompts
   - Comprehensive help

4. **Documentation**
   - Installation guide
   - Quick start tutorial
   - Architecture explanation
   - Troubleshooting

## Success Metrics

- Setup time: < 2 minutes from install to first agent launch
- Zero confusion about agent identity/context
- Works on macOS, Linux (Windows later)
- Delightful enough that users recommend it

## Open Questions

1. How do agents communicate? (GitHub Issues, files, other?)
2. Config file format: YAML, TOML, or JSON?
3. How to handle agent personality/profile management?

## Next Steps

1. Validate plan with stakeholders
2. Create agenter repository
3. Scaffold Go CLI structure
4. Implement `check` command first
5. Iterate based on user feedback

