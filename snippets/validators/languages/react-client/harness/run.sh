#!/bin/sh
# Builds the staged React snippet against the pre-baked Vite project,
# starts a preview server, and runs the Playwright DOM check against it.
#
# Two snippet variants flow through this same harness: the legacy CRA
# pattern uses src/index.tsx, while the createApp/Vite pattern uses
# src/main.tsx. We rewrite index.html to point at whichever entrypoint
# the snippet declared.
set -eu

. /harness-shared/lib.sh
require_env LAUNCHDARKLY_CLIENT_SIDE_ID LAUNCHDARKLY_FLAG_KEY SNIPPET_ENTRYPOINT

# Stage the snippet body (App.tsx) plus its companion (index.tsx or
# main.tsx) into the pre-baked React project.
cp "/snippet/$SNIPPET_ENTRYPOINT" "/opt/hello-react/$SNIPPET_ENTRYPOINT"
for f in /snippet/src/*.tsx; do
    [ -f "$f" ] || continue
    bn=$(basename "$f")
    cp "$f" "/opt/hello-react/src/$bn"
done

# Point index.html at whichever entrypoint the snippet uses (index.tsx
# for legacy, main.tsx for createApp).
ENTRY_BASENAME=$(basename "$SNIPPET_ENTRYPOINT")
ENTRY_FILE="src/$(basename "$SNIPPET_ENTRYPOINT" .tsx).tsx"
# Companion may also be the entrypoint script. Pick whichever isn't App.tsx.
SCRIPT_SRC=""
for f in /opt/hello-react/src/*.tsx; do
    bn=$(basename "$f")
    case "$bn" in
        App.tsx) ;;
        *) SCRIPT_SRC="/src/$bn"; break ;;
    esac
done
if [ -z "$SCRIPT_SRC" ]; then
    echo "harness: could not find non-App entrypoint in /opt/hello-react/src" >&2
    exit 1
fi

cd /opt/hello-react
sed -i "s|/src/main.tsx|$SCRIPT_SRC|" index.html

npm run build >/tmp/build.log 2>&1 \
    || { cat /tmp/build.log >&2; exit 1; }

PREVIEW_LOG=$(mktemp)
npm run preview >"$PREVIEW_LOG" 2>&1 &
PREVIEW_PID=$!

for _ in $(seq 1 20); do
    if grep -q 'Local:' "$PREVIEW_LOG" 2>/dev/null; then
        break
    fi
    sleep 0.2
done

cleanup() {
    kill -TERM "$PREVIEW_PID" 2>/dev/null || true
    wait "$PREVIEW_PID" 2>/dev/null || true
}
trap cleanup EXIT

REACT_PREVIEW_URL=http://localhost:4173 exec node /harness/check.js
