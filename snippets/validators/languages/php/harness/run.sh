#!/bin/sh
# Runs the staged PHP snippet against a real LaunchDarkly environment.
set -eu

. /harness-shared/lib.sh
require_env LAUNCHDARKLY_SDK_KEY LAUNCHDARKLY_FLAG_KEY SNIPPET_ENTRYPOINT

WORK=$(mktemp -d)
trap 'rm -rf "$WORK"' EXIT
cp -r /snippet/. "$WORK/"
cd "$WORK"

# composer require <pkg>... per requirements line. The image bundles a
# system-wide `composer` binary, so we use it directly rather than
# bootstrapping composer.phar like the snippet does.
if [ -f requirements.txt ]; then
    pkgs=""
    while IFS= read -r line; do
        [ -z "$line" ] && continue
        pkgs="$pkgs $line"
    done < requirements.txt
    if [ -n "$pkgs" ]; then
        # shellcheck disable=SC2086
        composer require --quiet --no-interaction --no-progress $pkgs
    fi
fi

LOG=$(mktemp)

# PHP snippet loops forever with sleep(1); time out and SIGTERM after match.
timeout --signal=TERM 90s php -d output_buffering=Off "$SNIPPET_ENTRYPOINT" >"$LOG" 2>&1 &
PID=$!

deadline=$(( $(date +%s) + 80 ))
if await_success_line "$LOG" "$PID" "$deadline"; then
    exit 0
fi

kill -TERM "$PID" 2>/dev/null || true
wait "$PID" 2>/dev/null || true
fail_with_log "$LOG" "did not see expected line: feature flag evaluates to true"
