#!/bin/sh
# Runs the staged Node.js snippet against a real LaunchDarkly environment.
# Inputs (env): LAUNCHDARKLY_SDK_KEY (or _CLIENT_SIDE_ID for client SDKs),
#               LAUNCHDARKLY_FLAG_KEY, SNIPPET_ENTRYPOINT.
#
# The snippet's `validation.requirements` line(s) are passed in via a
# stage-time requirements.txt file (the dispatcher writes it). For Node,
# each line is treated as an `npm install <line>` argument.
set -eu

. /harness-shared/lib.sh
require_env LAUNCHDARKLY_FLAG_KEY SNIPPET_ENTRYPOINT

# Stage to a writable workdir; npm install writes node_modules.
WORK=$(mktemp -d)
trap 'rm -rf "$WORK"' EXIT
cp -r /snippet/. "$WORK/"
cd "$WORK"

# Initialize package.json if absent (snippet's `mkdir + npm init` step
# isn't reproduced verbatim — npm init is interactive).
if [ ! -f package.json ]; then
    npm init -y >/dev/null
fi

if [ -f requirements.txt ]; then
    # Each non-empty line is an npm install target.
    while IFS= read -r line; do
        [ -z "$line" ] && continue
        npm install --silent --no-audit --no-fund --no-progress "$line"
    done < requirements.txt
fi

LOG=$(mktemp)

CI=1 timeout --signal=TERM 60s node "$SNIPPET_ENTRYPOINT" >"$LOG" 2>&1 &
PID=$!

deadline=$(( $(date +%s) + 50 ))
if await_success_line "$LOG" "$PID" "$deadline"; then
    exit 0
fi

kill -TERM "$PID" 2>/dev/null || true
wait "$PID" 2>/dev/null || true
fail_with_log "$LOG" "did not see expected line: feature flag evaluates to true"
