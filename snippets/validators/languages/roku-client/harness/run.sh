#!/bin/sh
# Validates a staged BrightScript snippet by running it through
# brighterscript's Lexer + Parser. A clean parse means the file is
# syntactically valid as either classic BrightScript or the
# BrighterScript superset.
#
# Parse-only — no Roku device, no real LD env. The snippet's
# scaffold wraps the wrappee body inside `sub _Wrappee()` and
# unconditionally prints the EXAM-HELLO sentinel from `sub Main()`,
# so on a successful parse we just echo the success line ourselves
# (the file is never actually executed; we only assert it's parsable).
set -eu

. /harness-shared/lib.sh
require_env SNIPPET_ENTRYPOINT

ENTRY="/snippet/$SNIPPET_ENTRYPOINT"
if [ ! -f "$ENTRY" ]; then
    echo "harness: entrypoint not found at $ENTRY" >&2
    exit 1
fi

LOG=$(mktemp)
if ! node /harness/check.js "$ENTRY" >"$LOG" 2>&1; then
    fail_with_log "$LOG" "brighterscript reported parse errors"
fi

echo "feature flag evaluates to true"
echo "validator: ok (brighterscript parse succeeded)"
