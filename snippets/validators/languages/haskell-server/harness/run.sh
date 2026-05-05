#!/bin/sh
# Runs the staged Haskell snippet against a real LaunchDarkly environment.
# The Dockerfile pre-bootstrapped a cabal project at /opt/hello-haskell
# with launchdarkly-server-sdk + text already compiled. Per-validate just
# swaps in the user's Main.hs and does an incremental rebuild.
set -eu

. /harness-shared/lib.sh
require_env LAUNCHDARKLY_SDK_KEY LAUNCHDARKLY_FLAG_KEY SNIPPET_ENTRYPOINT

cp "/snippet/$SNIPPET_ENTRYPOINT" /opt/hello-haskell/app/Main.hs

# If the staged Main.hs contains the TOP_LIFT_MARKER comment (used by
# the haskell-syntax-only scaffold), lift any line below the marker
# that starts at column 0 with a top-level keyword (`import`, `data`,
# `type`, `newtype`, `class`, `instance`) or that looks like a
# type-signature / top-level binding (`identifier ::` or `identifier
# =` at column 0) up above the marker. Haskell forbids these
# constructs inside a `do` block, so doc fragments that mix top-level
# decls with in-block code would otherwise fail to compile.
if grep -q 'TOP_LIFT_MARKER' /opt/hello-haskell/app/Main.hs; then
    awk '
    BEGIN { state = "head"; }
    /TOP_LIFT_MARKER/ {
        # Buffer everything after the marker, then re-emit with lifted
        # lines moved above the marker.
        marker = NR;
        print;
        state = "body";
        next;
    }
    state == "head" { print; next; }
    state == "body" {
        # First column non-whitespace and matches a top-level keyword?
        if (match($0, /^(import |data |type |newtype |class |instance )/) ||
            match($0, /^[A-Za-z_][A-Za-z0-9_'\'']*[[:space:]]+(::|=)/)) {
            top[++ntop] = $0;
        } else {
            rest[++nrest] = $0;
        }
    }
    END {
        for (i = 1; i <= ntop; i++) print top[i];
        for (i = 1; i <= nrest; i++) print rest[i];
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
