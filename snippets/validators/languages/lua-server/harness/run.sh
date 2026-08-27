#!/bin/sh
# Runs the staged Lua snippet against a real LaunchDarkly environment.
# The Dockerfile pre-built the C++ Server SDK shared lib, installed the
# Lua wrapper rock against it, and pinned LUA_PATH/LUA_CPATH/LD_LIBRARY_PATH
# so `lua` finds the launchdarkly_server_sdk module without any per-validate
# setup.
set -eu

. /harness-shared/lib.sh
require_env LAUNCHDARKLY_SDK_KEY LAUNCHDARKLY_FLAG_KEY SNIPPET_ENTRYPOINT
# SNIPPET_SUCCESS_RE (optional, via validation.env) overrides the awaited
# success line for wrappees that own their whole program and print their
# own line (e.g. sdk-info init snippets).
SUCCESS_RE="${SNIPPET_SUCCESS_RE:-}"

LOG=$(mktemp)

timeout --signal=TERM 60s lua5.3 "/snippet/$SNIPPET_ENTRYPOINT" >"$LOG" 2>&1 &
PID=$!

deadline=$(( $(date +%s) + 55 ))
if await_success_line "$LOG" "$PID" "$deadline" "$SUCCESS_RE"; then
    # CI's verify-hello-app asserts the cell output contains the
    # EXAM-HELLO line; when awaiting a custom success line, emit the
    # canonical sentinel too.
    if [ -n "$SUCCESS_RE" ]; then
        echo "feature flag evaluates to true"
    fi
    exit 0
fi

kill -TERM "$PID" 2>/dev/null || true
wait "$PID" 2>/dev/null || true
fail_with_log "$LOG" "did not see expected line: ${SUCCESS_RE:-feature flag evaluates to true}"
