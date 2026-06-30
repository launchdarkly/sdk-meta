#!/bin/sh
# Batch validator for Haskell server SDK snippets. The Dockerfile
# pre-bootstrapped a cabal project at /opt/hello-haskell with
# launchdarkly-server-sdk + deps already compiled; per-snippet just swaps
# in the user's Main.hs and does an incremental rebuild.
#
# The Go runner stages every matching snippet under /snippet and writes
# /snippet/manifest.tsv; run_batch loops over them in the single warm
# project (one container instead of one per snippet). Dispatch on
# SNIPPET_CHECK (snippets of one kind share a batch group):
#   - parse: a clean `cabal build` is the success condition (the module
#     defines its own main that would reach a real database if run).
#   - runtime: build then `cabal run` and await the success line.
set -eu

. /harness-shared/lib.sh
require_env LAUNCHDARKLY_SDK_KEY LAUNCHDARKLY_FLAG_KEY SNIPPET_BATCH
CHECK="${SNIPPET_CHECK:-runtime}"

cd /opt/hello-haskell

# lift_main <Main.hs>: if the staged module uses the haskell-syntax-only
# scaffold's body marker pair, lift top-level constructs (import/data/type/
# class/instance and `name ::`/`name =` bindings) out of the `_wrappee = do`
# block up to the TOP_LIFT_TARGET marker at module scope, indenting the
# remaining body lines to sit inside the do-block.
lift_main() {
    grep -qF -- '--TOP_LIFT_TARGET--' "$1" || return 0
    awk '
    BEGIN { in_body = 0; target_seen = 0; body_done = 0; }
    /^--TOP_LIFT_TARGET--$/ {
        target_seen = 1;
        pre[++npre] = $0;
        target_index = npre;
        next;
    }
    /^--BODY_BEGIN--$/ { in_body = 1; next; }
    /^--BODY_END--$/ { in_body = 0; body_done = 1; next; }
    {
        if (in_body) {
            if ($0 ~ /^(import |data |type |newtype |class |instance )/ ||
                $0 ~ /^[A-Za-z_][A-Za-z0-9_'\'']*[ \t]+(::|=)/) {
                lift[++nlift] = $0;
            } else if ($0 ~ /^[)\]}]/) {
                rest[++nrest] = "    " $0;
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
    ' "$1" > "$1.lifted"
    mv "$1.lifted" "$1"
}

validate_one() {
    relpath=$1
    cp "/snippet/$relpath" app/Main.hs
    lift_main app/Main.hs

    BUILD_LOG=$(mktemp)
    if ! cabal build >"$BUILD_LOG" 2>&1; then
        cat "$BUILD_LOG" >&2
        rm -f "$BUILD_LOG"
        return 1
    fi
    rm -f "$BUILD_LOG"

    if [ "$CHECK" = "parse" ]; then
        echo "feature flag evaluates to true"
        return 0
    fi

    LOG=$(mktemp)
    timeout --signal=TERM 60s cabal run hello-haskell-exe >"$LOG" 2>&1 &
    PID=$!
    deadline=$(( $(date +%s) + 55 ))
    if await_success_line "$LOG" "$PID" "$deadline"; then
        rm -f "$LOG"
        return 0
    fi
    kill -TERM "$PID" 2>/dev/null || true
    wait "$PID" 2>/dev/null || true
    dump_redacted "$LOG" >&2
    rm -f "$LOG"
    return 1
}

run_batch validate_one
