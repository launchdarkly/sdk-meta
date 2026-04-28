#!/bin/sh
# Runs the staged HTML snippet in headless Chromium and watches the page
# text for the EXAM-HELLO success line.
set -eu

. /harness-shared/lib.sh
require_env LAUNCHDARKLY_CLIENT_SIDE_ID LAUNCHDARKLY_FLAG_KEY SNIPPET_ENTRYPOINT

# The check.js harness reads SNIPPET_ENTRYPOINT and locates /snippet/<file>.
# All output goes to stdout/stderr — match handling lives in check.js since
# the success criterion is the page DOM text, not the program log.
exec node /harness/check.js
