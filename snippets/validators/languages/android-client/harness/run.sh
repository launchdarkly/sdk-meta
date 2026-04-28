#!/bin/sh
# Validates the Android snippet under Robolectric (no emulator). The
# Dockerfile pre-bakes a hello-android scaffold + Robolectric test;
# per-validate we just swap the snippet's two Kotlin files in and run
# the JUnit test.
set -eu

. /harness-shared/lib.sh
require_env LAUNCHDARKLY_MOBILE_KEY LAUNCHDARKLY_FLAG_KEY SNIPPET_ENTRYPOINT

# The snippet declares main-application + main-activity as separate
# files. Both come through as part of the staging dir — copy whichever
# .kt files /snippet has into the scaffold's main source tree.
SCAFFOLD=/opt/hello-android
PKG_DIR="${SCAFFOLD}/app/src/main/java/com/launchdarkly/hello_android"

for f in /snippet/app/src/main/java/com/launchdarkly/hello_android/*.kt; do
    [ -f "$f" ] || continue
    cp "$f" "${PKG_DIR}/$(basename "$f")"
done

cd "${SCAFFOLD}"

LOG=$(mktemp)
timeout --signal=TERM 600s ./gradlew --no-daemon \
        testDebugUnitTest --tests='*HelloAppTest*' --console=plain \
        >"$LOG" 2>&1 &
PID=$!

deadline=$(( $(date +%s) + 590 ))
if await_success_line "$LOG" "$PID" "$deadline"; then
    exit 0
fi

kill -TERM "$PID" 2>/dev/null || true
wait "$PID" 2>/dev/null || true
fail_with_log "$LOG" "did not see expected line: feature flag evaluates to true"
