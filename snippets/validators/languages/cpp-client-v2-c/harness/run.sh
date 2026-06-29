#!/bin/sh
# Batch parse-and-type-check for legacy v2.x C client SDK fragments. Each
# snippet is gcc-compiled (parse + type-check only) against the stub
# <launchdarkly/api.h> baked into the image — no SDK build, no flag eval.
# Batch mode loops every staged snippet in one container invocation rather
# than paying a container start per snippet.
set -eu

. /harness-shared/lib.sh
require_env SNIPPET_BATCH

validate_one() {
    relpath=$1
    LOG=$(mktemp)
    if gcc -x c -std=c11 -Wall -c "/snippet/$relpath" -o /dev/null >"$LOG" 2>&1; then
        rm -f "$LOG"
        return 0
    fi
    cat "$LOG" >&2
    rm -f "$LOG"
    return 1
}

run_batch validate_one
