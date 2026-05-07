#!/bin/sh
# Runs the staged .NET client snippet against a real LaunchDarkly environment.
# The snippet uses top-level statements (no namespace), so the synthesized
# .csproj is the bare minimum a `dotnet new console` would produce.
set -eu

. /harness-shared/lib.sh
require_env LAUNCHDARKLY_MOBILE_KEY LAUNCHDARKLY_FLAG_KEY SNIPPET_ENTRYPOINT

WORK=$(mktemp -d)
trap 'rm -rf "$WORK"' EXIT
cp -r /snippet/. "$WORK/"
cd "$WORK"

# If the staged Program.cs (csharp-client-syntax-only scaffold)
# contains the USING_LIFT_MARKER, lift any `using …;` line below the
# marker up to the marker. C# wants `using` at file/namespace scope.
if [ -f "$SNIPPET_ENTRYPOINT" ] && grep -q 'USING_LIFT_MARKER' "$SNIPPET_ENTRYPOINT"; then
    python3 - "$SNIPPET_ENTRYPOINT" <<'PYEOF'
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
    m = re.match(r'using\s+([A-Za-z_][A-Za-z0-9_.]*)\s*;\s*$', s)
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

cat > HelloDotNetClient.csproj <<'EOF'
<Project Sdk="Microsoft.NET.Sdk">
  <PropertyGroup>
    <OutputType>Exe</OutputType>
    <TargetFramework>net8.0</TargetFramework>
    <Nullable>enable</Nullable>
    <ImplicitUsings>enable</ImplicitUsings>
  </PropertyGroup>
</Project>
EOF

if [ -f requirements.txt ]; then
    while IFS= read -r line; do
        [ -z "$line" ] && continue
        dotnet add package "$line" --no-restore >/dev/null
    done < requirements.txt
fi

dotnet restore --verbosity quiet >/dev/null 2>&1 || true

LOG=$(mktemp)

# .NET client snippet exits naturally after evaluating the flag (calls
# client.Dispose()), so we don't need CI=1 nor a long timeout.
timeout --signal=TERM 90s dotnet run --project . --verbosity quiet >"$LOG" 2>&1 &
PID=$!

deadline=$(( $(date +%s) + 80 ))
if await_success_line "$LOG" "$PID" "$deadline"; then
    exit 0
fi

kill -TERM "$PID" 2>/dev/null || true
wait "$PID" 2>/dev/null || true
fail_with_log "$LOG" "did not see expected line: feature flag evaluates to true"
