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
#   gem install …      — route gems to a per-run GEM_HOME (so we don't
#                        need root), then assert by walking that dir.
#   cargo add …        — `cargo init`, then run body, then grep Cargo.toml.
#   php composer.phar
#     require …        — symlink /opt/composer.phar into $WORK, init a
#                        minimal composer project, run body, check vendor/.
#   composer require … — same flow as `php composer.phar require`, with a
#                        thin `composer` wrapper staged on PATH so the
#                        bare invocation resolves.
#   dotnet add
#     package …        — `dotnet new console`, then run body, then grep
#                        the generated .csproj for <PackageReference>.
#   Install-Package …  — legacy NuGet PowerShell. Stripped from multi-line
#                        bodies (we keep validating the dotnet CLI variant
#                        in the same body); causes a hard error if it's
#                        the only line.
#   implementation
#     group: …         — Gradle DSL fragment. Wrapped in a minimal
#                        build.gradle's dependencies block; `gradle
#                        dependencies` resolves and we grep the output
#                        for the artifact name.
#
# Any unrecognized leading token is a hard error: a new install style means
# the harness needs to learn it, not silently no-op.
set -eu

. /harness-shared/lib.sh
require_env SNIPPET_ENTRYPOINT

WORK=$(mktemp -d)
# `set -e` kills before assert_node_modules / fail_with_log can dump the LOG
# captured by run_in_log, so a failed install command otherwise produces zero
# diagnostic output. Dump the LOG on non-zero exit so CI artifacts surface the
# real error.
cleanup() {
    rc=$?
    if [ "$rc" -ne 0 ] && [ -n "${LOG:-}" ] && [ -s "$LOG" ]; then
        echo "=== shell-install/run.sh exiting with $rc; LOG dump follows ==="
        cat "$LOG"
    fi
    rm -rf "$WORK"
}
trap cleanup EXIT
cd "$WORK"

BODY="/snippet/$SNIPPET_ENTRYPOINT"
if [ ! -f "$BODY" ]; then
    echo "validator: snippet entrypoint not found: $BODY" >&2
    exit 1
fi

# Body is one (sometimes few) shell command lines. Read them all so a
# multi-line install (e.g. `cd … && npm i …`) still works.
COMMAND=$(cat "$BODY")

# Drop legacy NuGet PowerShell `Install-Package` lines. The cmdlet only
# exists in PowerShell with the NuGet PackageManagement module — not in
# our toolchain. Snippets that show both the legacy `Install-Package`
# and the modern `dotnet add package` forms (for backward-compat
# documentation) get the dotnet variant validated; the Install-Package
# variant is structurally unrunnable on Linux. Bodies whose only line is
# Install-Package fall through to the `*` case below and fail loudly.
if printf '%s' "$COMMAND" | grep -q '^[[:space:]]*Install-Package'; then
    COMMAND=$(printf '%s' "$COMMAND" | sed '/^[[:space:]]*Install-Package/d')
fi

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
#
# Multi-line bodies (e.g. observability snippets that install the SDK and
# the o11y plugin in two separate commands) are folded by the END block:
# `last` accumulates across every non-empty line and only the final value
# is printed, so the assertion targets the last package the body installs.
last_pkg() {
    printf '%s' "$1" | awk '
        {
            # Strip inline shell comments so trailing `# foo bar` annotations
            # in multi-line install bodies do not get picked up as package names.
            sub(/[ \t]*#.*$/, "", $0);
            $0 = $0;  # force re-split into $1..$NF after the substitution
        }
        NF >= 3 {
            for (i = 3; i <= NF; i++) {
                if ($i ~ /^-/) continue;
                last = $i;
            }
        }
        END { print last }
    '
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
    gem)
        if [ "$SUB" != "install" ]; then
            fail_with_log "$LOG" "unrecognized gem subcommand for install snippet: $SUB"
        fi
        # Route gems to a clean per-run dir via GEM_HOME so we don't need
        # root and concurrent runs don't collide. PATH gets the bin dir
        # too in case any post-install asserts shell out.
        export GEM_HOME="$WORK/.gems"
        export GEM_PATH="$GEM_HOME"
        export PATH="$GEM_HOME/bin:$PATH"
        run_in_log
        pkg=$(last_pkg "$COMMAND")
        # Walk GEM_HOME/gems for a directory named `<pkg>-<version>`. The
        # `gem list -i` route is finicky with custom GEM_HOME paths; a
        # filesystem check is unambiguous.
        if [ -d "$GEM_HOME/gems" ] && ls "$GEM_HOME/gems" | grep -E "^${pkg}-[0-9]" >/dev/null 2>&1; then
            echo "validator: ok — $pkg installed in $GEM_HOME"
        else
            fail_with_log "$LOG" "expected gem $pkg to be installed under $GEM_HOME/gems"
        fi
        ;;
    cargo)
        if [ "$SUB" != "add" ]; then
            fail_with_log "$LOG" "unrecognized cargo subcommand for install snippet: $SUB"
        fi
        cargo init --quiet --name install-sanity --vcs none . >/dev/null 2>&1
        run_in_log
        pkg=$(last_pkg "$COMMAND")
        if grep -q "$pkg" Cargo.toml; then
            echo "validator: ok — $pkg present in Cargo.toml"
        else
            fail_with_log "$LOG" "expected $pkg in Cargo.toml after cargo add"
        fi
        ;;
    php)
        # Body shape: `php composer.phar require …`. Stage composer.phar
        # alongside the body so the literal reference resolves, then
        # initialize a minimal composer project so `require` has somewhere
        # to write.
        if [ "$SUB" != "composer.phar" ]; then
            fail_with_log "$LOG" "unrecognized php install shape: php $SUB"
        fi
        cp /opt/composer.phar ./composer.phar
        php composer.phar init --quiet --name=example/sanity --no-interaction >/dev/null 2>&1
        run_in_log
        pkg=$(last_pkg "$COMMAND")
        if [ -d "vendor/$pkg" ]; then
            echo "validator: ok — $pkg present under vendor/"
        else
            fail_with_log "$LOG" "expected vendor/$pkg to exist after composer require"
        fi
        ;;
    composer)
        # Body shape: `composer require <vendor>/<pkg>`. The image ships
        # composer only as /opt/composer.phar (no global `composer`
        # binary), so stage a thin wrapper on PATH that the body's bare
        # `composer` invocation resolves to, then initialize a minimal
        # composer project so `require` has somewhere to write.
        if [ "$SUB" != "require" ]; then
            fail_with_log "$LOG" "unrecognized composer subcommand for install snippet: $SUB"
        fi
        mkdir -p bin
        printf '#!/bin/sh\nexec php /opt/composer.phar "$@"\n' > bin/composer
        chmod +x bin/composer
        php /opt/composer.phar init --quiet --name=example/sanity --no-interaction >/dev/null 2>&1
        PATH="$WORK/bin:$PATH" run_in_log
        pkg=$(last_pkg "$COMMAND")
        if [ -d "vendor/$pkg" ]; then
            echo "validator: ok — $pkg present under vendor/"
        else
            fail_with_log "$LOG" "expected vendor/$pkg to exist after composer require"
        fi
        ;;
    dotnet)
        if [ "$SUB" != "add" ]; then
            fail_with_log "$LOG" "unrecognized dotnet subcommand for install snippet: $SUB"
        fi
        dotnet new console -n InstallSanity --force >/dev/null 2>&1
        cd InstallSanity
        run_in_log
        pkg=$(last_pkg "$COMMAND")
        if grep -q "<PackageReference Include=\"$pkg\"" InstallSanity.csproj; then
            echo "validator: ok — $pkg in InstallSanity.csproj"
        else
            fail_with_log "$LOG" "expected <PackageReference Include=\"$pkg\"…> in InstallSanity.csproj"
        fi
        ;;
    implementation)
        # Gradle DSL fragment — wrap in a minimal build.gradle's
        # dependencies block and run `gradle dependencies` to confirm
        # the artifact resolves from mavenCentral. We grep the
        # dependency-tree output for the artifact name (the snippet's
        # `name: '…'` field).
        cat > build.gradle <<'EOF'
plugins { id 'java' }
repositories { mavenCentral() }
dependencies {
EOF
        printf '%s\n' "$COMMAND" >> build.gradle
        echo '}' >> build.gradle
        cat > settings.gradle <<'EOF'
rootProject.name = 'install-sanity'
EOF
        gradle --no-daemon --quiet dependencies >"$LOG" 2>&1
        artifact=$(printf '%s' "$COMMAND" | sed -n "s/.*name:[[:space:]]*'\([^']*\)'.*/\1/p")
        if [ -n "$artifact" ] && grep -q "$artifact" "$LOG"; then
            echo "validator: ok — $artifact resolved by gradle"
        else
            fail_with_log "$LOG" "expected gradle dependencies output to mention artifact"
        fi
        ;;
    *)
        fail_with_log "$LOG" "unrecognized install-snippet leading token: $LEAD"
        ;;
esac
