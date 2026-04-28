#!/bin/sh
# Runs the staged Rust snippet against a real LaunchDarkly environment.
# The harness reproduces gonfalon's `cargo new` + `cargo add` flow:
# bootstrap a Cargo project, drop the snippet's src/main.rs over the
# default template, add the SDK + tokio dependencies, and run it.
set -eu

. /harness-shared/lib.sh
require_env LAUNCHDARKLY_SDK_KEY LAUNCHDARKLY_FLAG_KEY SNIPPET_ENTRYPOINT

WORK=$(mktemp -d)
trap 'rm -rf "$WORK"' EXIT
cd "$WORK"

cargo new --quiet --bin hello-rust
cd hello-rust

# Replace the default src/main.rs with the snippet body. The snippet's
# `file:` is `src/main.rs`, so /snippet/src/main.rs holds it.
cp "/snippet/$SNIPPET_ENTRYPOINT" "$SNIPPET_ENTRYPOINT"

cargo add --quiet launchdarkly-server-sdk
cargo add --quiet tokio@1 -F rt,macros

LOG=$(mktemp)

timeout --signal=TERM 300s cargo run --quiet >"$LOG" 2>&1 &
PID=$!

deadline=$(( $(date +%s) + 290 ))
if await_success_line "$LOG" "$PID" "$deadline"; then
    exit 0
fi

kill -TERM "$PID" 2>/dev/null || true
wait "$PID" 2>/dev/null || true
fail_with_log "$LOG" "did not see expected line: feature flag evaluates to true"
