#!/bin/sh
# Batch validator for C++ client SDK snippets. Each snippet is a single
# main.cpp; gonfalon's Get Started flow has the user clone cpp-sdks
# alongside their project and add it via CMake. The Dockerfile pre-cloned
# cpp-sdks and pre-baked a CONFIGURED project (build dir already generated
# against add_subdirectory(cpp-sdks)), so per-snippet we only swap main.cpp
# and run an incremental `cmake --build --target hello` — no per-snippet
# CMake configure, which is what made the old one-per-snippet path slow.
#
# The Go runner stages every matching snippet under /snippet and writes
# /snippet/manifest.tsv; run_batch loops over it in the single warm project.
set -eu

. /harness-shared/lib.sh
require_env LAUNCHDARKLY_MOBILE_KEY LAUNCHDARKLY_FLAG_KEY SNIPPET_BATCH

cd /opt/hello-cpp

validate_one() {
    relpath=$1
    cp "/snippet/$relpath" main.cpp

    BUILD_LOG=$(mktemp)
    if ! cmake --build build --target hello >"$BUILD_LOG" 2>&1; then
        cat "$BUILD_LOG" >&2
        rm -f "$BUILD_LOG"
        return 1
    fi
    rm -f "$BUILD_LOG"

    LOG=$(mktemp)
    timeout --signal=TERM 60s ./build/hello >"$LOG" 2>&1 &
    PID=$!
    deadline=$(( $(date +%s) + 55 ))
    if await_success_line "$LOG" "$PID" "$deadline"; then
        rm -f "$LOG"
        return 0
    fi
    kill -TERM "$PID" 2>/dev/null || true
    wait "$PID" 2>/dev/null || true
    dump_redacted "$LOG" >&2
    rm -f "$LOG"
    return 1
}

run_batch validate_one
