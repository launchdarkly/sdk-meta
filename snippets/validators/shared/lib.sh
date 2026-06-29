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

# run_batch <validate-fn>
# Drives batch mode: loops over the manifest at $SNIPPET_BATCH (TSV lines
# of `<relpath>\t<label>`, where relpath is the path under /snippet — or
# $SNIPPET_DIR for native harnesses — to the snippet's entry file) and
# invokes <validate-fn> <relpath> <label> for each, inside one warm
# workspace the harness has already set up. The callback returns 0 on
# success, non-zero on failure.
#
# run_batch tallies results, prints `ok:`/`FAIL:` per snippet and a final
# `batch: <passed>/<total> passed`, and returns non-zero if any snippet
# failed (the run continues past failures so one bad fragment doesn't hide
# the rest). The callback is invoked in an `if`, so POSIX shells suspend
# `set -e` for its duration — callbacks must check command exit codes
# explicitly rather than relying on errexit.
run_batch() {
    _vb_fn=$1
    require_env SNIPPET_BATCH
    _vb_tab=$(printf '\t')
    _vb_fail=0
    _vb_total=0
    _vb_pass=0
    while IFS="$_vb_tab" read -r _vb_rel _vb_label; do
        [ -n "$_vb_rel" ] || continue
        _vb_total=$((_vb_total + 1))
        if "$_vb_fn" "$_vb_rel" "$_vb_label"; then
            _vb_pass=$((_vb_pass + 1))
            echo "ok: $_vb_label"
        else
            echo "FAIL: $_vb_label" >&2
            _vb_fail=1
        fi
    done < "$SNIPPET_BATCH"
    echo "batch: $_vb_pass/$_vb_total passed"
    return "$_vb_fail"
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
    # The process can print the success line and exit between our last grep
    # and the liveness check — common for syntax-only hellos that print and
    # return immediately, where the write may not have flushed when we
    # grepped. Do one final grep before giving up so that race doesn't read
    # as a failure.
    if grep -E "feature flag evaluates to [Tt]rue" "$log" >/dev/null 2>&1; then
        wait "$pid" 2>/dev/null || true
        grep -E "feature flag evaluates to [Tt]rue" "$log" | head -1
        echo "validator: ok"
        return 0
    fi
    return 1
}

# dump_redacted <log-file>
# Prints the log to stderr with LAUNCHDARKLY_SDK_KEY (and optionally
# LAUNCHDARKLY_MOBILE_KEY / LAUNCHDARKLY_CLIENT_SIDE_ID) replaced with a
# placeholder. Defense in depth — today's snippets never echo a key, but a
# future snippet could and this log gets piped into CI output.
# Escapes a string for use as a sed substitution pattern with `|` as the
# delimiter. Backslash-escapes the BRE metacharacters and the chosen
# delimiter so a credential containing any of them can't break out of
# the pattern half of the s|…|…| expression.
_sed_escape_pattern() {
    printf '%s' "$1" | sed 's/[][\\.*^$|/&]/\\&/g'
}

dump_redacted() {
    log=$1
    # Build a sed program that redacts each defined key. Empty values are
    # skipped so we don't substitute the empty string.
    sed_args=""
    for var in LAUNCHDARKLY_SDK_KEY LAUNCHDARKLY_MOBILE_KEY LAUNCHDARKLY_CLIENT_SIDE_ID; do
        eval "val=\${$var-}"
        if [ -n "$val" ]; then
            esc=$(_sed_escape_pattern "$val")
            sed_args="$sed_args -e s|$esc|<REDACTED_$var>|g"
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
