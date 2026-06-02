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
    if re.match(r"^\s*(override\s+)?fun\s+", line):
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

new_top = []
for imp in in_func_imports:
    if imp not in file_imports:
        new_top.append(imp)
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
