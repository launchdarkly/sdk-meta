#!/bin/sh
# Batch validator for Rust server SDK snippets. Reproduces gonfalon's
# `cargo new` + `cargo add` + run flow, but ONCE per job rather than once
# per snippet: the Dockerfile pre-baked /opt/hello-rust with the SDK +
# tokio + futures + transport dependency tree already compiled, so each
# snippet only recompiles the binary crate (the swapped src/main.rs) and
# links against the cached dependency objects.
#
# The Go runner stages every matching snippet under the bind-mounted
# /snippet dir and writes /snippet/manifest.tsv (one `<relpath>\t<label>`
# line per snippet). We loop over it in the single warm project. Exit
# non-zero if any snippet fails; the run continues past failures so one
# bad fragment doesn't hide the rest.
set -eu

. /harness-shared/lib.sh
require_env LAUNCHDARKLY_SDK_KEY LAUNCHDARKLY_FLAG_KEY SNIPPET_BATCH

cd /opt/hello-rust

# Version-pinned batch groups (rust-syntax-only-v1, etc.) set
# LD_RUST_SDK_VERSION via validation.env so older-API fragments compile
# against the SDK major they document. Re-pin and warm the cache once for
# the whole shard; the default (unpinned) group reuses the image's
# pre-built dependency tree untouched, keeping the common path a pure
# cache hit.
if [ -n "${LD_RUST_SDK_VERSION:-}" ]; then
    cargo add --quiet "launchdarkly-server-sdk@${LD_RUST_SDK_VERSION}"
    # Warm the re-pinned dependency tree once. If it doesn't compile, say so
    # loudly rather than letting every snippet fail later with a confusing
    # `cargo run` error; we still continue so each snippet's own result is
    # reported.
    if ! cargo build --quiet; then
        echo "validator: warm build for launchdarkly-server-sdk@${LD_RUST_SDK_VERSION} failed; snippet results below will reflect this" >&2
    fi
fi

validate_one() {
    relpath=$1
    cp "/snippet/$relpath" src/main.rs

    LOG=$(mktemp)
    timeout --signal=TERM 300s cargo run --quiet >"$LOG" 2>&1 &
    PID=$!
    deadline=$(( $(date +%s) + 290 ))
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
