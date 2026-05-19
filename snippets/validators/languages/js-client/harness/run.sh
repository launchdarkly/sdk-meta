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

# If the staged entrypoint uses the js-syntax-only scaffold's body
# marker pair, lift any `import …;` directives from inside the body
# block up to module scope. ESM forbids `import` inside a function
# body, so a wrappee body containing an install-time import would
# otherwise fail to parse. Mirrors react-client/harness/run.sh.
ENTRY_FILE="/opt/hello-js/${SNIPPET_ENTRYPOINT:-src/app.ts}"
if [ -f "$ENTRY_FILE" ] && grep -qF -- '//IMPORT_LIFT_TARGET' "$ENTRY_FILE"; then
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
            if (in_multi_import) {
                sub(/^[ \t]+/, "", $0);
                multi_import_buf = multi_import_buf "\n" $0;
                if ($0 ~ /;[ \t]*$/) {
                    lift[++nlift] = multi_import_buf;
                    in_multi_import = 0;
                    multi_import_buf = "";
                }
                next;
            }
            if ($0 ~ /^[ \t]*import[ \t]/) {
                sub(/^[ \t]+/, "", $0);
                # Single-line import: ends in `;` or trailing module-name quote.
                if ($0 ~ /;[ \t]*$/ || $0 ~ /[\x27"][ \t]*$/) {
                    lift[++nlift] = $0;
                } else {
                    multi_import_buf = $0;
                    in_multi_import = 1;
                }
                next;
            }
            sub(/^[ \t]*export[ \t]+default[ \t]+/, "", $0);
            sub(/^[ \t]*export[ \t]+/, "", $0);
            rest[++nrest] = $0;
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
    ' "$ENTRY_FILE" > /tmp/lifted.ts
    mv /tmp/lifted.ts "$ENTRY_FILE"
fi

cd /opt/hello-js

npm run build >/tmp/build.log 2>&1 \
    || { cat /tmp/build.log >&2; exit 1; }

exec node /harness/check.js
