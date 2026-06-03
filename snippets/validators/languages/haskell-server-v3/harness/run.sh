#!/bin/sh
# Runs the staged Haskell snippet against a real LaunchDarkly environment.
# The Dockerfile pre-bootstrapped a cabal project at /opt/hello-haskell
# with launchdarkly-server-sdk + text already compiled. Per-validate just
# swaps in the user's Main.hs and does an incremental rebuild.
set -eu

. /harness-shared/lib.sh
require_env LAUNCHDARKLY_SDK_KEY LAUNCHDARKLY_FLAG_KEY SNIPPET_ENTRYPOINT

cp "/snippet/$SNIPPET_ENTRYPOINT" /opt/hello-haskell/app/Main.hs

# If the staged Main.hs uses the haskell-syntax-only scaffold's body
# marker pair, lift top-level constructs out of the body region. The
# scaffold places the body inside a `_wrappee = do` block, so any
# body line at column 0 that is `import`/`data`/`type`/`newtype`/
# `class`/`instance` or matches a `name :: ...` / `name = ...`
# top-level binding shape gets relocated to the TOP_LIFT_TARGET marker
# at module scope. Body lines that aren't top-level shape stay in
# place but get indented two spaces so they sit inside the do-block.
# The scaffold's own bindings live above BODY_BEGIN and aren't touched.
if grep -qF -- '--TOP_LIFT_TARGET--' /opt/hello-haskell/app/Main.hs; then
    awk '
    BEGIN { in_body = 0; target_seen = 0; body_done = 0; }
    /^--TOP_LIFT_TARGET--$/ {
        target_seen = 1;
        pre[++npre] = $0;
        target_index = npre;
        next;
    }
    /^--BODY_BEGIN--$/ {
        in_body = 1;
        next;
    }
    /^--BODY_END--$/ {
        in_body = 0;
        body_done = 1;
        next;
    }
    {
        if (in_body) {
            if ($0 ~ /^(import |data |type |newtype |class |instance )/ ||
                $0 ~ /^[A-Za-z_][A-Za-z0-9_'\'']*[ \t]+(::|=)/) {
                lift[++nlift] = $0;
            } else if ($0 ~ /^[^ \t]/ && length($0) > 0) {
                rest[++nrest] = "  " $0;
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
    ' /opt/hello-haskell/app/Main.hs > /tmp/Main.hs.lifted
    mv /tmp/Main.hs.lifted /opt/hello-haskell/app/Main.hs
fi

cd /opt/hello-haskell
cabal build >/tmp/build.log 2>&1 \
    || { cat /tmp/build.log >&2; exit 1; }

LOG=$(mktemp)

timeout --signal=TERM 60s cabal run hello-haskell-exe >"$LOG" 2>&1 &
PID=$!

deadline=$(( $(date +%s) + 55 ))
if await_success_line "$LOG" "$PID" "$deadline"; then
    exit 0
fi

kill -TERM "$PID" 2>/dev/null || true
wait "$PID" 2>/dev/null || true
fail_with_log "$LOG" "did not see expected line: feature flag evaluates to true"
