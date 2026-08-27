#!/bin/sh
# Runs the staged Ruby snippet against a real LaunchDarkly environment.
set -eu

. /harness-shared/lib.sh
require_env LAUNCHDARKLY_SDK_KEY LAUNCHDARKLY_FLAG_KEY SNIPPET_ENTRYPOINT
# SNIPPET_SUCCESS_RE (optional, via validation.env) overrides the awaited
# success line for wrappees that own their whole program and print their
# own line (e.g. sdk-info init snippets).
SUCCESS_RE="${SNIPPET_SUCCESS_RE:-}"

WORK=$(mktemp -d)
trap 'rm -rf "$WORK"' EXIT
cp -r /snippet/. "$WORK/"
cd "$WORK"

# Generate a Gemfile from validation.requirements (newline-separated gem names).
if [ -f requirements.txt ] && [ ! -f Gemfile ]; then
    {
        echo "source 'https://rubygems.org'"
        while IFS= read -r line; do
            [ -z "$line" ] && continue
            echo "gem '$line'"
        done < requirements.txt
    } > Gemfile
fi

if [ -f Gemfile ]; then
    bundle install --quiet
fi

LOG=$(mktemp)

# Ruby snippet blocks on Thread sleep; we time it out and SIGTERM after match.
# Force line-buffered stdout — Ruby block-buffers when stdout isn't a tty,
# which would hide the success line until the deadline fires. stdbuf only
# affects libc-level buffering and doesn't help with Ruby's own IO layer,
# so we wrap the snippet in a one-liner that sets $stdout.sync = true
# before loading it. This keeps the snippet itself unmodified.
timeout --signal=TERM 90s ruby -e '$stdout.sync = true; load ENV["SNIPPET_ENTRYPOINT"]' >"$LOG" 2>&1 &
PID=$!

deadline=$(( $(date +%s) + 80 ))
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
