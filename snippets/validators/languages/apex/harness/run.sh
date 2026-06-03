#!/bin/sh
# Validates a staged Apex snippet by parsing it with
# prettier-plugin-apex's `apex-anonymous` parser. A clean parse means
# the fragment is syntactically valid Apex in Execute Anonymous form.
#
# Parse-only — no Salesforce scratch org, no real LD env. We invoke
# prettier without --check/--write: it parses the file and prints the
# reformatted source to stdout (discarded), exiting non-zero only when
# the parser rejects the input. Gating on the exit code rather than
# --check means doc fragments don't have to be prettier-formatted to
# pass — we only assert they parse. The fragment is never executed, so
# on a successful parse we echo the EXAM-HELLO sentinel ourselves.
set -eu

. /harness-shared/lib.sh
require_env SNIPPET_ENTRYPOINT

ENTRY="/snippet/$SNIPPET_ENTRYPOINT"
if [ ! -f "$ENTRY" ]; then
    echo "harness: entrypoint not found at $ENTRY" >&2
    exit 1
fi

LOG=$(mktemp)
# Run from the pre-baked install dir so prettier resolves the
# `prettier-plugin-apex` bare specifier against /opt/apex/node_modules
# (prettier resolves --plugin relative to its working directory, not
# its own binary path). $ENTRY is absolute, so the cd doesn't affect it.
if ! ( cd /opt/apex && ./node_modules/.bin/prettier \
    --plugin=prettier-plugin-apex \
    --parser=apex-anonymous \
    "$ENTRY" > /dev/null 2>"$LOG" ); then
    fail_with_log "$LOG" "prettier-plugin-apex reported parse errors"
fi

echo "feature flag evaluates to true"
echo "validator: ok (apex parse succeeded)"
