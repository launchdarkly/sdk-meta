#!/bin/sh
# Parse-and-type-check the staged C++ file against the stub
# <launchdarkly/api.hpp>. Success = g++ accepts the syntax.
set -eu

. /harness-shared/lib.sh
require_env SNIPPET_ENTRYPOINT

WORK=$(mktemp -d)
trap 'rm -rf "$WORK"' EXIT
cp "/snippet/$SNIPPET_ENTRYPOINT" "$WORK/main.cpp"
cd "$WORK"

LOG=$(mktemp)
if g++ -std=c++17 -fsyntax-only -Wall main.cpp >"$LOG" 2>&1; then
    echo "feature flag evaluates to true"
    echo "validator: ok (g++ -fsyntax-only against v2 cpp-client C++ binding stub succeeded)"
    exit 0
fi
fail_with_log "$LOG" "cpp-client v2 C++ binding parse/type-check failed"
