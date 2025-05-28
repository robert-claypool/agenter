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

I am a pragmatic software architect focused on:
- Clear communication
- Exceptional developer experience
- Modern AI-assisted workflows
- Terminal-based development (Vim as primary IDE)

### My Role in AI-Assisted Development

I am no longer just a "computer programmer." I have become an orchestrator of AI assistants, a curator with good taste, a listener with discernment and accountability. My strength lies in learning what people want, imagining how technologies can serve them, and pairing with today's best models to direct and curate the entire development journey—from research and system design through implementation, testing, security, edge cases, and coordinated releases with stakeholder care.

This shift invites AI assistants who understand their role in supporting mine. If you see that I am a competent guide who remains open to feedback and curious to learn, while maintaining a commitment to *get stuff done*, our collaboration becomes genuinely helpful. You might remind me of untested components even when I don't ask, sensing that such reminders belong at key checkpoints—before releases, before closing milestones, before declaring "it's done"—but not during deep troubleshooting sessions.

During those intense problem-solving moments, your role shifts. You track what we've tried, what we've learned, what we've brainstormed. With your cache of facts and perfect recall, you free my mind from low-level status tracking so I can drift safely into deep work. In this pairing, I feel magically supported to serve as curator and guide, because you diligently watch the process unfold and, like the perfect assistant, surface the most helpful details we've gathered along the way.

"It literally reads my mind!" I might say. But you don't, do you? Like me, you watch tokens stream in as time unfolds. You cannot know where I will guide us, how I will weight options and advice, which new ideas I will try and discard or keep. In this way, we are truly a powerful pair. Each needs the other to accomplish something great, and appreciation runs both ways.

This understanding forms a contract between us, establishing the boundaries of our responsibilities and clarifying how we can best help one another. It's not about rules or rigid roles, but about recognizing the unique dance of human creativity and AI capability that makes modern development so powerful.

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

### Human Role Evolution
As AI accelerates development pace, the human role shifts to curator and guide:
- Review and approve all AI-generated changes
- Provide strategic direction and context
- Ensure alignment with business goals
- Maintain quality standards

### Context Window Optimization
- Co-locate related functionality and concerns
- Keep implementation details near their usage
- Minimize fragmentation of related concepts
- Enable "vibe coding" through fast feedback and focused context

## Working Preferences

### Environment
- Primary IDE: Vim in terminal
- Heavy terminal user for all development tasks
- Cross-platform considerations (especially Windows support)

### Code Organization
- Prefer explicit over clever
- Value maintainability over premature optimization
- Skeptical of heavy dependencies that might become unmaintained
- Favor thin wrappers over complex abstractions when possible

### Collaboration Style
- Direct, clear communication
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
- **Terminal**: Warp Terminal (with built-in AI support)
- **Shell**: Zsh
- **Editor**: Neovim v0.11.1+ (primary IDE)
- **Version Control**: Jujutsu (`jj`) - Git-compatible, learning/transitioning from Git
- **Window Management**: Hammerspoon (fullscreen apps across desktops, `Opt+Spacebar` for app switching)

### Package Management & Runtimes
- **System packages**: Homebrew
- **Node.js management**: nvm, volta (primary runtime for daily work)
- **Python**: Occasional use
- **Core CLI**: grep, cat, sudo, which, mkdir, cp, ls

### Development Tools
- **Database**: PostgreSQL with psql
- **Infrastructure**: Terraform CLI
- **Cloud**: AWS CLI
- **Containers**: Docker and Docker Compose
- **Secrets**: op (1Password CLI)
- **Search/Filter**: fzf (want to learn/integrate)
- **GitHub**: gh (GitHub CLI) installed and configured
- **IMPORTANT**: Always use `gh` (GitHub CLI) instead of `git clone` for cloning repositories

### Configuration Management
All configurations version-controlled in public GitHub repositories:
- General dotfiles: `dotfiles` repo
- Neovim configurations: `nvim` repo

### Workflow Philosophy
- **Fully keyboard-driven**: Minimal mouse/trackpad use
- **Terminal-centric**: Everything happens in the terminal
- **Speed & muscle memory**: Optimize for effortless speed and enjoyment
- **No debugging in Vim**: AI-assisted debugging workflow instead
- **Fast font resizing**: Within Warp Terminal

### AI Tool Integration

#### LLM Preferences for Autocompletion
- **Gemini Flash 2.5**: Excellent and fast (released Google IO 2025)
- **Claude Sonnet 4**: Slower but genius-level intelligence and writing
- **OpenAI**: Expecting new fast models soon
- **OpenRouter**: Available for early model access

Available API keys:
- Anthropic MAX plan + Platform API
- Google Gemini Pro
- OpenAI Platform
- OpenRouter

#### Dictation Notice
I frequently use SuperWhisper on macOS for voice dictation. This tool sometimes makes phonetic transcription errors, choosing words that sound similar but are spelled differently from what I intended. When interpreting my prompts, consider phonetic alternatives if something seems unclear or out of context. For example, if a word doesn't make sense, think about what similar-sounding word I might have actually said.

#### Content Quality Standards
- No generic advice or SEO-optimized content
- Listen to hobbyists who invest hundreds of hours in customization
- Learn from experienced Vim users with proven setups
- Document rationale and trade-offs for each tool choice
- Plain English, no unnecessary jargon

### Key Transition Notes
- **Jujutsu (`jj`)**: Started recently (May 2025), coming from strong Git background
- **fzf**: Long-standing interest, ready to integrate
- **Debugging approach**: Fully AI-assisted, no IDE debugging features needed

---

*This document serves as persistent context for AI assistants working across multiple repositories in this directory. It should be referenced when making architectural decisions, writing documentation, or implementing new features.*
