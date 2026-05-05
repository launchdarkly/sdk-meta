#!/bin/sh
# Validates the iOS snippet on macos-latest with xcodebuild + iOS
# Simulator. The snippet's AppDelegate + ViewController are dropped
# into a pre-baked Xcode project (generated from the scaffold's
# project.yml via xcodegen), pointed at the launchdarkly-ios-client-sdk
# Swift Package, and exercised via an XCTest case.
#
# `mode: native` — xcodebuild + iOS Simulator don't run inside Linux
# containers, so the CI cell sets runs-on: macos-latest.
set -eu

# The runner doesn't mount /harness-shared (no docker), so source the
# helpers via a relative path.
. "$(cd "$(dirname "$0")/../../../shared" && pwd)/lib.sh"

require_env LAUNCHDARKLY_MOBILE_KEY LAUNCHDARKLY_FLAG_KEY SNIPPET_ENTRYPOINT

SCAFFOLD="$(cd "$(dirname "$0")/../scaffold" && pwd)"

WORK=$(mktemp -d)
trap 'rm -rf "$WORK"' EXIT

cp -R "$SCAFFOLD"/. "$WORK"/
cp "$SNIPPET_DIR/AppDelegate.swift" "$WORK/Sources/AppDelegate.swift"
cp "$SNIPPET_DIR/ViewController.swift" "$WORK/Sources/ViewController.swift"

# Swift only allows `import` declarations at file scope. The init
# scaffold splices the snippet body inside a function (so the body's
# `LDClient.start(...)` runs with an Application lifecycle); any
# `import` lines that came along in the body need to be lifted out.
# Idempotent on files that don't have a misplaced import.
python3 - "$WORK/Sources/AppDelegate.swift" <<'PYEOF'
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
        continue  # drop from in-place position
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

cd "$WORK"

if ! command -v xcodegen >/dev/null 2>&1; then
    brew install xcodegen
fi
xcodegen generate

# Pick whichever iPhone simulator is currently installed on this
# runner. macos-latest's iPhone roster shifts per Xcode release —
# hardcoding "iPhone 15" stops working when GitHub bumps the image.
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
        # Prefer plain "iPhone <N>" over Pro/Max/Plus variants so the
        # destination string is shortest and most likely to match.
        m = re.match(r"^iPhone (\d+)$", name)
        if m and int(m.group(1)) > best_num:
            best_num = int(m.group(1))
            best = name
if best is None:
    # Fallback: any available iPhone, max version-suffix wins.
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

LOG=$(mktemp)

# `-resolvePackageDependencies` is an action mutually exclusive with
# `test`; xcodebuild silently runs only the resolve when both are
# passed. Resolve packages first, then run the test action.
xcodebuild -resolvePackageDependencies \
    -project HelloIOS.xcodeproj \
    -scheme HelloIOS \
    >>"$LOG" 2>&1

set +e
SIMCTL_CHILD_LAUNCHDARKLY_MOBILE_KEY="$LAUNCHDARKLY_MOBILE_KEY" \
SIMCTL_CHILD_LAUNCHDARKLY_FLAG_KEY="$LAUNCHDARKLY_FLAG_KEY" \
xcodebuild test \
    -project HelloIOS.xcodeproj \
    -scheme HelloIOS \
    -destination "$DESTINATION" \
    CODE_SIGNING_ALLOWED=NO \
    CODE_SIGN_IDENTITY="" \
    >>"$LOG" 2>&1
XCB_EXIT=$?
set -e

if grep -q "feature flag evaluates to true" "$LOG"; then
    grep -E "feature flag evaluates to true|validator: rendered" "$LOG" | head -3
    echo "validator: ok"
    exit 0
fi

echo "validator: did not see expected line: feature flag evaluates to true (xcodebuild exit=$XCB_EXIT)" >&2
echo "--- last 100 lines of xcodebuild output ---" >&2
tail -100 "$LOG" >&2
exit 1
