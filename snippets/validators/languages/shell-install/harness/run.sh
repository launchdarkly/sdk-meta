#!/bin/sh
# Runs an sdk-info install-snippet body in a clean working dir and asserts
# the package manager actually fetched the package.
#
# Inputs (env): SNIPPET_ENTRYPOINT — path under /snippet to the install
#               command body (one shell line, by convention).
#
# The harness sniffs the body's leading token to pick a pre-state and a
# post-check:
#
#   npm i …            — `npm init -y`, then run body, then assert package
#                        appears under node_modules.
#   pnpm i …           — same pre-state; pnpm uses node_modules too.
#   yarn add …         — same pre-state.
#   pip3 install …     — create a venv, run inside it, assert with `pip show`.
#   pip install …      — same as pip3.
#   go get …           — `go mod init`, then run body, then grep go.mod.
#   bower install …    — install bower locally, then run body, then assert
#                        bower_components dir exists.
#
# Any unrecognized leading token is a hard error: a new install style means
# the harness needs to learn it, not silently no-op.
set -eu

. /harness-shared/lib.sh
require_env SNIPPET_ENTRYPOINT

WORK=$(mktemp -d)
trap 'rm -rf "$WORK"' EXIT
cd "$WORK"

BODY="/snippet/$SNIPPET_ENTRYPOINT"
if [ ! -f "$BODY" ]; then
    echo "validator: snippet entrypoint not found: $BODY" >&2
    exit 1
fi

# Body is one (sometimes few) shell command lines. Read them all so a
# multi-line install (e.g. `cd … && npm i …`) still works.
COMMAND=$(cat "$BODY")
# Sniff the leading non-whitespace token to pick a strategy.
LEAD=$(printf '%s' "$COMMAND" | awk 'NF{print $1; exit}')
SUB=$(printf '%s' "$COMMAND" | awk 'NF{print $2; exit}')

LOG=$(mktemp)

run_in_log() {
    # shellcheck disable=SC2086 # COMMAND is intentionally split.
    sh -c "$COMMAND" >"$LOG" 2>&1
}

# Extract the package name from a `<tool> <verb> <pkg>` install line. Each
# of npm/pnpm/yarn accepts multiple targets; we assert on the LAST one
# (which is where every sdk-info install command puts the package). The
# extraction is intentionally simple: discard the first two tokens and the
# pre-existing flag tokens, return the remaining last token.
last_pkg() {
    printf '%s' "$1" | awk '{
        for (i = 3; i <= NF; i++) {
            if ($i ~ /^-/) continue;
            last = $i;
        }
        print last;
    }'
}

assert_node_modules() {
    pkg=$(last_pkg "$COMMAND")
    if [ -z "$pkg" ]; then
        fail_with_log "$LOG" "could not extract package name from: $COMMAND"
    fi
    if [ -d "node_modules/$pkg" ]; then
        echo "validator: ok — $pkg present under node_modules"
        return 0
    fi
    fail_with_log "$LOG" "expected node_modules/$pkg to exist after install"
}

case "$LEAD" in
    npm)
        if [ "$SUB" != "i" ] && [ "$SUB" != "install" ] && [ "$SUB" != "add" ]; then
            fail_with_log "$LOG" "unrecognized npm subcommand for install snippet: $SUB"
        fi
        npm init -y >/dev/null
        run_in_log
        assert_node_modules
        ;;
    pnpm)
        if [ "$SUB" != "i" ] && [ "$SUB" != "install" ] && [ "$SUB" != "add" ]; then
            fail_with_log "$LOG" "unrecognized pnpm subcommand for install snippet: $SUB"
        fi
        npm init -y >/dev/null
        run_in_log
        assert_node_modules
        ;;
    yarn)
        if [ "$SUB" != "add" ] && [ "$SUB" != "install" ]; then
            fail_with_log "$LOG" "unrecognized yarn subcommand for install snippet: $SUB"
        fi
        npm init -y >/dev/null
        run_in_log
        assert_node_modules
        ;;
    pip|pip3)
        python3 -m venv .venv >/dev/null
        # Re-point the install command at the venv's pip so we don't pollute
        # the system python (and so the venv's pip wins on $PATH inside the
        # subshell).
        PATH="$WORK/.venv/bin:$PATH" run_in_log
        pkg=$(last_pkg "$COMMAND")
        if "$WORK/.venv/bin/pip" show "$pkg" >/dev/null 2>&1; then
            echo "validator: ok — $pkg installed in venv"
        else
            fail_with_log "$LOG" "expected venv pip to have $pkg installed"
        fi
        ;;
    go)
        if [ "$SUB" != "get" ]; then
            fail_with_log "$LOG" "only `go get` is supported by this harness, got: go $SUB"
        fi
        go mod init example/install-sanity >/dev/null 2>&1
        run_in_log
        # `go get` writes the require line into go.mod; grep for the module
        # path (last token of the body, stripped of any version suffix).
        target=$(last_pkg "$COMMAND")
        # `go get pkg@v1.2.3` — strip the `@version` for the grep.
        modpath=$(printf '%s' "$target" | awk -F'@' '{print $1}')
        if grep -q "$modpath" go.mod; then
            echo "validator: ok — $modpath present in go.mod"
        else
            fail_with_log "$LOG" "expected $modpath in go.mod after install"
        fi
        ;;
    bower)
        # Install bower locally so the snippet body's `bower install …` finds
        # the binary on PATH without us editing the body.
        npm init -y >/dev/null
        npm install --silent --no-audit --no-fund --no-progress bower >/dev/null
        PATH="$WORK/node_modules/.bin:$PATH" run_in_log
        if [ -d bower_components ]; then
            echo "validator: ok — bower_components present"
        else
            fail_with_log "$LOG" "expected bower_components/ to exist after install"
        fi
        ;;
    *)
        fail_with_log "$LOG" "unrecognized install-snippet leading token: $LEAD"
        ;;
esac
