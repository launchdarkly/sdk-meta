#!/bin/sh
# Runs the staged PHP snippet against a real LaunchDarkly environment.
set -eu

. /harness-shared/lib.sh
require_env LAUNCHDARKLY_SDK_KEY LAUNCHDARKLY_FLAG_KEY SNIPPET_ENTRYPOINT

WORK=$(mktemp -d)
trap 'rm -rf "$WORK"' EXIT
cp -r /snippet/. "$WORK/"
cd "$WORK"

# PHP `use` (import) statements are only legal at file/namespace scope,
# never inside a function body. The syntax-only scaffold wraps the
# wrappee in `function _wrappee() { ... }`, so a fragment that opens
# with `use ...;` lines (e.g. the migration config example) would fail
# to parse. Lift any `use ...;` lines that appear after the first
# `function ` declaration up to just after the opening `<?php` tag.
# Keyed on the function boundary so runtime snippets (no wrapping
# function) are untouched.
python3 - "$SNIPPET_ENTRYPOINT" <<'PYEOF'
import re, sys
fp = sys.argv[1]
lines = open(fp).read().splitlines()
func_idx = next((i for i, l in enumerate(lines) if re.match(r'\s*function\s', l)), None)
if func_idx is not None:
    uses, rest = [], []
    for i, l in enumerate(lines):
        if i > func_idx and re.match(r'\s*use\s+\\?[A-Za-z_][\\A-Za-z0-9_]*.*;\s*$', l):
            uses.append(l.strip())
        else:
            rest.append(l)
    if uses:
        php_idx = next((i for i, l in enumerate(rest) if l.strip().startswith('<?php')), -1)
        rest[php_idx + 1:php_idx + 1] = uses
        open(fp, 'w').write('\n'.join(rest) + '\n')
PYEOF

# composer require <pkg>... per requirements line. The image bundles a
# system-wide `composer` binary, so we use it directly rather than
# bootstrapping composer.phar like the snippet does.
if [ -f requirements.txt ]; then
    pkgs=""
    while IFS= read -r line; do
        [ -z "$line" ] && continue
        pkgs="$pkgs $line"
    done < requirements.txt
    if [ -n "$pkgs" ]; then
        # shellcheck disable=SC2086
        composer require --quiet --no-interaction --no-progress $pkgs
    fi
fi

LOG=$(mktemp)

# PHP snippet loops forever with sleep(1); time out and SIGTERM after match.
timeout --signal=TERM 90s php -d output_buffering=Off "$SNIPPET_ENTRYPOINT" >"$LOG" 2>&1 &
PID=$!

deadline=$(( $(date +%s) + 80 ))
if await_success_line "$LOG" "$PID" "$deadline"; then
    exit 0
fi

kill -TERM "$PID" 2>/dev/null || true
wait "$PID" 2>/dev/null || true
fail_with_log "$LOG" "did not see expected line: feature flag evaluates to true"
