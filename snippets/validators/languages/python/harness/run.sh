#!/bin/sh
# Runs the staged Python snippet against a real LaunchDarkly environment.
# Inputs (env): LAUNCHDARKLY_SDK_KEY, LAUNCHDARKLY_FLAG_KEY, SNIPPET_ENTRYPOINT.
set -eu

. /harness-shared/lib.sh
require_env LAUNCHDARKLY_SDK_KEY LAUNCHDARKLY_FLAG_KEY SNIPPET_ENTRYPOINT

if [ -f /snippet/requirements.txt ]; then
    pip install --quiet --no-input -r /snippet/requirements.txt
fi

LOG=$(mktemp)
trap 'rm -f "$LOG"' EXIT
cd /snippet

# Hello-world programs block on Event().wait(); run with a timeout. The
# shared lib watches the log for the success line and SIGTERMs the process
# as soon as it appears.
PYTHONUNBUFFERED=1 timeout --signal=TERM 30s python "$SNIPPET_ENTRYPOINT" >"$LOG" 2>&1 &
PID=$!

deadline=$(( $(date +%s) + 25 ))
if await_success_line "$LOG" "$PID" "$deadline"; then
    exit 0
fi

kill -TERM "$PID" 2>/dev/null || true
wait "$PID" 2>/dev/null || true
fail_with_log "$LOG" "did not see expected line: feature flag evaluates to true"
