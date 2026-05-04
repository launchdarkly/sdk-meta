#!/bin/sh
# Builds the staged Vue snippet against the pre-baked Vite project,
# starts a preview server, and runs the Playwright DOM check against it.
set -eu

. /harness-shared/lib.sh
require_env LAUNCHDARKLY_CLIENT_SIDE_ID LAUNCHDARKLY_FLAG_KEY SNIPPET_ENTRYPOINT

# Stage every src/* file the snippet ships into the pre-baked Vue
# project. The snippet's `file:` paths are project-relative
# (src/App.vue, src/main.js); init scaffolds also stage a companion
# App.vue. We copy whatever the snippet provided rather than naming
# specific files so an init scaffold can ship its own App.vue companion
# alongside main.js.
for f in /snippet/src/*; do
    [ -f "$f" ] || continue
    cp "$f" "/opt/hello-vue/src/$(basename "$f")"
done
if [ -n "${SNIPPET_ENTRYPOINT:-}" ] && [ -f "/snippet/$SNIPPET_ENTRYPOINT" ]; then
    cp "/snippet/$SNIPPET_ENTRYPOINT" "/opt/hello-vue/$SNIPPET_ENTRYPOINT"
fi

cd /opt/hello-vue
npm run build >/tmp/build.log 2>&1 \
    || { cat /tmp/build.log >&2; exit 1; }

# Start vite preview in the background and let Playwright probe it.
PREVIEW_LOG=$(mktemp)
npm run preview >"$PREVIEW_LOG" 2>&1 &
PREVIEW_PID=$!

# Wait briefly for the server to come up. vite preview prints a "Local:"
# line within ~1s on this image.
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

VUE_PREVIEW_URL=http://localhost:4173 exec node /harness/check.js
