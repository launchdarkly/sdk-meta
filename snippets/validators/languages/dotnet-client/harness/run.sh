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
