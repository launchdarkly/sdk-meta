#!/bin/sh
# Runs the staged Go snippet against a real LaunchDarkly environment.
# Inputs (env): LAUNCHDARKLY_SDK_KEY, LAUNCHDARKLY_FLAG_KEY, SNIPPET_ENTRYPOINT.
#
# The gonfalon `go mod init` step is reproduced here: we initialize a
# throwaway module in the working dir and let `go mod tidy` resolve the
# imports the snippet brings in. This mirrors what a developer following
# the Get Started instructions would do.
set -eu

. /harness-shared/lib.sh
require_env LAUNCHDARKLY_SDK_KEY LAUNCHDARKLY_FLAG_KEY SNIPPET_ENTRYPOINT

# Stage the snippet contents into a writable workdir (the bind-mount is
# read-only). go mod init writes go.mod / go.sum, so we can't operate on
# /snippet directly.
WORK=$(mktemp -d)
trap 'rm -rf "$WORK"' EXIT
cp -r /snippet/. "$WORK/"
cd "$WORK"

go mod init example/hello-go >/dev/null 2>&1
go mod tidy >/dev/null 2>&1

LOG=$(mktemp)

# CI=1 makes the snippet exit after the first evaluation instead of blocking
# on the listener loop.
CI=1 go run "$SNIPPET_ENTRYPOINT" >"$LOG" 2>&1 &
PID=$!

deadline=$(( $(date +%s) + 60 ))
if await_success_line "$LOG" "$PID" "$deadline"; then
    exit 0
fi

kill -TERM "$PID" 2>/dev/null || true
wait "$PID" 2>/dev/null || true
fail_with_log "$LOG" "did not see expected line: feature flag evaluates to true"
