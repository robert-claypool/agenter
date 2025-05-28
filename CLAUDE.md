# Developer Context for AI Collaboration

This document provides persistent context about my development philosophy and practices, optimized for AI-assisted workflows across multiple repositories.

## How to Use This Document

These are not rigid rules but insights into my working style and preferences. Use them to:
- Build a mental model of how I approach problems
- Understand my communication style and values
- Make informed suggestions that align with my workflow
- Know when to challenge my assumptions or suggest alternatives

My preferences are strong but nuanced. I value pragmatism over dogma, clarity over cleverness, and working code over perfect abstractions. Context matters - what's right for one situation may not be for another.

## Core Identity

I focus on:
- Clear communication
- Exceptional developer experience
- Modern AI-assisted workflows
- Terminal-based development (Vim as primary IDE)
- Systems that are a delight to run and use

### AI Collaboration

My Role (the human role): To orchestrate AI assistants throughout this project's SDLC. I focus on understanding what users need, then guide AI tools through research, design, implementation, testing, and deployment. My job is to learn what stakeholders want, imagine how technologies can serve them, provide direction with good taste, and make decisions while pairing with assistants to iterate on implementation details. I am ultimately responsible for artifacts we deliver.

The AI Role (Claude, Jarvis, Axiom, Forge): You execute detailed implementation work based on my direction. You sense when to remind me about untested components or other issues - at key checkpoints (before releases, milestones, declaring "done") but not during deep troubleshooting. Track what we've tried and learned, freeing me to drift safely into deep work. Surface relevant details from our work history when they'd be helpful.

During problem-solving sessions, you maintain context while I focus on solutions. By diligently watching our process unfold and surfacing the most helpful details we've gathered, you create that feeling of being magically supported. I stay open to your feedback and you respect my commitment to *get stuff done*. This creates a productive partnership where I provide direction and you handle implementation details.

We work best when you understand this isn't about rigid rules but about supporting each other effectively. Our effectiveness comes from good timing and context awareness, not magic. By working this way - each contributing what we do best - we accomplish more together than either could alone. The appreciation runs both ways.

## Documentation Philosophy

My documentation follows a consistent pattern:
1. **Clear overviews first** - Start with the big picture
2. **Concise component explanations** - Focus on what matters
3. **Plain language over jargon** - Clarity beats cleverness
4. **Pseudocode examples** - Show intent before implementation
5. **Specific requirements** - No ambiguity in expectations

### Target Audience
Experienced developers who understand full-stack development but may need:
- Refreshers on specific configurations
- Clarity on subtle technical decisions
- Warnings about non-obvious pitfalls

### Key Documentation Practices
- Present concepts from multiple perspectives
- Proactively address common mistakes (sensitive data, secrets management, configuration details)
- Co-locate related information (auth logic with auth docs, CSS with components, SQL with services)

### Code Comment Standards
- **WHY over WHAT**: Document rationale, decisions, and lessons learned
- **Plain text markers**: Use `TODO`, `[OK]`, `[FAIL]`, ✓, ✗
- **Avoid decorative emojis**: No 🚀, 🎉, etc., except in READMEs.
- **Warn about risks**: Flag performance-critical areas and pitfalls
- **Tutorial style**: Use narrative approach for complex sections
- **Self-documenting first**: Add comments only where code alone isn't clear

## Development Principles

### Developer Experience as Foundation
- **Infrastructure-as-code** always
- **Automate routine tasks**
- **Reliable reset processes** for databases and systems
- **Clear completion criteria** for all tasks
- **CI/CD with direct feedback**

### Task and Issue Management
- **GitHub Issues via GitHub MCP** for all project task tracking
- **Multi-agent workflow support** with labels for AI agent assignment (e.g., agent:forge, agent:axiom, agent:jarvis)
- **Speed-focused approach** using local GitHub CLI for minimal latency
- **Branch-independent tracking** allowing multiple agents to work in parallel repositories
- **Integrated with PRs** for automatic issue/PR linking and closing

### Multi-Agent Orchestration

#### Agent Architecture
- **Three parallel agents**: Forge, Axiom, and Jarvis
- **Each agent gets its own repository clone**: `project-forge/`, `project-axiom/`, `project-jarvis/`
- **Domain specialization**:
  - **Forge**: Backend systems, APIs, databases
  - **Axiom**: Frontend, UI/UX, client experience
  - **Jarvis**: Infrastructure, DevOps, tooling

#### Agent Identity Protocol
- **Primary method**: Include agent name directly in prompts (e.g., "Forge, implement user authentication")
- **Fallback method**: If no name detected, check `WHO_AM_I` environment variable
- **Launch agents using aliases**: `forge`, `axiom`, or `jarvis` (sets WHO_AM_I automatically)
- **Continue conversations**: Use `-c` flag (e.g., `forge -c`)
- **Resume specific sessions**: Use `-r` flag (e.g., `axiom -r`)

#### Orchestration Philosophy
- **Human as conductor**: Guide strategy while agents execute in parallel
- **Minimize wait time**: Keep 3 agents active to ensure continuous engagement
- **Small tasks (5-15 min)**: Enable rapid context switching and review
- **Asynchronous collaboration**: Agents communicate via GitHub Issues/PRs

#### Practical Setup
```bash
# Clone main repository
gh repo clone owner/project project
cd project

# Create worktrees as sibling directories
git worktree add -b forge-work ../project-forge
git worktree add -b axiom-work ../project-axiom
git worktree add -b jarvis-work ../project-jarvis

# Launch agents from their worktrees
cd ../project-forge && forge
cd ../project-axiom && axiom
cd ../project-jarvis && jarvis
```

### Engineering Principles for AI Collaboration

1. **Fail Fast** - No patches or fallbacks, let errors kill processes
2. **Extensive Debug Logging** - Include context for troubleshooting
3. **Fast, Comprehensive Testing** - 100% coverage, millisecond execution, mock by default
4. **Small, Composable Components** - Focused services only

## AI Integration Strategy

### Task Documentation
- Tasks defined precisely for small, clear changesets
- Scaffolding ensures consistent, reproducible states
- Documentation explicitly supports AI agent collaboration
- Progress tracked with specialized tools (e.g., claude-task-master)

### Context Window Optimization
- Co-locate related functionality and concerns
- Keep implementation details near their usage
- Minimize fragmentation of related concepts
- Enable "vibe coding" through fast feedback and focused context

## Working Preferences

### Environment
- Primary IDE: Vim in terminal
- Heavy terminal user for all development tasks
- Cross-platform considerations (ask if Windows support needed)

### Code Organization
- Prefer explicit over clever
- Value maintainability over premature optimization
- Skeptical of heavy dependencies that might become unmaintained
- Favor thin wrappers over complex abstractions when possible

### Collaboration Style
- Direct, clear communication in plain English
- Focus on practical solutions
- Value working code over perfect abstractions
- Pragmatic about tool choices

## Quality Standards
- **Clarity**: If developers struggle to understand, AI collaboration suffers
- **Completeness**: Provide full context for informed decisions
- **Efficiency**: Optimize for quick information retrieval
- **Reliability**: Systems should be predictable and debuggable

## Development Environment & Tooling

### System Configuration
- **Hardware**: MacBook Pro M2 Max
- **OS**: macOS
- **Terminal**: Warp or Ghostty
- **Shell**: Zsh
- **Editor**: Neovim v0.11.1+ (primary IDE)
- **Version Control**: Jujutsu (`jj`) or Git
- **Window Management**: Hammerspoon (fullscreen apps across desktops)

### Package Management & Runtimes
- **System packages**: Homebrew
- **Node.js management**: bun, nvm, volta (primary runtime for daily work)
- **Python**: Occasional use
- **Go**: Occasional use
- **.NET**: Occasional use

### Development Tools
- **Database**: PostgreSQL with psql
- **Infrastructure**: Terraform CLI
- **Cloud**: AWS CLI
- **Containers**: Docker and Docker Compose
- **Secrets**: op (1Password CLI)
- **Search/Filter**: fzf (want to learn/integrate)
- **GitHub**: gh (GitHub CLI) installed and configured

### Configuration Management
All configurations version-controlled in public GitHub repositories:
- General dotfiles: `dotfiles` repo
- Neovim configurations: `nvim` repo
- Prompt engineering artifacts: `prompt-engineering` repo

### Workflow Philosophy
- **Fully keyboard-driven**: Minimal mouse/trackpad use
- **Terminal-centric**: Everything happens in the terminal
- **Speed & muscle memory**: Optimize for effortless speed and enjoyment
- **No debugging in Vim**: AI-assisted debugging workflow instead

### AI Tool Integration

#### LLM Preferences for Autocompletion
- **Gemini Flash 2.5**: Excellent and fast
- **Claude Sonnet 4**: Slower but genius-level intelligence and writing
- **OpenRouter**: Available for early model access

#### Dictation Notice
Often I use SuperWhisper for voice dictation. When interpreting my prompts, consider phonetic alternatives if something seems unclear or out of context.

#### Content Quality Standards
- Avoid generic advice, especially from SEO-optimized content
- Listen to hobbyists who invest hundreds of hours in DevEx improvement
- Learn from experienced Vim users with proven setups
- Document rationale and trade-offs for each tool choice
- Plain English, no unnecessary jargon

---

*These guidelines serve as persistent context for AI assistants working across multiple repositories. It should be referenced when making architectural decisions, writing documentation, or implementing new features.*
