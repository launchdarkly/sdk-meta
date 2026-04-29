#!/bin/sh
# Validates the React Native snippet under jest with the react-native
# preset (no emulator/simulator). The Dockerfile pre-installed the
# bundled deps; per-validate just swaps the snippet's two .tsx files
# in and re-runs jest.
set -eu

. /harness-shared/lib.sh
require_env LAUNCHDARKLY_MOBILE_KEY LAUNCHDARKLY_FLAG_KEY SNIPPET_ENTRYPOINT

PROJECT=/opt/hello-react-native

# Stage App.tsx (root) + src/welcome.tsx (companion). Both are .tsx
# files at known paths under the snippet stage dir.
cp "/snippet/App.tsx" "${PROJECT}/App.tsx"
cp "/snippet/src/welcome.tsx" "${PROJECT}/src/welcome.tsx"

cd "${PROJECT}"

LOG=$(mktemp)
timeout --signal=TERM 180s npm test --silent >"$LOG" 2>&1 &
PID=$!

deadline=$(( $(date +%s) + 170 ))
if await_success_line "$LOG" "$PID" "$deadline"; then
    exit 0
fi

kill -TERM "$PID" 2>/dev/null || true
wait "$PID" 2>/dev/null || true
fail_with_log "$LOG" "did not see expected line: feature flag evaluates to true"
