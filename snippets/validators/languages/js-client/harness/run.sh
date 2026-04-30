#!/bin/sh
# Builds the staged JavaScript snippet against the pre-baked tsdown
# project, then runs the Playwright DOM check against the resulting
# bundle loaded via a file:// URL.
#
# The snippet ships src/app.ts as its entrypoint and index.html as its
# companion; both get copied into the pre-baked project before tsdown
# runs.
set -eu

. /harness-shared/lib.sh
require_env LAUNCHDARKLY_CLIENT_SIDE_ID LAUNCHDARKLY_FLAG_KEY

# Stage the snippet's TypeScript source plus the index.html companion
# into the pre-baked project.
for f in /snippet/src/*.ts; do
    [ -f "$f" ] || continue
    bn=$(basename "$f")
    cp "$f" "/opt/hello-js/src/$bn"
done

if [ -f /snippet/index.html ]; then
    cp /snippet/index.html /opt/hello-js/index.html
fi

cd /opt/hello-js

npm run build >/tmp/build.log 2>&1 \
    || { cat /tmp/build.log >&2; exit 1; }

exec node /harness/check.js
