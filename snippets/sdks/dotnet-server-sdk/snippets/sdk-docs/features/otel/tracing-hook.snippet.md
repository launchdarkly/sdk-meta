---
id: dotnet-server-sdk/sdk-docs/features/otel/tracing-hook
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: OpenTelemetry tracing hook configuration for the .NET (server-side) SDK.
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only
---

```csharp
using LaunchDarkly.Sdk.Server.Hooks;
using LaunchDarkly.Sdk.Server.Telemetry;

var config = Configuration.Builder("YOUR_SDK_KEY")
  .Hooks(Components.Hooks()
    .Add(TracingHook.Default())
  ).Build();

var client = new LdClient(config);
```
