#!/bin/sh
# Builds the staged Flutter snippet against the pre-baked web project and
# runs the Playwright DOM check against a static-served bundle.
#
# The snippet's main.dart pulls credentials via
# `CredentialSource.fromEnvironment()`, which reads the
# LAUNCHDARKLY_CLIENT_SIDE_ID / LAUNCHDARKLY_MOBILE_KEY values that were
# baked into the build via --dart-define. Mobile keys aren't valid for
# the web target, so we only forward CLIENT_SIDE_ID here.
set -eu

. /harness-shared/lib.sh
require_env LAUNCHDARKLY_CLIENT_SIDE_ID LAUNCHDARKLY_FLAG_KEY SNIPPET_ENTRYPOINT

cp "/snippet/$SNIPPET_ENTRYPOINT" /opt/hello_flutter/lib/main.dart

cd /opt/hello_flutter

flutter build web --release --no-pub \
    --dart-define LAUNCHDARKLY_CLIENT_SIDE_ID="${LAUNCHDARKLY_CLIENT_SIDE_ID}" \
    >/tmp/build.log 2>&1 \
    || { cat /tmp/build.log >&2; exit 1; }

# Static-serve the build output. python3 is in the playwright base image.
PREVIEW_LOG=$(mktemp)
( cd build/web && python3 -m http.server 4173 ) >"$PREVIEW_LOG" 2>&1 &
PREVIEW_PID=$!

# python's http.server writes "Serving HTTP on …" once it binds.
for _ in $(seq 1 20); do
    if grep -q 'Serving' "$PREVIEW_LOG" 2>/dev/null; then
        break
    fi
    sleep 0.2
done

cleanup() {
    kill -TERM "$PREVIEW_PID" 2>/dev/null || true
    wait "$PREVIEW_PID" 2>/dev/null || true
}
trap cleanup EXIT

FLUTTER_PREVIEW_URL=http://localhost:4173 exec node /harness/check.js
