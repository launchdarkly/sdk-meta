#!/bin/sh
# Validator entrypoint. Runs the staged snippet against a real LaunchDarkly
# environment.
#
# Required env (passed in by `snippets validate`):
#   LAUNCHDARKLY_SDK_KEY    server-side SDK key for the test environment
#   LAUNCHDARKLY_FLAG_KEY   the flag the snippet is templated to evaluate
#   SNIPPET_ENTRYPOINT      file under /snippet to run (e.g. main.py)
#
# Success criterion: the snippet prints
#   `*** The <LAUNCHDARKLY_FLAG_KEY> feature flag evaluates to ...`
# within the timeout. That line only appears on a successful SDK init + flag
# evaluation, so matching it is a real signal.
set -eu

: "${LAUNCHDARKLY_SDK_KEY:?LAUNCHDARKLY_SDK_KEY not set}"
: "${LAUNCHDARKLY_FLAG_KEY:?LAUNCHDARKLY_FLAG_KEY not set}"
: "${SNIPPET_ENTRYPOINT:?SNIPPET_ENTRYPOINT not set}"

if [ -f /snippet/requirements.txt ]; then
    pip install --quiet --no-input -r /snippet/requirements.txt
fi

LOG=$(mktemp)
trap 'rm -f "$LOG"' EXIT

cd /snippet

# Hello-world programs block on Event().wait(); run with a timeout. We watch
# the log for the success line and SIGTERM the process as soon as it appears.
PYTHONUNBUFFERED=1 timeout --signal=TERM 30s python "$SNIPPET_ENTRYPOINT" >"$LOG" 2>&1 &
PID=$!

deadline=$(( $(date +%s) + 25 ))
prefix="*** The ${LAUNCHDARKLY_FLAG_KEY} feature flag evaluates to "

while [ "$(date +%s)" -lt "$deadline" ]; do
    if grep -F -- "$prefix" "$LOG" >/dev/null 2>&1; then
        kill -TERM "$PID" 2>/dev/null || true
        wait "$PID" 2>/dev/null || true
        # Echo the matched line so the caller sees the actual value.
        grep -F -- "$prefix" "$LOG" | head -1
        echo "validator: ok"
        exit 0
    fi
    if ! kill -0 "$PID" 2>/dev/null; then
        break
    fi
    sleep 0.2
done

kill -TERM "$PID" 2>/dev/null || true
wait "$PID" 2>/dev/null || true
echo "validator: did not see expected line: ${prefix}<value>" >&2
echo "--- snippet output (LAUNCHDARKLY_SDK_KEY redacted) ---" >&2
# Defense-in-depth: today's snippets never print LAUNCHDARKLY_SDK_KEY, but a
# future snippet could (intentionally or by accident) and this log gets piped
# into CI output. Redact any literal occurrence of the key before dumping.
sed -e "s|${LAUNCHDARKLY_SDK_KEY}|<REDACTED_LAUNCHDARKLY_SDK_KEY>|g" "$LOG" >&2
exit 1
