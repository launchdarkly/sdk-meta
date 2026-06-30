#!/bin/sh
# Batch validator for iOS snippets on macos-latest (xcodebuild + iOS
# Simulator). Each snippet's AppDelegate + ViewController are dropped into a
# pre-baked Xcode project (generated from the scaffold's project.yml via
# xcodegen) pointed at the launchdarkly-ios-client-sdk Swift Package.
#
# `mode: native` — xcodebuild + the Simulator don't run inside Linux
# containers, so the CI cell sets runs-on: macos-latest.
#
# The old path did this once per snippet: xcodegen generate +
# `-resolvePackageDependencies` (which resolves AND builds the LD SDK
# package) + `xcodebuild test` (which boots a simulator). For the ~80
# syntax-only sdk-docs fragments that was ~90s each, almost all of it
# re-resolving/re-building the SDK and booting a simulator that only exists
# to print a hardcoded line. Batch mode does the project setup and SPM
# resolve ONCE, into a shared DerivedData, then loops the staged snippets:
#
#   - parse (sdk-docs / experimentation, syntax-only): `xcodebuild build`
#     against the iphonesimulator SDK — a compile/type-check, no simulator
#     boot. The wrappee body lives in a never-instantiated function, so a
#     clean compile is the whole signal; we emit the canonical line.
#   - runtime (init): `xcodebuild test`, which boots the simulator and runs
#     LDClient.start end-to-end, then greps the captured log.
#
# The Go runner stages every matching snippet under $SNIPPET_DIR and writes
# $SNIPPET_BATCH (the manifest); run_batch loops over it in the warm project.
set -eu

# The runner doesn't mount /harness-shared (no docker), so source the
# helpers via a relative path.
. "$(cd "$(dirname "$0")/../../../shared" && pwd)/lib.sh"

require_env LAUNCHDARKLY_FLAG_KEY SNIPPET_BATCH
CHECK="${SNIPPET_CHECK:-runtime}"
if [ "$CHECK" = "runtime" ]; then
    require_env LAUNCHDARKLY_MOBILE_KEY
fi

SCAFFOLD="$(cd "$(dirname "$0")/../scaffold" && pwd)"

WORK=$(mktemp -d)
trap 'rm -rf "$WORK"' EXIT
cp -R "$SCAFFOLD"/. "$WORK"/
cd "$WORK"

# Snapshot the scaffold's baseline Sources so each snippet starts from a
# clean tree. The project compiles every file under Sources/, so without a
# reset a differently-named .swift file from an earlier fragment could
# linger and affect a later compile.
BASELINE_SOURCES=$(mktemp -d)
cp -R "$WORK/Sources/." "$BASELINE_SOURCES/"

if ! command -v xcodegen >/dev/null 2>&1; then
    brew install xcodegen
fi
xcodegen generate

DERIVED="$WORK/DerivedData"

# Resolve + build the Swift Package dependencies ONCE into the shared
# DerivedData. Every per-snippet build below reuses these built products
# instead of re-resolving and re-compiling the SDK.
xcodebuild -resolvePackageDependencies \
    -project HelloIOS.xcodeproj \
    -scheme HelloIOS \
    -derivedDataPath "$DERIVED" \
    >/tmp/resolve.log 2>&1 || { cat /tmp/resolve.log >&2; exit 1; }

# Runtime mode boots a simulator; pick whichever iPhone is installed on
# this runner (the roster shifts per Xcode release). Parse mode compiles
# against the iphonesimulator SDK without booting anything, so it skips
# this.
DESTINATION=""
if [ "$CHECK" = "runtime" ]; then
    SIM_NAME=$(xcrun simctl list devices available --json | python3 -c '
import sys, json, re
data = json.load(sys.stdin)
best = None
best_num = -1
for runtime, devs in data.get("devices", {}).items():
    if "iOS" not in runtime:
        continue
    for dev in devs:
        name = dev.get("name", "")
        if not dev.get("isAvailable", False):
            continue
        if not name.startswith("iPhone"):
            continue
        m = re.match(r"^iPhone (\d+)$", name)
        if m and int(m.group(1)) > best_num:
            best_num = int(m.group(1))
            best = name
if best is None:
    cand = []
    for runtime, devs in data.get("devices", {}).items():
        if "iOS" not in runtime:
            continue
        for dev in devs:
            if dev.get("isAvailable") and dev.get("name", "").startswith("iPhone"):
                cand.append(dev["name"])
    best = sorted(cand)[-1] if cand else "iPhone 16"
print(best)
')
    DESTINATION="platform=iOS Simulator,name=$SIM_NAME"
    echo "validator: targeting $DESTINATION"
fi

# lift_imports <file>: Swift only allows `import` at file scope. The init
# scaffold splices the snippet body inside a function, so any `import` lines
# that came along in the body must be lifted out. Idempotent.
lift_imports() {
    python3 - "$1" <<'PYEOF'
import re, sys
path = sys.argv[1]
with open(path) as f:
    text = f.read()

lines = text.splitlines()
file_imports = set()
in_func_imports = []
saw_func = False
out = []
for line in lines:
    s = line.strip()
    if s.startswith("func "):
        saw_func = True
    if saw_func and re.match(r"^\s*import\s+[A-Za-z_][A-Za-z0-9_.]*\s*$", line):
        m = re.match(r"^\s*import\s+([A-Za-z_][A-Za-z0-9_.]*)\s*$", line)
        in_func_imports.append(m.group(1))
        continue
    if not saw_func:
        m = re.match(r"^\s*import\s+([A-Za-z_][A-Za-z0-9_.]*)\s*$", line)
        if m:
            file_imports.add(m.group(1))
    out.append(line)

new_top = []
for mod in in_func_imports:
    if mod not in file_imports:
        new_top.append(f"import {mod}")
        file_imports.add(mod)

if new_top:
    insert_at = 0
    for i, line in enumerate(out):
        if line.strip().startswith("import "):
            insert_at = i + 1
        elif line.strip() and insert_at:
            break
    out[insert_at:insert_at] = new_top

with open(path, "w") as f:
    f.write("\n".join(out) + ("\n" if text.endswith("\n") else ""))
PYEOF
}

stage_snippet() {
    idx=$1
    # Reset Sources to the scaffold baseline so a prior snippet's files don't
    # leak into this compile, then copy in whatever .swift files the unit
    # provides (an AppDelegate + its ViewController companion).
    rm -rf "$WORK/Sources"
    mkdir -p "$WORK/Sources"
    cp -R "$BASELINE_SOURCES/." "$WORK/Sources/"
    for f in "$SNIPPET_DIR/$idx"/*.swift; do
        [ -f "$f" ] || continue
        cp "$f" "$WORK/Sources/$(basename "$f")"
    done
    lift_imports "$WORK/Sources/AppDelegate.swift"
}

validate_one() {
    relpath=$1
    idx=$(printf '%s' "$relpath" | cut -d/ -f1)
    stage_snippet "$idx"

    LOG=$(mktemp)
    if [ "$CHECK" = "runtime" ]; then
        SIMCTL_CHILD_LAUNCHDARKLY_MOBILE_KEY="$LAUNCHDARKLY_MOBILE_KEY" \
        SIMCTL_CHILD_LAUNCHDARKLY_FLAG_KEY="$LAUNCHDARKLY_FLAG_KEY" \
        xcodebuild test \
            -project HelloIOS.xcodeproj \
            -scheme HelloIOS \
            -destination "$DESTINATION" \
            -derivedDataPath "$DERIVED" \
            CODE_SIGNING_ALLOWED=NO \
            CODE_SIGN_IDENTITY="" \
            >"$LOG" 2>&1 || true
        if grep -q "feature flag evaluates to true" "$LOG"; then
            # Re-emit the matched line on stdout: the verify-hello-app
            # wrapper greps the command's combined output for it, and the
            # xcodebuild log went to $LOG, not stdout.
            grep -E "feature flag evaluates to true" "$LOG" | head -1
            rm -f "$LOG"
            return 0
        fi
        tail -100 "$LOG" >&2
        rm -f "$LOG"
        return 1
    fi

    # parse mode: compile-only against the simulator SDK, no boot.
    if xcodebuild build \
        -project HelloIOS.xcodeproj \
        -scheme HelloIOS \
        -sdk iphonesimulator \
        -destination "generic/platform=iOS Simulator" \
        -derivedDataPath "$DERIVED" \
        CODE_SIGNING_ALLOWED=NO \
        CODE_SIGN_IDENTITY="" \
        >"$LOG" 2>&1; then
        # The wrappee body never executes; a clean compile is the signal.
        echo "feature flag evaluates to true"
        rm -f "$LOG"
        return 0
    fi
    tail -100 "$LOG" >&2
    rm -f "$LOG"
    return 1
}

run_batch validate_one
