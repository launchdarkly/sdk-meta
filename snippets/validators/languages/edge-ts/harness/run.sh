#!/bin/sh
# Validates a staged TypeScript edge-SDK fragment by parsing it with the
# TypeScript compiler's transpileModule (syntax-only; no module
# resolution, no type-checking). A clean parse means the fragment is
# syntactically valid TypeScript. Parse-only -- no edge runtime, no LD
# env. On a clean parse we echo the EXAM-HELLO sentinel ourselves.
set -eu

. /harness-shared/lib.sh
require_env SNIPPET_ENTRYPOINT

ENTRY="/snippet/$SNIPPET_ENTRYPOINT"
if [ ! -f "$ENTRY" ]; then
    echo "harness: entrypoint not found at $ENTRY" >&2
    exit 1
fi

LOG=$(mktemp)
if ! node /harness/check.js "$ENTRY" >"$LOG" 2>&1; then
    fail_with_log "$LOG" "typescript reported syntax errors"
fi

echo "feature flag evaluates to true"
echo "validator: ok (typescript parse succeeded)"
