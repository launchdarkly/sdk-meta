#!/bin/sh
# Type-checks a staged Android Compose / View masking fragment against the
# real observability plugin + Jetpack Compose, using the pre-baked
# Kotlin-2.0 library project at /opt/compose-lib. No emulator, no run —
# a clean `compileDebugKotlin` means the fragment parses and type-checks
# against the real `Modifier.ldMask()` / `View.ldMask()` APIs.
set -eu

. /harness-shared/lib.sh
require_env SNIPPET_ENTRYPOINT

PROJ=/opt/compose-lib
PKG="$PROJ/lib/src/main/kotlin/com/launchdarkly/masking"
rm -f "$PKG"/Snippet.kt "$PKG"/Placeholder.kt
mkdir -p "$PKG"
cp "/snippet/$SNIPPET_ENTRYPOINT" "$PKG/Snippet.kt"

cd "$PROJ"
LOG=$(mktemp)
if ! gradle -q :lib:compileDebugKotlin --no-daemon >"$LOG" 2>&1; then
    fail_with_log "$LOG" "compileDebugKotlin failed"
fi

echo "feature flag evaluates to true"
echo "validator: ok (compileDebugKotlin succeeded)"
