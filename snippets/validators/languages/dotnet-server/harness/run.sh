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

# .NET wants a project file; gonfalon's flow uses Visual Studio's "new
# console app" wizard which creates one. We synthesize the minimum.
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
