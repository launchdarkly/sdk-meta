#!/bin/sh
# Validates the Android snippet under Robolectric (no emulator). The
# Dockerfile pre-bakes a hello-android scaffold + Robolectric test;
# per-validate we just swap the snippet's two Kotlin files in and run
# the JUnit test.
set -eu

. /harness-shared/lib.sh
require_env LAUNCHDARKLY_MOBILE_KEY LAUNCHDARKLY_FLAG_KEY SNIPPET_ENTRYPOINT

# The snippet declares main-application + main-activity as separate
# files. Both come through as part of the staging dir — copy whichever
# .kt files /snippet has into the scaffold's main source tree.
SCAFFOLD=/opt/hello-android
PKG_DIR="${SCAFFOLD}/app/src/main/java/com/launchdarkly/hello_android"

for f in /snippet/app/src/main/java/com/launchdarkly/hello_android/*.kt; do
    [ -f "$f" ] || continue
    cp "$f" "${PKG_DIR}/$(basename "$f")"
done

# The init scaffold's MainApplication.kt splices the snippet body
# inside `onCreate()`. Two transforms are needed before kotlinc will
# accept it:
#   1. Lift any `import` statements out of the function body to file
#      scope (Kotlin only allows imports between `package` and the
#      first top-level declaration).
#   2. Rewrite the body's `this@BaseApplication` reference (the
#      gonfalon snippet's literal Application name) to
#      `this@MainApplication`, which is the class name this scaffold
#      produces and the existing HelloAppTest expects via
#      `@Config(application = MainApplication::class)`.
# Idempotent on files that don't have a misplaced import or the
# `BaseApplication` token.
APP_KT="${PKG_DIR}/MainApplication.kt"
if [ -f "$APP_KT" ]; then
    python3 - "$APP_KT" <<'PYEOF'
import re, sys
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

# Substitute BaseApplication for MainApplication so the body's
# `this@BaseApplication` resolves to the class this scaffold defines.
new_text = "\n".join(out) + ("\n" if text.endswith("\n") else "")
new_text = re.sub(r"\bBaseApplication\b", "MainApplication", new_text)

with open(path, "w") as f:
    f.write(new_text)
PYEOF
fi

cd "${SCAFFOLD}"

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
