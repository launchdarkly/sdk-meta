#!/bin/sh
# Batch parse-and-type-check for C++-flavored fragments of the legacy v2.x
# C server SDK. Each snippet is g++ -fsyntax-only checked against the stub
# <launchdarkly/api.h> baked into the image — no SDK build, no flag eval.
# Batch mode loops every staged snippet in one container invocation rather
# than paying a container start per snippet.
set -eu

. /harness-shared/lib.sh
require_env SNIPPET_BATCH

validate_one() {
    relpath=$1
    LOG=$(mktemp)
    if g++ -x c++ -std=c++17 -fsyntax-only -Wall "/snippet/$relpath" >"$LOG" 2>&1; then
        rm -f "$LOG"
        return 0
    fi
    cat "$LOG" >&2
    rm -f "$LOG"
    return 1
}

run_batch validate_one
