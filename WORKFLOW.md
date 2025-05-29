# Agenter Workflow

Run three AI agents in parallel. Each handles different parts of your project.

## Quick Start

Open 3 terminals:

```bash
# Terminal 1
cd ~/git/project-jarvis && jarvis
"Jarvis, improve our linter config"

# Terminal 2  
cd ~/git/project-forge && forge
"Forge, ensure that a user in tenant A cannot fetch data belonging to tenant B"

# Terminal 3
cd ~/git/project-axiom && axiom
"Axiom, change all 301 redirects to HTTP 302"
```

Always use the agent's name in your prompt.

## Work First, Branch When Ready

Just like regular git workflow, you don't need a branch name until you know what you're building:

```bash
# Start exploring
"Jarvis, check why our tests are flaky"

# Once direction is clear, create topic branch
"Jarvis, let's call this 'fix-flaky-tests' and create a topic branch"
# Jarvis runs: worktree_make_topic fix-flaky-tests

# Or if you know upfront
"Forge, create topic 'add-rate-limiting' and implement rate limiting on all endpoints"
# Forge runs: worktree_make_topic add-rate-limiting
```

Topics are git branches. Each agent works from their own base branch (forge-worktree, axiom-worktree, jarvis-worktree) and creates topic branches from there.

## Branch Commands

```bash
worktree_make_topic feature-name    # Create topic branch
worktree_push_topic                 # Push for PR
worktree_next_topic [next-name]     # Return to base (and optionally start next)
```

## Agent Roles

You decide what each agent does. Split the work however it makes sense for your project, but be mindful of merge conflicts. The agents must collaborate with each other.

## Coordination

Agents communicate through GitHub Issues and PRs:

```
"Axiom, create an issue for Forge: Need /api/users endpoint"
"Forge, check PR #45 from Axiom for the UI changes"
```

## Full Example

```bash
# Start exploring
"Jarvis, look at our memory usage patterns"

# Create topic when ready
"Jarvis, create topic 'reduce-memory-usage' for these optimizations"
# Jarvis: worktree_make_topic reduce-memory-usage

# Work and commit
"Jarvis, apply the memory optimizations we discussed"

# Push for review
"Jarvis, push this for review"
# Jarvis: worktree_push_topic

# Move to next task
"Jarvis, return to base"
# Jarvis: worktree_next_topic
```

## Common Problems

`ERROR: forge can only run in directories ending with '-forge'`
→ You're in the wrong directory. Each agent needs their own worktree.

`ERROR: Failed to create topic`
→ Branch already exists or you're already on a topic. Run `worktree_next_topic` first.

`Can't push worktree branch`
→ Correct. Create a topic branch for your actual work.