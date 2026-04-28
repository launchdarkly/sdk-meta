#!/bin/sh
# Builds the staged Erlang snippet against the pre-baked rebar3 project,
# then runs `rebar3 eval` to start the gen_server, evaluate the flag,
# and print the EXAM-HELLO canonical line.
#
# The snippet is a gen_server module — the user-facing Get Started flow
# expects the user to launch `rebar3 shell` and call into it manually.
# For CI we synthesize the equivalent: ensure_all_started, sleep so the
# SDK has time to fetch flags, call hello_erlang_server:get/3, format
# the canonical line, halt.
set -eu

. /harness-shared/lib.sh
require_env LAUNCHDARKLY_SDK_KEY LAUNCHDARKLY_FLAG_KEY SNIPPET_ENTRYPOINT

cp "/snippet/$SNIPPET_ENTRYPOINT" /opt/hello_erlang/src/hello_erlang_server.erl

cd /opt/hello_erlang
rebar3 compile >/tmp/build.log 2>&1 \
    || { cat /tmp/build.log >&2; exit 1; }

LOG=$(mktemp)

# rebar3 doesn't ship `eval` in the bundled tasks, so drive `erl`
# directly. The compiled beams + transitive deps live under
# _build/default/lib/*/ebin. -noshell + -s init stop wraps the eval in
# a non-interactive session that exits when init:stop/0 fires.
EVAL_EXPR="application:ensure_all_started(hello_erlang),
timer:sleep(3000),
FlagKey = <<\"$LAUNCHDARKLY_FLAG_KEY\">>,
Result = hello_erlang_server:get(FlagKey, false, <<\"example-user-key\">>),
io:format(\"The ~s feature flag evaluates to ~p~n\", [FlagKey, Result]),
init:stop()."

timeout --signal=TERM 60s erl \
    -pa _build/default/lib/*/ebin \
    -noshell \
    -eval "$EVAL_EXPR" >"$LOG" 2>&1 &
PID=$!

deadline=$(( $(date +%s) + 55 ))
if await_success_line "$LOG" "$PID" "$deadline"; then
    exit 0
fi

kill -TERM "$PID" 2>/dev/null || true
wait "$PID" 2>/dev/null || true
fail_with_log "$LOG" "did not see expected line: feature flag evaluates to true"
