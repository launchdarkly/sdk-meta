#!/bin/sh
# Runs the staged .NET server snippet against a real LaunchDarkly environment.
# Synthesizes a minimal .csproj around the snippet's Program.cs and pulls
# whatever NuGet packages the snippet's `validation.requirements` lists.
set -eu

. /harness-shared/lib.sh
require_env LAUNCHDARKLY_SDK_KEY LAUNCHDARKLY_FLAG_KEY SNIPPET_ENTRYPOINT

WORK=$(mktemp -d)
trap 'rm -rf "$WORK"' EXIT
cp -r /snippet/. "$WORK/"
cd "$WORK"

# If the staged Program.cs (csharp-syntax-only scaffold) contains the
# USING_LIFT_MARKER, lift any `using …;` line that's between the
# marker and the next class declaration up to the marker — C# wants
# `using` at file or namespace scope, not inside method bodies. This
# is the same idea as the JVM harness's IMPORT_LIFT_MARKER pre-step.
if [ -f "$SNIPPET_ENTRYPOINT" ] && grep -q 'USING_LIFT_MARKER' "$SNIPPET_ENTRYPOINT"; then
    python3 - "$SNIPPET_ENTRYPOINT" <<'PYEOF'
import sys
fp = sys.argv[1]
with open(fp) as f:
    lines = f.read().splitlines()
marker_idx = next((i for i, l in enumerate(lines) if 'USING_LIFT_MARKER' in l), None)
if marker_idx is None:
    sys.exit(0)
# Walk lines BELOW the marker. Collect any `using …;` line into a
# block to splice ABOVE the marker, replace each lifted line with a
# blank to preserve numbering. `using (var x = …)` C# statement form
# is excluded by requiring a `;` immediately after the namespace path
# (no `(`). A trailing `// ...` line comment (docs sometimes annotate a
# directive with the minimum package version) lifts along with the
# directive.
import re
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
fi

# .NET wants a project file; gonfalon's flow uses Visual Studio's "new
# console app" wizard which creates one. We synthesize the minimum unless
# the snippet's scaffold staged its own `.csproj` (e.g. an ASP.NET Core
# init that needs `Microsoft.NET.Sdk.Web`).
if ! ls *.csproj >/dev/null 2>&1; then
    cat > HelloDotNet.csproj <<'EOF'
<Project Sdk="Microsoft.NET.Sdk">
  <PropertyGroup>
    <OutputType>Exe</OutputType>
    <TargetFramework>net8.0</TargetFramework>
    <Nullable>disable</Nullable>
    <RootNamespace>HelloDotNet</RootNamespace>
    <AssemblyName>HelloDotNet</AssemblyName>
  </PropertyGroup>
</Project>
EOF
fi

if [ -f requirements.txt ]; then
    while IFS= read -r line; do
        [ -z "$line" ] && continue
        # `dotnet add package` does not accept --verbosity; redirect noise to /dev/null.
        dotnet add package "$line" --no-restore >/dev/null
    done < requirements.txt
fi

dotnet restore --verbosity quiet >/dev/null 2>&1 || true

LOG=$(mktemp)

CI=1 timeout --signal=TERM 180s dotnet run --project . --verbosity quiet >"$LOG" 2>&1 &
PID=$!

deadline=$(( $(date +%s) + 170 ))
if await_success_line "$LOG" "$PID" "$deadline"; then
    exit 0
fi

kill -TERM "$PID" 2>/dev/null || true
wait "$PID" 2>/dev/null || true
fail_with_log "$LOG" "did not see expected line: feature flag evaluates to true"
