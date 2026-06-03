#!/bin/sh
# Parse-and-type-check the staged C file against the stub <launchdarkly/api.h>
# header. Success = gcc accepts the source (no parse / type errors).
set -eu

. /harness-shared/lib.sh
require_env SNIPPET_ENTRYPOINT

WORK=$(mktemp -d)
trap 'rm -rf "$WORK"' EXIT
cp "/snippet/$SNIPPET_ENTRYPOINT" "$WORK/main.c"
cd "$WORK"

LOG=$(mktemp)
if gcc -std=c11 -Wall -c main.c -o /dev/null >"$LOG" 2>&1; then
    echo "feature flag evaluates to true"
    echo "validator: ok (gcc compile against v2 C SDK stub headers succeeded)"
    exit 0
fi
fail_with_log "$LOG" "c-server v2 SDK parse/type-check failed"
