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

# If the staged main.dart uses the flutter-syntax-only scaffold's
# body marker pair, lift any `import 'package:...';` directives from
# inside the body up to module scope. Dart forbids imports inside a
# function body, so doc fragments that show install-time imports
# would otherwise fail to compile.
if grep -qF -- '//IMPORT_LIFT_TARGET' /opt/hello_flutter/lib/main.dart; then
    awk '
    /^\/\/IMPORT_LIFT_TARGET$/ {
        target_seen = 1;
        pre[++npre] = $0;
        target_index = npre;
        next;
    }
    /^\/\/BODY_BEGIN$/ {
        in_body = 1;
        next;
    }
    /^\/\/BODY_END$/ {
        in_body = 0;
        body_done = 1;
        next;
    }
    {
        if (in_body) {
            if ($0 ~ /^[ \t]*import [\x27"]/) {
                sub(/^[ \t]+/, "", $0);
                lift[++nlift] = $0;
            } else {
                rest[++nrest] = $0;
            }
        } else if (body_done) {
            post[++npost] = $0;
        } else if (target_seen) {
            mid[++nmid] = $0;
        } else {
            pre[++npre] = $0;
        }
    }
    END {
        for (i = 1; i <= npre; i++) {
            print pre[i];
            if (target_seen && i == target_index) {
                for (j = 1; j <= nlift; j++) print lift[j];
            }
        }
        for (i = 1; i <= nmid; i++) print mid[i];
        for (i = 1; i <= nrest; i++) print rest[i];
        for (i = 1; i <= npost; i++) print post[i];
    }
    ' /opt/hello_flutter/lib/main.dart > /tmp/main.dart.lifted
    mv /tmp/main.dart.lifted /opt/hello_flutter/lib/main.dart
fi

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
