FROM mcr.microsoft.com/dotnet/sdk:8.0

# python3 is used by the harness's USING_LIFT_MARKER / TYPE_LIFT_TARGET
# pre-steps (see run.sh) to hoist `using …;` directives and lift type
# declarations out of doc-fragment method bodies before the C# compiler
# runs.
RUN apt-get update && apt-get install -y --no-install-recommends \
        python3 \
    && rm -rf /var/lib/apt/lists/*

# Pre-bake a warm project so batch per-validate cycles only recompile the
# swapped Program.cs incrementally instead of restoring NuGet packages from
# scratch each time. The pre-restored package set is the superset of every
# snippet's requirements (ServerSdk + Observability + ServerSdk.Ai +
# ServerSdk.Redis + ServerSdk.Telemetry + ServerSdk.Consul +
# ServerSdk.DynamoDB); extra references are harmless for
# snippets that use fewer, so the per-snippet requirements.txt is a no-op
# against this baseline. The init scaffold stages its own
# Microsoft.NET.Sdk.Web .csproj (which carries no package references), so for
# that snippet the harness adds its requirements and restores.
WORKDIR /opt/hello-dotnet
RUN printf '%s\n' \
    '<Project Sdk="Microsoft.NET.Sdk">' \
    '  <PropertyGroup>' \
    '    <OutputType>Exe</OutputType>' \
    '    <TargetFramework>net8.0</TargetFramework>' \
    '    <Nullable>disable</Nullable>' \
    '    <RootNamespace>HelloDotNet</RootNamespace>' \
    '    <AssemblyName>HelloDotNet</AssemblyName>' \
    '  </PropertyGroup>' \
    '</Project>' > HelloDotNet.csproj \
    && printf 'class Program { static void Main() { System.Console.WriteLine("placeholder"); } }\n' > Program.cs \
    && dotnet add package LaunchDarkly.ServerSdk \
    && dotnet add package LaunchDarkly.Observability \
    && dotnet add package LaunchDarkly.ServerSdk.Ai \
    && dotnet add package LaunchDarkly.ServerSdk.Redis \
    && dotnet add package LaunchDarkly.ServerSdk.Telemetry \
    && dotnet add package LaunchDarkly.ServerSdk.Consul \
    && dotnet add package LaunchDarkly.ServerSdk.DynamoDB \
    && dotnet add package OpenFeature \
    && dotnet add package LaunchDarkly.OpenFeature.ServerProvider \
    && dotnet build -c Debug \
    && cp HelloDotNet.csproj /opt/baseline.csproj

WORKDIR /work

COPY shared /harness-shared
COPY languages/dotnet-server/harness /harness

ENTRYPOINT ["/harness/run.sh"]
