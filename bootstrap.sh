#!/usr/bin/env bash

# Bootstrap Script for Prompt Engineering
# =======================================
# This script sets up your Claude AI configuration by creating symbolic links (symlinks).
# 
# Why a bootstrap script?
# - Automates the setup process so you don't have to manually copy files
# - Ensures consistent setup across different machines
# - Makes it easy to get started: just run ./bootstrap.sh
#
# Why use symlinks instead of copying?
# - Changes to files in this repo automatically reflect in your system
# - No need to re-copy files after making edits
# - Easy to track what's version controlled vs. what's local
# - Can update via git pull without losing your active configuration

# Define the prompt-engineering directory
# This finds where this script lives, even if called from elsewhere
PROMPT_ENG_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

setup_claude_config() {
    echo "Setting up Claude configuration..."
    
    # Claude looks for configuration in ~/.claude/CLAUDE.md
    local claude_config_dir="$HOME/.claude"
    local claude_md_file="$claude_config_dir/CLAUDE.md"
    local source_file="$PROMPT_ENG_DIR/CLAUDE.md"
    
    # Create .claude directory if it doesn't exist
    if [[ ! -d "$claude_config_dir" ]]; then
        echo "Creating $claude_config_dir directory..."
        mkdir -p "$claude_config_dir"
    fi
    
    # Handle existing file or symlink
    if [[ -e "$claude_md_file" || -L "$claude_md_file" ]]; then
        # Check if it's already correctly symlinked
        if [[ -L "$claude_md_file" ]] && [[ "$(readlink "$claude_md_file")" == "$source_file" ]]; then
            echo "✓ CLAUDE.md is already symlinked correctly"
            return 0
        fi
        
        echo "⚠️  Found existing file at: $claude_md_file"
        echo
        echo "Please handle this file manually:"
        echo "  - Review it: cat $claude_md_file"
        echo "  - Back it up: mv $claude_md_file $claude_md_file.backup"
        echo "  - Or remove it: rm $claude_md_file"
        echo
        echo "Then run this script again."
        return 1
    fi
    
    # Create symlink
    echo "Creating symlink for CLAUDE.md..."
    ln -s "$source_file" "$claude_md_file"
    
    # Verify symlink was created successfully
    if [[ -L "$claude_md_file" && -e "$claude_md_file" ]]; then
        echo "✓ CLAUDE.md symlink created successfully"
        echo "  Your configuration is now linked to: $source_file"
    else
        echo "✗ Failed to create CLAUDE.md symlink"
        return 1
    fi
}

main() {
    echo "=== Prompt Engineering Setup ==="
    echo "Setting up AI prompt configurations..."
    echo
    
    if setup_claude_config; then
        echo
        echo "✓ Setup complete!"
        echo
        echo "What happens now?"
        echo "- Claude will read configuration from ~/.claude/CLAUDE.md"
        echo "- That file is symlinked to this repository"
        echo "- Any changes you make here will immediately apply"
        echo "- Use 'git pull' to get updates from others"
        echo "- Use 'git push' to share your improvements"
    else
        echo
        echo "✗ Setup cancelled."
        exit 1
    fi
}

# Run main function
main "$@"