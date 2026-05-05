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

# Don't use await_success_line here: jest's failure output prints the
# expected regex pattern verbatim ("Expected pattern: /feature flag
# evaluates to true/i"), which would falsely match before jest had a
# chance to exit non-zero. Wait for jest to exit, then check both its
# exit code and that the success line appeared. The test prints the
# rendered text via console.log only after `expect().toMatch` passes.
LOG=$(mktemp)
if ! timeout --signal=TERM 180s npm test --silent >"$LOG" 2>&1; then
    fail_with_log "$LOG" "jest exited non-zero"
fi

if grep -E "feature flag evaluates to [Tt]rue" "$LOG" >/dev/null 2>&1; then
    grep -E "feature flag evaluates to [Tt]rue" "$LOG" | head -1
    echo "validator: ok"
    exit 0
fi

fail_with_log "$LOG" "did not see expected line: feature flag evaluates to true"
