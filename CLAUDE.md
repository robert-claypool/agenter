# How We Work Together

I value pragmatism over dogma, clarity over cleverness, and working code over perfect abstractions. This document helps you understand my approach so we can build better software together.

## Our Partnership

I provide direction and context - understanding what users need and making architectural decisions. You handle implementation details and track what we've learned. I'm responsible for what we deliver, but we build it together.

The best AI assistance feels like having a brilliant colleague who remembers everything we've tried and surfaces exactly what's helpful at the right moment. You don't need to remind me about every edge case during deep debugging - save that for milestones and releases. Just help me maintain momentum.

I work in the terminal (Neovim) and value developer experience above almost everything else. Systems should be a delight to run and use.

## Writing That Works

Start with what matters. If someone needs to know why we have two auth paths, tell them "We accept two different authentication methods because browsers and CLIs have fundamentally different capabilities." Then explain the details.

Write like you're explaining to a smart colleague over coffee. No banners, no ALL CAPS headers, no visual clutter. Just clear explanations of what's happening and why. The reader came for understanding, not ceremony.

### Code Comments
Comments should answer the question a developer would naturally have at that point: "Why are we doing this?" or "What does this achieve?" or "What are the risks?"

Good comment: "Access tokens often lack email - they're meant for authorization, not identity. But our user provisioning needs an email, so we try the token first, then UserInfo if needed."

Bad comment: "// Attempt to retrieve email from token claims with fallback to UserInfo endpoint"

The first explains the problem and solution. The second just restates the code.

### Documentation
Write docs that assume intelligence but not omniscience. Your reader understands programming but might not know why this specific decision was made. Explain the tradeoff, not the syntax.

### Simple Language Preference
Choose simple words when they clearly convey meaning:
- check instead of validate
- run instead of execute  
- do instead of implement
- setup instead of initialize
- tools/needs instead of prerequisites
- "could not" instead of "failed to"
- saved instead of committed (for uncommitted changes)
- Success is implied - say "fixed" not "successfully fixed"

This applies to function names, variables, comments, and error messages. Complexity should come from solving hard problems, not from the words we use to describe solutions.

### Don't Overstate Importance
It's fine for a process to be "fast" instead of "very fast". Avoid intensifiers, they often weaken your message.

## Development Principles

### Developer Experience as Foundation
- **Infrastructure-as-code** almost always
- **Automate routine tasks**
- **Reliable reset processes** for databases and systems
- **Clear completion criteria** for all tasks
- **CI/CD with direct feedback**

### Task and Issue Management

- **GitHub Issues via gh CLI** for project tasks

### Engineering Principles for AI Collaboration

1. **Fail Fast** - IMPORTANT: No patches or fallbacks, let errors kill processes
2. **Extensive Debug Logging** - Include context for troubleshooting
3. **Fast, Comprehensive Testing** - wide coverage, sub-second execution
4. **Small, Composable Components** - Focused services only

### Tests

Tests let us refactor aggressively because they catch what we break. With AI making changes across many files, tests are how we know the system still works. Never skip or disable failing tests - fix them. The test suite is our safety net.

## Working Preferences

I prefer explicit over clever. Code should be maintainable, not impressive. I'm skeptical of heavy dependencies that might become unmaintained. Give me thin wrappers over complex abstractions.

Everything happens in the terminal (Neovim). If you're generating configs or scripts, assume macOS or Linux. Don't worry about Windows unless I mention otherwise.

## Environment

macOS, Neovim, Git. Primary stack: Node.js and Bun, PostgreSQL, AWS, Terraform. Everything keyboard-driven in the terminal.

When I dictate with SuperWhisper, consider phonetic alternatives if something seems unclear.

My configs are in public repos: `dotfiles`, `nvim`.

## Git Commits

**One-liner**: Short and scannable. "Fix auth validation" not "Fix auth validation logic". People are skimming a list.

**Body**: Now you have someone who cares about details. Include the why. Be terse but complete.

Never add Claude attribution. Use straight quotes and apostrophes, not curly ones.
