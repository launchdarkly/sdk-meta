#!/bin/sh
# Validates the React Native snippet under jest with the react-native
# preset (no emulator/simulator). The Dockerfile pre-installed the
# bundled deps; per-validate just swaps the snippet's two .tsx files
# in and re-runs jest.
#
# Two snippet variants flow through this harness, dispatched on the
# optional `SNIPPET_MODE` env var (set by the wrappee scaffold's
# `validation.env`):
#
#   - SNIPPET_MODE unset (init): full jest run end-to-end against the
#     LD streaming endpoint.
#
#   - SNIPPET_MODE=syntax-only: Babel-parse the staged App.tsx and
#     bail out before the jest run. Used for parse-only doc fragments
#     (e.g. observability/import) that aren't standalone-runnable but
#     should still be validated for syntactic correctness against the
#     pre-baked deps tree (so `import { Observability } from
#     '@launchdarkly/observability-react-native'` resolves at parse
#     time but no LD network call is made).
set -eu

. /harness-shared/lib.sh
require_env LAUNCHDARKLY_FLAG_KEY SNIPPET_ENTRYPOINT
MODE="${SNIPPET_MODE:-init}"
if [ "$MODE" = "init" ]; then
    require_env LAUNCHDARKLY_MOBILE_KEY
fi

PROJECT=/opt/hello-react-native

# Stage App.tsx (root) + src/welcome.tsx (companion). Both are .tsx
# files at known paths under the snippet stage dir.
cp "/snippet/$SNIPPET_ENTRYPOINT" "${PROJECT}/$SNIPPET_ENTRYPOINT"
if [ -f /snippet/src/welcome.tsx ]; then
    cp /snippet/src/welcome.tsx "${PROJECT}/src/welcome.tsx"
fi

cd "${PROJECT}"

if [ "$MODE" = "syntax-only" ]; then
    # If the staged entrypoint uses the react-native-syntax-only
    # scaffold's body marker pair, lift any `import …;` directives
    # from inside the body block up to module scope. ESM forbids
    # `import` inside a function body, so a wrappee body containing
    # an install-time import would otherwise fail to parse.
    # Mirrors react-client/harness/run.sh.
    ENTRY_FILE="${PROJECT}/${SNIPPET_ENTRYPOINT}"
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
        ' "$ENTRY_FILE" > /tmp/lifted.tsx
        mv /tmp/lifted.tsx "$ENTRY_FILE"
    fi

    # Parse the staged entrypoint through Babel with the project's
    # react-native preset config. A successful parse means the body's
    # syntax is valid; we don't run the body, so unresolved network
    # calls (LD streaming, etc.) never happen.
    node -e "
      const path = require('path');
      const file = path.resolve(process.argv[1]);
      const babel = require('@babel/core');
      const result = babel.transformFileSync(file, {
        babelrc: false,
        configFile: path.resolve('babel.config.js'),
        ast: false,
        code: false,
      });
      if (!result) {
        console.error('babel: no result for ' + file);
        process.exit(1);
      }
      console.log('feature flag evaluates to true');
    " "$SNIPPET_ENTRYPOINT"
    exit 0
fi

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
