#!/usr/bin/env bash

# Test script for agenter.sh
# 
# TEMPORARY: These bash scripts are prototypes. The production version
# will be a Go CLI with proper tests and type safety. Until then, this
# validates our bash implementation works correctly.

set -e  # Exit on error

echo "=== Testing Agenter Script ==="
echo

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Test function
test_case() {
    local name="$1"
    echo -n "Testing: $name... "
}

pass() {
    echo -e "${GREEN}PASS${NC}"
}

fail() {
    echo -e "${RED}FAIL${NC}: $1"
    exit 1
}

# Store the script location
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
SCRIPT_PATH="$SCRIPT_DIR/agenter.sh"

# Source the script
source "$SCRIPT_PATH"

# Create a test repository
TEST_DIR="/tmp/agenter-test-$$"
mkdir -p "$TEST_DIR"
cd "$TEST_DIR"

test_case "Create test repository"
git init test-repo >/dev/null 2>&1 || fail "Could not init repo"
cd test-repo
echo "test" > README.md
git add README.md
git commit -m "Initial commit" >/dev/null 2>&1 || fail "Could not commit"
pass

# Test 1: Agent functions should fail outside worktrees
test_case "Agent functions fail outside worktrees"
if forge 2>&1 | grep -q "ERROR.*forge.*-forge"; then
    pass
else
    fail "forge should fail outside worktree"
fi

# Test 2: Create worktrees
test_case "Create agent worktrees"
if create_agent_worktrees | grep -q "Created worktrees"; then
    pass
else
    fail "Could not create worktrees"
fi

# Test 3: Verify worktrees exist
test_case "Verify worktrees created"
if [[ -d "../test-repo-forge" && -d "../test-repo-axiom" && -d "../test-repo-jarvis" ]]; then
    pass
else
    fail "Worktrees not created"
fi

# Test 4: Run create_agent_worktrees again (should handle existing)
test_case "Handle existing worktrees"
if create_agent_worktrees | grep -q "already exists"; then
    pass
else
    fail "Should detect existing worktrees"
fi

# Test 5: Agent functions work in correct worktrees
test_case "Agent functions work in worktrees"
cd ../test-repo-forge
# The functions should already be available since we sourced the script
if type forge >/dev/null 2>&1; then
    pass
else
    fail "forge function not available"
fi

# Test 6: Topic management - make topic
test_case "Create topic branch"
if worktree_make_topic test-feature | grep -q "Now working on topic: test-feature"; then
    pass
else
    fail "Could not create topic"
fi

# Test 7: Push topic (without actual push to avoid remote requirement)
test_case "Push topic validates current branch"
# Should try to push but fail due to no remote - that's OK for testing
if worktree_push_topic 2>&1 | grep -q -E "(Pushing topic branch|fatal.*origin)"; then
    pass
else
    fail "Push topic failed"
fi

# Test 8: Next topic without name
test_case "Next topic returns to worktree branch"
if worktree_next_topic | grep -q "Ready for next topic"; then
    pass
else
    fail "Next topic failed"
fi

# Test 9: Next topic with name
test_case "Next topic with new name"
if worktree_next_topic another-feature | grep -q "Now working on topic: another-feature"; then
    pass
else
    fail "Next topic with name failed"
fi

# Test 10: Error on worktree branch push
test_case "Cannot push from worktree branch"
git checkout forge-worktree 2>/dev/null
if worktree_push_topic 2>&1 | grep -q "No topic to push"; then
    pass
else
    fail "Should prevent pushing worktree branch"
fi

# Cleanup
cd "$TEST_DIR"
rm -rf test-repo test-repo-forge test-repo-axiom test-repo-jarvis

echo
echo "=== All tests passed! ==="
echo
echo "Manual testing suggestions:"
echo "1. In a real project with 'origin' remote:"
echo "   - Test that worktree_push_topic generates correct PR URL"
echo "   - Test git pull from main works correctly"
echo
echo "2. Test the actual Claude integration:"
echo "   - cd to each worktree and run forge/axiom/jarvis"
echo "   - Verify .claude/ directories are created separately"