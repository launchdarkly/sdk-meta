#!/bin/sh
# Batch validator for Java server SDK snippets. The Dockerfile pre-baked a
# warm maven project at /opt/hello-java (deps + plugins in ~/.m2, one compile
# done), so per-snippet we reset the source tree, drop the snippet's class in
# at its declared package, and run `mvn -o` (offline) — no per-snippet
# `clean`, dep re-resolution, or fat-jar assembly.
#
# The Go runner stages every matching snippet under /snippet and writes
# /snippet/manifest.tsv; run_batch loops over them in the single warm
# project (one container instead of one per snippet).
set -eu

. /harness-shared/lib.sh
require_env LAUNCHDARKLY_SDK_KEY LAUNCHDARKLY_FLAG_KEY SNIPPET_BATCH

PROJ=/opt/hello-java
cd "$PROJ"

# import_lift <file>: hoist `import …;` lines that appear after the
# IMPORT_LIFT_MARKER (i.e. inside the wrappee body) up to the marker — Java
# forbids imports inside method bodies.
import_lift() {
    grep -q 'IMPORT_LIFT_MARKER' "$1" || return 0
    python3 - "$1" <<'PYEOF'
import sys
fp = sys.argv[1]
with open(fp) as f:
    lines = f.read().splitlines()
marker_idx = next((i for i, l in enumerate(lines) if 'IMPORT_LIFT_MARKER' in l), None)
if marker_idx is None:
    sys.exit(0)
imports = []
for i in range(marker_idx + 1, len(lines)):
    stripped = lines[i].lstrip()
    if stripped.startswith('import ') and stripped.rstrip().endswith(';'):
        imp = stripped.rstrip()
        if imp not in imports:
            imports.append(imp)
        lines[i] = ''
if imports:
    lines = lines[:marker_idx] + imports + lines[marker_idx:]
with open(fp, 'w') as f:
    f.write('\n'.join(lines) + '\n')
PYEOF
}

validate_one() {
    relpath=$1
    idx=$(printf '%s' "$relpath" | cut -d/ -f1)
    unit="/snippet/$idx"
    ep=${relpath#*/}                       # entrypoint path within the unit
    entry="src/main/java/${ep#src/main/java/}"

    # Reset the source tree, then copy the whole staged source tree in — a
    # snippet may stage companions (e.g. the init runner's Main + Runner), so
    # copying only the entrypoint would drop them and break compilation.
    rm -rf src/main/java
    if [ -d "$unit/src/main/java" ]; then
        mkdir -p src/main/java
        cp -r "$unit/src/main/java/." src/main/java/
    else
        # No src tree (a bare body): place the entrypoint at its path.
        mkdir -p "$(dirname "$entry")"
        cp "$unit/$ep" "$entry"
    fi
    [ -f "$entry" ] || { echo "entrypoint not found: $entry" >&2; return 1; }

    import_lift "$entry"

    # Bodies without a package declaration want the tutorial layout, matching
    # the old harness.
    if ! head -1 "$entry" | grep -q '^package '; then
        prefixed=$(mktemp)
        printf 'package com.launchdarkly.tutorial;\n\n' > "$prefixed"
        cat "$entry" >> "$prefixed"
        mv "$prefixed" "$entry"
    fi

    PKG=$(grep -m1 '^package ' "$entry" | sed -e 's/^package //' -e 's/;.*//')
    CLS=$(basename "$entry" .java)
    MAIN_CLASS="$PKG.$CLS"

    LOG=$(mktemp)
    if ! mvn -o -B -q compile >"$LOG" 2>&1; then
        cat "$LOG" >&2
        rm -f "$LOG"
        return 1
    fi

    CI=1 timeout --signal=TERM 60s mvn -o -B -q exec:java -Dexec.mainClass="$MAIN_CLASS" >"$LOG" 2>&1 &
    PID=$!
    deadline=$(( $(date +%s) + 50 ))
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
