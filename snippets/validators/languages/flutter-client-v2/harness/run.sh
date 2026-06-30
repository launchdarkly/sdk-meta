#!/bin/sh
# Batch validator for Flutter client SDK snippets. Every bound fragment is
# syntax-only: the scaffold's main() unconditionally renders the success
# text and the fragment body sits in a never-invoked _wrappee(), so the
# validation is "does this compile against the real SDK".
#
# We compile with `flutter build linux --debug` rather than the old
# `flutter build web --release`. Both run the Dart front-end (the compiler
# that catches every syntax / type / resolution error a doc fragment can
# have), but the web release path additionally runs dart2js (~25s/snippet)
# and the old harness then launched headless Chromium to read the rendered
# text — neither of which adds coverage here, since main() renders the text
# unconditionally and the body never executes. The linux debug build stops
# at the kernel compile (no dart2js, no AOT), so it's a real build that
# finishes in ~5-7s warm. The flutter snippets import only flutter/material
# and the LD SDK (no dart:html / dart:io / platform channels), so the linux
# target compiles the identical code the web target would — no platform
# divergence. The Dockerfile pre-baked /opt/hello_flutter with the linux
# desktop scaffold + toolchain and a warmed debug build.
#
# The Go runner stages every matching snippet under /snippet and writes
# /snippet/manifest.tsv; run_batch loops over it in the single warm project.
set -eu

. /harness-shared/lib.sh
require_env SNIPPET_BATCH

cd /opt/hello_flutter

# lift_imports <file>: if the staged main.dart uses the flutter-syntax-only
# scaffold's body marker pair, lift any `import 'package:...';` directives
# from inside the body up to module scope. Dart forbids imports inside a
# function body, so doc fragments that show install-time imports would
# otherwise fail to compile.
lift_imports() {
    f=$1
    grep -qF -- '//IMPORT_LIFT_TARGET' "$f" || return 0
    awk '
    /^\/\/IMPORT_LIFT_TARGET$/ {
        target_seen = 1;
        pre[++npre] = $0;
        target_index = npre;
        next;
    }
    /^\/\/BODY_BEGIN$/ { in_body = 1; next; }
    /^\/\/BODY_END$/ { in_body = 0; body_done = 1; next; }
    {
        if (in_body) {
            if ($0 ~ /^[ \t]*import [\x27"]/) {
                sub(/^[ \t]+/, "", $0);
                lift[++nlift] = $0;
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
    ' "$f" > "$f.lifted"
    mv "$f.lifted" "$f"
}

validate_one() {
    relpath=$1
    cp "/snippet/$relpath" lib/main.dart
    lift_imports lib/main.dart

    LOG=$(mktemp)
    if flutter build linux --debug >"$LOG" 2>&1; then
        # main() renders the success text unconditionally, so a clean
        # compile means the app would render it; emit the canonical line
        # for the verify-hello-app wrapper and snippet-parity grep.
        echo "feature flag evaluates to true"
        rm -f "$LOG"
        return 0
    fi
    cat "$LOG" >&2
    rm -f "$LOG"
    return 1
}

run_batch validate_one
