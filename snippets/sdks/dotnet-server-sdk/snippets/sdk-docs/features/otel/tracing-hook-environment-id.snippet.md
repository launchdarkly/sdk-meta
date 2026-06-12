---
id: dotnet-server-sdk/sdk-docs/features/otel/tracing-hook-environment-id
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: OpenTelemetry tracing hook with an explicit environment ID for the .NET (server-side) SDK.
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only
---

```csharp
using LaunchDarkly.Sdk.Server.Hooks; // v1.1 or later
using LaunchDarkly.Sdk.Server.Telemetry;

var config = Configuration.Builder("YOUR_SDK_KEY")
  .Hooks(Components.Hooks()
    .Add(TracingHook.Builder()
      .EnvironmentId("example-client-side-id")
      .Build()
    )
  ).Build();

var client = new LdClient(config);
```
