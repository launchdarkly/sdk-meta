# Shared helpers for validator harnesses. Source from each language's
# harness/run.sh:
#
#   . /harness-shared/lib.sh
#
# Or, for native (non-docker) harnesses, the per-language run.sh references
# this via a relative path under validators/shared/lib.sh.

# require_env <NAME>... exits 1 with a clear message if any var is unset.
require_env() {
    for name in "$@"; do
        eval "val=\${$name-}"
        if [ -z "$val" ]; then
            echo "$name not set" >&2
            exit 1
        fi
    done
}

# await_success_line <log-file> <pid> <deadline-epoch>
# Returns 0 once the log file contains the EXAM-HELLO success line, or 1
# when the deadline elapses (or the program exits before matching).
#
# The success regex is intentionally lenient on framing — different SDK
# hello-worlds quote the flag key differently ("sample-feature",
# 'sample-feature', sample-feature) and Python prints `True` while every
# other language prints `true`. The phrase `feature flag evaluates to
# [Tt]rue` is the one canonical fragment all of them emit on a successful
# init+evaluation against the EXAM-HELLO `sample-feature` flag.
await_success_line() {
    log=$1
    pid=$2
    deadline=$3
    while [ "$(date +%s)" -lt "$deadline" ]; do
        if grep -E "feature flag evaluates to [Tt]rue" "$log" >/dev/null 2>&1; then
            kill -TERM "$pid" 2>/dev/null || true
            wait "$pid" 2>/dev/null || true
            grep -E "feature flag evaluates to [Tt]rue" "$log" | head -1
            echo "validator: ok"
            return 0
        fi
        if ! kill -0 "$pid" 2>/dev/null; then
            break
        fi
        sleep 0.2
    done
    return 1
}

# dump_redacted <log-file>
# Prints the log to stderr with LAUNCHDARKLY_SDK_KEY (and optionally
# LAUNCHDARKLY_MOBILE_KEY / LAUNCHDARKLY_CLIENT_SIDE_ID) replaced with a
# placeholder. Defense in depth — today's snippets never echo a key, but a
# future snippet could and this log gets piped into CI output.
dump_redacted() {
    log=$1
    # Build a sed program that redacts each defined key. Empty values are
    # skipped so we don't substitute the empty string.
    sed_args=""
    for var in LAUNCHDARKLY_SDK_KEY LAUNCHDARKLY_MOBILE_KEY LAUNCHDARKLY_CLIENT_SIDE_ID; do
        eval "val=\${$var-}"
        if [ -n "$val" ]; then
            sed_args="$sed_args -e s|$val|<REDACTED_$var>|g"
        fi
    done
    if [ -n "$sed_args" ]; then
        # shellcheck disable=SC2086
        sed $sed_args "$log" >&2
    else
        cat "$log" >&2
    fi
}

# fail_with_log <log-file> <message>
# Convenience: print the failure message and the redacted log, then exit 1.
fail_with_log() {
    log=$1
    msg=$2
    echo "validator: $msg" >&2
    echo "--- snippet output (keys redacted) ---" >&2
    dump_redacted "$log"
    exit 1
}
