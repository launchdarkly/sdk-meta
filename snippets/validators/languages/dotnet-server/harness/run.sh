#!/bin/sh
# Batch validator for .NET server SDK snippets. The Dockerfile pre-baked a
# warm project at /opt/hello-dotnet with the SDK package superset restored
# and built, so per-snippet we reset to that baseline, swap in the snippet's
# Program.cs (and its own .csproj, if it staged one), and rebuild
# incrementally instead of restoring NuGet from scratch each time.
#
# The Go runner stages every matching snippet under /snippet and writes
# /snippet/manifest.tsv; run_batch loops over them in the single warm
# project (one container instead of one per snippet).
set -eu

. /harness-shared/lib.sh
require_env LAUNCHDARKLY_SDK_KEY LAUNCHDARKLY_FLAG_KEY SNIPPET_BATCH

PROJ=/opt/hello-dotnet
cd "$PROJ"

# using_lift / type_lift: the same source rewrites the old per-snippet
# harness applied, factored into functions invoked per snippet.
using_lift() {
    grep -q 'USING_LIFT_MARKER' "$1" || return 0
    python3 - "$1" <<'PYEOF'
import sys, re
fp = sys.argv[1]
with open(fp) as f:
    lines = f.read().splitlines()
marker_idx = next((i for i, l in enumerate(lines) if 'USING_LIFT_MARKER' in l), None)
if marker_idx is None:
    sys.exit(0)
lift = []
for i in range(marker_idx + 1, len(lines)):
    s = lines[i].lstrip()
    m = re.match(r'using\s+([A-Za-z_][A-Za-z0-9_.]*)\s*;\s*(//.*)?$', s)
    if m:
        if s not in lift:
            lift.append(s.rstrip())
        lines[i] = ''
if lift:
    lines = lines[:marker_idx] + lift + lines[marker_idx:]
with open(fp, 'w') as f:
    f.write('\n'.join(lines) + '\n')
PYEOF
}

type_lift() {
    grep -q 'TYPE_LIFT_TARGET' "$1" || return 0
    python3 - "$1" <<'PYEOF'
import re, sys
fp = sys.argv[1]
with open(fp) as f:
    lines = f.read().splitlines()
target_idx = next((i for i, l in enumerate(lines) if 'TYPE_LIFT_TARGET' in l), None)
begin_idx = next((i for i, l in enumerate(lines) if 'BODY_BEGIN' in l), None)
end_idx = next((i for i, l in enumerate(lines) if 'BODY_END' in l), None)
if target_idx is None or begin_idx is None or end_idx is None:
    sys.exit(0)
decl_re = re.compile(
    r'^\s*((public|private|protected|internal|static|sealed|abstract|partial)\s+)*'
    r'(class|interface|enum|struct|record)\s+\w')
lifted = []
i = begin_idx + 1
depth = 0
while i < end_idx:
    line = lines[i]
    if depth == 0 and decl_re.match(line):
        block = []
        bdepth = 0
        seen_brace = False
        while i < end_idx:
            block.append(lines[i])
            bdepth += lines[i].count('{') - lines[i].count('}')
            if '{' in lines[i] or '}' in lines[i]:
                seen_brace = True
            lines[i] = ''
            i += 1
            if seen_brace and bdepth == 0:
                break
        lifted.extend(block)
        continue
    depth += line.count('{') - line.count('}')
    i += 1
if lifted:
    lines = lines[:target_idx + 1] + lifted + lines[target_idx + 1:]
with open(fp, 'w') as f:
    f.write('\n'.join(lines) + '\n')
PYEOF
}

validate_one() {
    relpath=$1
    idx=$(printf '%s' "$relpath" | cut -d/ -f1)
    unit="/snippet/$idx"

    # Reset to the warm baseline: restore the baseline .csproj and drop any
    # .csproj a previous snippet staged, so a snippet that needs the Web SDK
    # doesn't leak into the next one.
    cp /opt/baseline.csproj HelloDotNet.csproj
    find . -maxdepth 1 -name '*.csproj' ! -name HelloDotNet.csproj -delete 2>/dev/null || true

    # Stage the snippet's Program.cs (the entrypoint) and, if the scaffold
    # staged its own .csproj (e.g. the ASP.NET Core init), use it instead of
    # the baseline.
    cp "$unit/$(basename "$relpath")" Program.cs
    staged_csproj=""
    for cs in "$unit"/*.csproj; do
        [ -f "$cs" ] || continue
        rm -f HelloDotNet.csproj
        cp "$cs" .
        staged_csproj=$(basename "$cs")
    done

    using_lift Program.cs
    type_lift Program.cs

    LOG=$(mktemp)
    # A snippet that staged its own .csproj (the ASP.NET Core init) starts
    # from a project with no package references, so add the packages its
    # requirements.txt lists and restore. Baseline snippets reuse the warm
    # pre-restored superset and skip this entirely.
    if [ -n "$staged_csproj" ]; then
        if [ -f "$unit/requirements.txt" ]; then
            while IFS= read -r pkg; do
                [ -z "$pkg" ] && continue
                dotnet add package "$pkg" --no-restore >>"$LOG" 2>&1 || true
            done < "$unit/requirements.txt"
        fi
        dotnet restore --verbosity quiet >>"$LOG" 2>&1 || true
    fi

    CI=1 timeout --signal=TERM 180s dotnet run --project . --verbosity quiet >"$LOG" 2>&1 &
    PID=$!
    deadline=$(( $(date +%s) + 170 ))
    if await_success_line "$LOG" "$PID" "$deadline"; then
        rm -f "$LOG"
        return 0
    fi
    kill -TERM "$PID" 2>/dev/null || true
    wait "$PID" 2>/dev/null || true
    dump_redacted "$LOG" >&2
    rm -f "$LOG"
    return 1
}

run_batch validate_one
