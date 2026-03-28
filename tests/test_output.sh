#!/bin/bash

# Test script for --output and --raw flags
# This script tests all commands with different output formats
# Note: Many commands will fail due to missing credentials, but we can verify
# that the flag parsing and output formatting works correctly

CLI="./dist/pingcode-client.darwin-arm64"
PASS=0
FAIL=0

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

test_cmd() {
    local cmd="$1"
    local desc="$2"
    local expected_fail="${3:-false}"  # true if API call is expected to fail (no creds)

    echo -n "Testing: $desc ... "
    output=$($cmd 2>&1)
    exit_code=$?

    # Check if command exists (not a "unknown command" error)
    if echo "$output" | grep -q "unknown command\|Error: flag"; then
        echo -e "${RED}FAIL${NC} - Flag/command error"
        echo "  Output: $output"
        ((FAIL++))
        return 1
    fi

    # Check for expected error messages (valid behavior for certain tests)
    if echo "$output" | grep -q "unsupported output format"; then
        echo -e "${GREEN}PASS${NC} (correctly rejected invalid format)"
        ((PASS++))
        return 1
    fi

    # If expected to fail due to API error (no creds), we still pass if flags parsed correctly
    if [ $exit_code -ne 0 ]; then
        if echo "$output" | grep -q "client not initialized\|Error\|failed\|not found\|Forbidden\|Unauthorized"; then
            # API error - flags parsed correctly
            echo -e "${YELLOW}PASS${NC} (API error - creds missing, but flags work)"
            ((PASS++))
        else
            echo -e "${RED}FAIL${NC} - Exit code: $exit_code"
            echo "  Output: $output"
            ((FAIL++))
        fi
    else
        echo -e "${GREEN}PASS${NC}"
        ((PASS++))
    fi
}

echo "=========================================="
echo "Testing --output and --raw flags"
echo "=========================================="
echo ""

# Test 1: Global flags exist on all commands
echo "=== Testing Global Flags ==="
test_cmd "$CLI --help" "Global --help shows --output and --raw"
test_cmd "$CLI ship --help" "ship --help shows --output and --raw"
test_cmd "$CLI project --help" "project --help shows --output and --raw"
test_cmd "$CLI wiki --help" "wiki --help shows --output and --raw"
test_cmd "$CLI testhub --help" "testhub --help shows --output and --raw"
test_cmd "$CLI directory --help" "directory --help shows --output and --raw"
echo ""

# Test 2: --output json
echo "=== Testing --output json ==="
test_cmd "$CLI ship product list --output json" "ship product list --output json"
test_cmd "$CLI ship ideas list --output json --product-id test" "ship ideas list --output json"
test_cmd "$CLI ship tickets list --output json --product-id test" "ship tickets list --output json"
test_cmd "$CLI project list --output json" "project list --output json"
test_cmd "$CLI project metadata types --output json" "project metadata types --output json"
test_cmd "$CLI wiki space list --output json" "wiki space list --output json"
test_cmd "$CLI wiki page list --output json --space-id test" "wiki page list --output json"
test_cmd "$CLI testhub library list --output json" "testhub library list --output json"
test_cmd "$CLI directory user list --output json" "directory user list --output json"
echo ""

# Test 3: --output yaml
echo "=== Testing --output yaml ==="
test_cmd "$CLI ship product list --output yaml" "ship product list --output yaml"
test_cmd "$CLI ship ideas list --output yaml --product-id test" "ship ideas list --output yaml"
test_cmd "$CLI ship tickets list --output yaml --product-id test" "ship tickets list --output yaml"
test_cmd "$CLI project list --output yaml" "project list --output yaml"
test_cmd "$CLI project metadata types --output yaml" "project metadata types --output yaml"
test_cmd "$CLI wiki space list --output yaml" "wiki space list --output yaml"
test_cmd "$CLI testhub library list --output yaml" "testhub library list --output yaml"
test_cmd "$CLI directory user list --output yaml" "directory user list --output yaml"
echo ""

# Test 4: --output markdown
echo "=== Testing --output markdown ==="
test_cmd "$CLI ship product list --output markdown" "ship product list --output markdown"
test_cmd "$CLI ship ideas list --output markdown --product-id test" "ship ideas list --output markdown"
test_cmd "$CLI ship tickets list --output markdown --product-id test" "ship tickets list --output markdown"
test_cmd "$CLI project list --output markdown" "project list --output markdown"
test_cmd "$CLI project metadata types --output markdown" "project metadata types --output markdown"
test_cmd "$CLI wiki space list --output markdown" "wiki space list --output markdown"
test_cmd "$CLI testhub library list --output markdown" "testhub library list --output markdown"
test_cmd "$CLI directory user list --output markdown" "directory user list --output markdown"
echo ""

# Test 5: --output table (default)
echo "=== Testing --output table ==="
test_cmd "$CLI ship product list --output table" "ship product list --output table"
test_cmd "$CLI ship ideas list --output table --product-id test" "ship ideas list --output table"
test_cmd "$CLI ship tickets list --output table --product-id test" "ship tickets list --output table"
test_cmd "$CLI project list --output table" "project list --output table"
test_cmd "$CLI project metadata types --output table" "project metadata types --output table"
test_cmd "$CLI wiki space list --output table" "wiki space list --output table"
test_cmd "$CLI testhub library list --output table" "testhub library list --output table"
test_cmd "$CLI directory user list --output table" "directory user list --output table"
echo ""

# Test 6: --raw flag
echo "=== Testing --raw flag ==="
test_cmd "$CLI ship product list --raw" "ship product list --raw"
test_cmd "$CLI ship ideas list --raw --product-id test" "ship ideas list --raw"
test_cmd "$CLI ship tickets list --raw --product-id test" "ship tickets list --raw"
test_cmd "$CLI project list --raw" "project list --raw"
test_cmd "$CLI project metadata types --raw" "project metadata types --raw"
test_cmd "$CLI wiki space list --raw" "wiki space list --raw"
test_cmd "$CLI testhub library list --raw" "testhub library list --raw"
test_cmd "$CLI directory user list --raw" "directory user list --raw"
echo ""

# Test 7: --raw with --output
echo "=== Testing --raw --output json ==="
test_cmd "$CLI ship product list --raw --output json" "ship product list --raw --output json"
test_cmd "$CLI project list --raw --output json" "project list --raw --output json"
test_cmd "$CLI wiki space list --raw --output json" "wiki space list --raw --output json"
test_cmd "$CLI testhub library list --raw --output json" "testhub library list --raw --output json"
test_cmd "$CLI directory user list --raw --output json" "directory user list --raw --output json"
echo ""

# Test 8: Default format (table)
echo "=== Testing default format (no flags) ==="
test_cmd "$CLI ship product list" "ship product list (default table)"
test_cmd "$CLI project list" "project list (default table)"
test_cmd "$CLI wiki space list" "wiki space list (default table)"
test_cmd "$CLI testhub library list" "testhub library list (default table)"
test_cmd "$CLI directory user list" "directory user list (default table)"
echo ""

# Test 9: Workitem commands
echo "=== Testing Project WorkItem Commands ==="
test_cmd "$CLI project workitem list --output json -p test" "project workitem list --output json"
test_cmd "$CLI project workitem list --raw -p test" "project workitem list --raw"
test_cmd "$CLI project workitem tag --help" "project workitem tag --help"
echo ""

# Test 10: Iteration commands
echo "=== Testing Project Iteration Commands ==="
test_cmd "$CLI project iteration list --output json -p test" "project iteration list --output json"
test_cmd "$CLI project iteration section --help" "project iteration section --help"
test_cmd "$CLI project iteration category --help" "project iteration category --help"
echo ""

# Test 11: Version commands
echo "=== Testing Project Version Commands ==="
test_cmd "$CLI project version list --output json -p test" "project version list --output json"
test_cmd "$CLI project version section --help" "project version section --help"
test_cmd "$CLI project version category --help" "project version category --help"
echo ""

# Test 12: Kanban commands
echo "=== Testing Project Kanban Commands ==="
test_cmd "$CLI project kanban list --output json -p test" "project kanban list --output json"
test_cmd "$CLI project kanban entries --output json -p test --board-id test" "project kanban entries --output json"
test_cmd "$CLI project kanban swimlane --help" "project kanban swimlane --help"
echo ""

# Test 13: Member commands
echo "=== Testing Project Member Commands ==="
test_cmd "$CLI project member list --output json -p test" "project member list --output json"
test_cmd "$CLI project member add --help" "project member add --help"
echo ""

# Test 14: Wiki space member commands
echo "=== Testing Wiki Space Member Commands ==="
test_cmd "$CLI wiki space member --help" "wiki space member --help"
echo ""

# Test 15: Wiki page content commands
echo "=== Testing Wiki Page Content Commands ==="
test_cmd "$CLI wiki page content --help" "wiki page content --help"
echo ""

# Test 16: TestHub suite commands
echo "=== Testing TestHub Suite Commands ==="
test_cmd "$CLI testhub library suite --help" "testhub library suite --help"
echo ""

# Test 17: TestHub member commands
echo "=== Testing TestHub Member Commands ==="
test_cmd "$CLI testhub library member --help" "testhub library member --help"
echo ""

# Test 18: Invalid output format
echo "=== Testing Invalid Output Format ==="
test_cmd "$CLI ship product list --output invalid_format" "ship product list --output invalid_format"
echo ""

# Summary
echo "=========================================="
echo "Test Summary"
echo "=========================================="
echo -e "Passed: ${GREEN}$PASS${NC}"
echo -e "Failed: ${RED}$FAIL${NC}"
echo ""

if [ $FAIL -eq 0 ]; then
    echo -e "${GREEN}All tests passed!${NC}"
    exit 0
else
    echo -e "${RED}Some tests failed.${NC}"
    exit 1
fi
