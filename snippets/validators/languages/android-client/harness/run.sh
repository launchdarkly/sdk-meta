#!/bin/sh
# Validates the Android snippet under Robolectric (no emulator). The
# Dockerfile pre-bakes a hello-android scaffold + Robolectric test;
# per-validate we just swap the snippet's two Kotlin files in and run
# the JUnit test.
#
# Two checks flow through this same harness, dispatched on the
# optional `SNIPPET_CHECK` env var (set by the snippet's
# `validation.env` or by the validate dispatcher):
#
#   - SNIPPET_CHECK unset OR "runtime" (default): full Robolectric run
#     end-to-end against the LD streaming endpoint. Asserts the
#     canonical EXAM-HELLO line lands in the activity TextView.
#
#   - SNIPPET_CHECK=parse: parse-and-type-check the staged Kotlin
#     file against the pre-baked `launchdarkly-android-client-sdk`
#     aar + AndroidX classpath via `./gradlew compileDebugKotlin`,
#     then bail out. Used for doc fragments (e.g. the
#     `sdk-docs/*-kotlin.snippet.md` reference snippets) that
#     aren't standalone-runnable but should still be checked for
#     syntactic correctness and type resolution against the real
#     android-client SDK surface.
set -eu

. /harness-shared/lib.sh
require_env LAUNCHDARKLY_FLAG_KEY SNIPPET_ENTRYPOINT
CHECK="${SNIPPET_CHECK:-runtime}"
if [ "$CHECK" = "runtime" ]; then
    require_env LAUNCHDARKLY_MOBILE_KEY
fi

# The snippet declares main-application + main-activity as separate
# files. Both come through as part of the staging dir — copy whichever
# .kt files /snippet has into the scaffold's main source tree.
SCAFFOLD=/opt/hello-android
PKG_DIR="${SCAFFOLD}/app/src/main/java/com/launchdarkly/hello_android"

for f in /snippet/app/src/main/java/com/launchdarkly/hello_android/*.kt; do
    [ -f "$f" ] || continue
    cp "$f" "${PKG_DIR}/$(basename "$f")"
done
# Java sources flow through the same path. The gradle project's
# `src/main/java/` source set accepts both, and `compileDebug*`
# invocations below cover Kotlin and Java separately.
for f in /snippet/app/src/main/java/com/launchdarkly/hello_android/*.java; do
    [ -f "$f" ] || continue
    cp "$f" "${PKG_DIR}/$(basename "$f")"
done

# Two transforms are needed before kotlinc will accept the staged
# Kotlin file:
#   1. Lift any `import` statements out of the function body to file
#      scope (Kotlin only allows imports between `package` and the
#      first top-level declaration).
#   2. In runtime mode only: rewrite the body's `this@BaseApplication`
#      reference (the gonfalon snippet's literal Application name)
#      to `this@MainApplication`, which is the class name the init
#      scaffold produces and the existing HelloAppTest expects via
#      `@Config(application = MainApplication::class)`. The
#      kotlin-syntax-only scaffold declares its class as
#      `BaseApplication` directly so the body's labeled-this
#      resolves without substitution; skip the rewrite there.
# Idempotent on files that don't have a misplaced import or the
# `BaseApplication` token.
ENTRY_KT="${PKG_DIR}/$(basename "$SNIPPET_ENTRYPOINT")"
if [ -f "$ENTRY_KT" ]; then
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
fi

# If the staged file contains the TYPE_LIFT_TARGET marker, move any
# brace-balanced type declarations found between the BODY_BEGIN/BODY_END
# markers up to the target at SnippetActivity member scope. Java rejects
# access modifiers on local classes, so doc fragments that define a
# `public class` alongside statements (e.g. a hook implementation plus
# the configuration that registers it) would otherwise fail to compile
# inside `onCreate()`. As nested member classes they compile, and they
# shadow any same-named file-scope stub for the statement residue left
# behind. Bodies without type declarations are untouched. Brace counting
# is line-based and does not see braces in string literals; the doc
# fragments this serves don't contain such literals.
if [ -f "$ENTRY_KT" ] && grep -q 'TYPE_LIFT_TARGET' "$ENTRY_KT"; then
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

cd "${SCAFFOLD}"

if [ "$CHECK" = "parse" ]; then
    # Compile-only path: kotlinc + javac against the real
    # android-client SDK aar + AndroidX runtime, no Robolectric run.
    # We invoke both `compileDebugKotlin` and `compileDebugJavaWithJavac`
    # so Kotlin-only and Java-only fragments are both covered. Each
    # task is cheap (no APK assembly, no lint), and gradle skips any
    # task whose inputs (the source set) didn't change.
    LOG=$(mktemp)
    if timeout --signal=TERM 600s ./gradlew --no-daemon \
            compileDebugKotlin compileDebugJavaWithJavac \
            --console=plain >"$LOG" 2>&1; then
        echo "feature flag evaluates to true"
        echo "validator: ok (compileDebug{Kotlin,JavaWithJavac} succeeded)"
        exit 0
    fi
    fail_with_log "$LOG" "android compile failed"
fi

LOG=$(mktemp)
timeout --signal=TERM 600s ./gradlew --no-daemon \
        testDebugUnitTest --tests='*HelloAppTest*' --console=plain \
        >"$LOG" 2>&1 &
PID=$!

deadline=$(( $(date +%s) + 590 ))
if await_success_line "$LOG" "$PID" "$deadline"; then
    exit 0
fi

kill -TERM "$PID" 2>/dev/null || true
wait "$PID" 2>/dev/null || true
fail_with_log "$LOG" "did not see expected line: feature flag evaluates to true"
