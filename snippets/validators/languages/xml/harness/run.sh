#!/bin/sh
# Validates a staged XML snippet by parsing it with xmllint. A clean
# parse means the file is well-formed XML.
#
# Parse-only -- no schema/DTD validation and no LD env. The snippet is
# never executed; on a successful parse we echo the EXAM-HELLO success
# line ourselves, mirroring the roku-client validator's approach.
set -eu

. /harness-shared/lib.sh
require_env SNIPPET_ENTRYPOINT

ENTRY="/snippet/$SNIPPET_ENTRYPOINT"
if [ ! -f "$ENTRY" ]; then
    echo "harness: entrypoint not found at $ENTRY" >&2
    exit 1
fi

LOG=$(mktemp)
if ! xmllint --noout "$ENTRY" >"$LOG" 2>&1; then
    fail_with_log "$LOG" "xmllint reported parse errors"
fi

echo "feature flag evaluates to true"
echo "validator: ok (xmllint well-formedness check succeeded)"
