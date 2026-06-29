#!/bin/sh
# Batch validator for Android snippets under Robolectric / kotlinc (no
# emulator). The Dockerfile pre-bakes a hello-android scaffold + Robolectric
# test and warms the gradle build cache; per-snippet we reset the package dir
# to the baseline scaffold, swap the snippet's Kotlin/Java files in, apply
# the import/type transforms, and compile (or, in runtime mode, run the
# Robolectric test).
#
# Unlike the old one-invocation-per-snippet path, the Go runner stages every
# matching snippet under /snippet and writes /snippet/manifest.tsv; run_batch
# loops over them in a single container. Crucially we let the gradle daemon
# stay alive across the loop (no --no-daemon), so only the first snippet pays
# JVM + gradle startup and the rest reuse the warm daemon + build cache —
# that, not the compile itself, was the dominant per-snippet cost.
#
# Two checks flow through this same harness, dispatched on SNIPPET_CHECK
# (snippets of one kind share a batch group, so a whole run is one mode):
#
#   - runtime (default): full Robolectric run end-to-end against the LD
#     streaming endpoint; asserts the canonical EXAM-HELLO line.
#   - parse: kotlinc + javac against the real android-client SDK aar +
#     AndroidX classpath via compileDebug{Kotlin,JavaWithJavac}, no run.
set -eu

. /harness-shared/lib.sh
require_env LAUNCHDARKLY_FLAG_KEY SNIPPET_BATCH
CHECK="${SNIPPET_CHECK:-runtime}"
if [ "$CHECK" = "runtime" ]; then
    require_env LAUNCHDARKLY_MOBILE_KEY
fi

SCAFFOLD=/opt/hello-android
PKG_DIR="${SCAFFOLD}/app/src/main/java/com/launchdarkly/hello_android"

# Snapshot the baseline package dir once. Each snippet may add or replace
# .kt/.java files; resetting to this baseline before staging the next
# snippet keeps stale files from a prior fragment out of the next compile.
BASELINE=/tmp/pkg-baseline
rm -rf "$BASELINE"
cp -a "$PKG_DIR" "$BASELINE"

cd "$SCAFFOLD"

# stage_files <unit-dir>: reset the package dir to baseline, then copy the
# snippet's staged Kotlin/Java files in. The snippet declares its source
# files under app/src/main/java/com/launchdarkly/hello_android/.
stage_files() {
    unit=$1
    rm -rf "$PKG_DIR"
    cp -a "$BASELINE" "$PKG_DIR"
    for f in "$unit"/app/src/main/java/com/launchdarkly/hello_android/*.kt \
             "$unit"/app/src/main/java/com/launchdarkly/hello_android/*.java; do
        [ -f "$f" ] || continue
        cp "$f" "${PKG_DIR}/$(basename "$f")"
    done
}

# transform_entry <entry-file>: applies the two source rewrites the staged
# entry file needs before kotlinc/javac will accept it.
#   1. Lift `import` statements out of the function body to file scope
#      (Kotlin/Java only allow imports before the first declaration).
#   2. In runtime mode only: rewrite the body's `this@BaseApplication`
#      (the gonfalon snippet's literal Application name) to
#      `this@MainApplication`, the class HelloAppTest's
#      `@Config(application = MainApplication::class)` expects. The
#      kotlin-syntax-only scaffold names its class `BaseApplication`
#      directly, so the parse path skips this rewrite.
#   3. If the file carries the TYPE_LIFT_TARGET marker, hoist
#      brace-balanced type declarations out of the body to member scope.
transform_entry() {
    ENTRY_KT=$1
    [ -f "$ENTRY_KT" ] || return 0
    SUBSTITUTE_BASE_APP=1
    if [ "$CHECK" != "runtime" ]; then
        SUBSTITUTE_BASE_APP=0
    fi
    SUBSTITUTE_BASE_APP="$SUBSTITUTE_BASE_APP" python3 - "$ENTRY_KT" <<'PYEOF'
import os, re, sys
path = sys.argv[1]
with open(path) as f:
    text = f.read()

# Pull `import com.…` lines out of the function body (anything after
# the first `fun ` or `override fun`), dedup against module-scope
# imports already present, and re-insert them at the top of the
# imports block.
lines = text.splitlines()
file_imports = set()
in_func_imports = []
saw_func = False
out = []
for line in lines:
    s = line.strip()
    # Mark the end of the file-scope import region. Kotlin bodies open
    # with `fun`/`override fun`; Java bodies have no `fun`, so also
    # trip on the first type declaration (`class`/`object`/`interface`/
    # `enum`). In both languages every legal file-scope import precedes
    # the first such line, so any `import` after it is a misplaced body
    # import to hoist.
    if re.match(r"^\s*(override\s+)?fun\s+", line) or re.match(
        r"^\s*((public|private|protected|final|abstract|open|internal|sealed|data|static)\s+)*(class|object|interface|enum)\s+",
        line,
    ):
        saw_func = True
    if saw_func:
        m = re.match(r"^\s*(import\s+[A-Za-z_][A-Za-z0-9_.]*\*?\s*;?\s*)$", line)
        if m:
            in_func_imports.append(m.group(1).rstrip(';').strip())
            continue
    if not saw_func:
        m = re.match(r"^\s*(import\s+[A-Za-z_][A-Za-z0-9_.]*\*?\s*;?\s*)$", line)
        if m:
            file_imports.add(m.group(1).rstrip(';').strip())
    out.append(line)

# Imports are collected with their trailing `;` stripped (so Kotlin and
# Java forms dedup against each other). Java requires the semicolon when
# re-inserted at file scope; Kotlin must not have one.
is_java = path.endswith(".java")
new_top = []
for imp in in_func_imports:
    if imp not in file_imports:
        new_top.append(imp + ";" if is_java else imp)
        file_imports.add(imp)

if new_top:
    insert_at = 0
    for i, line in enumerate(out):
        if line.strip().startswith("import ") or line.strip().startswith("package "):
            insert_at = i + 1
        elif line.strip() and insert_at:
            break
    out[insert_at:insert_at] = new_top

new_text = "\n".join(out) + ("\n" if text.endswith("\n") else "")
if os.environ.get("SUBSTITUTE_BASE_APP") == "1":
    new_text = re.sub(r"\bBaseApplication\b", "MainApplication", new_text)

with open(path, "w") as f:
    f.write(new_text)
PYEOF

    # If the staged file contains the TYPE_LIFT_TARGET marker, move any
    # brace-balanced type declarations found between the BODY_BEGIN/BODY_END
    # markers up to the target at SnippetActivity member scope. Java rejects
    # access modifiers on local classes, so doc fragments that define a
    # `public class` alongside statements would otherwise fail to compile
    # inside `onCreate()`. As nested member classes they compile, and they
    # shadow any same-named file-scope stub for the statement residue left
    # behind.
    if grep -q 'TYPE_LIFT_TARGET' "$ENTRY_KT"; then
        python3 - "$ENTRY_KT" <<'PYEOF'
import re
import sys

fp = sys.argv[1]
with open(fp) as f:
    lines = f.read().splitlines()

target_idx = next((i for i, l in enumerate(lines) if 'TYPE_LIFT_TARGET' in l), None)
begin_idx = next((i for i, l in enumerate(lines) if 'BODY_BEGIN' in l), None)
end_idx = next((i for i, l in enumerate(lines) if 'BODY_END' in l), None)
if target_idx is None or begin_idx is None or end_idx is None:
    sys.exit(0)

decl_re = re.compile(
    r'^\s*((public|private|protected|static|final|abstract)\s+)*'
    r'(class|interface|enum)\s+\w')

lifted = []
i = begin_idx + 1
depth = 0
while i < end_idx:
    line = lines[i]
    if depth == 0 and decl_re.match(line):
        block = []
        bdepth = 0
        seen_brace = False
        while i < end_idx:
            block.append(lines[i])
            bdepth += lines[i].count('{') - lines[i].count('}')
            if '{' in lines[i] or '}' in lines[i]:
                seen_brace = True
            lines[i] = ''
            i += 1
            if seen_brace and bdepth == 0:
                break
        lifted.extend(block)
        continue
    depth += line.count('{') - line.count('}')
    i += 1

if lifted:
    lines = lines[:target_idx + 1] + lifted + lines[target_idx + 1:]
with open(fp, 'w') as f:
    f.write('\n'.join(lines) + '\n')
PYEOF
    fi
}

validate_parse() {
    LOG=$(mktemp)
    if timeout --signal=TERM 600s ./gradlew \
            compileDebugKotlin compileDebugJavaWithJavac \
            --console=plain >"$LOG" 2>&1; then
        echo "feature flag evaluates to true"
        echo "validator: ok (compileDebug{Kotlin,JavaWithJavac} succeeded)"
        rm -f "$LOG"
        return 0
    fi
    cat "$LOG" >&2
    rm -f "$LOG"
    return 1
}

validate_runtime() {
    LOG=$(mktemp)
    timeout --signal=TERM 600s ./gradlew \
            testDebugUnitTest --tests='*HelloAppTest*' --console=plain \
            >"$LOG" 2>&1 &
    PID=$!
    deadline=$(( $(date +%s) + 590 ))
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

validate_one() {
    relpath=$1
    idx=$(printf '%s' "$relpath" | cut -d/ -f1)
    stage_files "/snippet/$idx"
    transform_entry "${PKG_DIR}/$(basename "$relpath")"

    if [ "$CHECK" = "parse" ]; then
        validate_parse
    else
        validate_runtime
    fi
}

run_batch validate_one
