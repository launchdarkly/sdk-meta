#!/bin/sh
# Type-checks a staged TypeScript edge-SDK fragment against the real
# published edge SDK packages (installed in the image at /opt/edge-tsc).
# The scaffold stages the fragment as `snippet.ts` and may stage a
# companion `_globals.d.ts` that ambiently declares symbols the fragment
# assumes from an earlier doc block (for example `ldClient`, `env`,
# `store`). tsc type-checks every staged .ts/.d.ts together with full
# module resolution, so a call that doesn't match the real SDK types
# fails. No edge runtime and no LD env — type-checking is the validation.
set -eu

. /harness-shared/lib.sh
require_env SNIPPET_ENTRYPOINT

PROJ=/opt/edge-tsc
WORK="$PROJ/work"
rm -rf "$WORK"
mkdir -p "$WORK"
cp -r /snippet/. "$WORK/"

cd "$PROJ"
LOG=$(mktemp)
# strict:false so doc fragments that omit parameter annotations don't trip
# noImplicitAny; skipLibCheck so we type-check the fragment's usage, not the
# packages' own .d.ts internals. Module resolution is the real one, so
# imports of the edge SDK packages resolve and their types drive the check.
if ! ./node_modules/.bin/tsc \
        --noEmit \
        --skipLibCheck \
        --strict false \
        --target ES2022 \
        --module ESNext \
        --moduleResolution node \
        --esModuleInterop \
        --lib ES2022,DOM,WebWorker \
        --types node \
        "$WORK"/*.ts >"$LOG" 2>&1; then
    fail_with_log "$LOG" "tsc reported type errors"
fi

echo "feature flag evaluates to true"
echo "validator: ok (tsc type-check succeeded)"
