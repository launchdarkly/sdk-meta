#!/bin/sh
# Runs the staged Haskell snippet against a real LaunchDarkly environment.
# The Dockerfile pre-bootstrapped a cabal project at /opt/hello-haskell
# with launchdarkly-server-sdk + text already compiled. Per-validate just
# swaps in the user's Main.hs and does an incremental rebuild.
set -eu

. /harness-shared/lib.sh
require_env LAUNCHDARKLY_SDK_KEY LAUNCHDARKLY_FLAG_KEY SNIPPET_ENTRYPOINT

cp "/snippet/$SNIPPET_ENTRYPOINT" /opt/hello-haskell/app/Main.hs

cd /opt/hello-haskell
cabal build >/tmp/build.log 2>&1 \
    || { cat /tmp/build.log >&2; exit 1; }

LOG=$(mktemp)

timeout --signal=TERM 60s cabal run hello-haskell-exe >"$LOG" 2>&1 &
PID=$!

deadline=$(( $(date +%s) + 55 ))
if await_success_line "$LOG" "$PID" "$deadline"; then
    exit 0
fi

kill -TERM "$PID" 2>/dev/null || true
wait "$PID" 2>/dev/null || true
fail_with_log "$LOG" "did not see expected line: feature flag evaluates to true"
