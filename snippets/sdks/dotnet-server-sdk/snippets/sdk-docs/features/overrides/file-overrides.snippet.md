---
id: dotnet-server-sdk/sdk-docs/features/overrides/file-overrides
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: File-based flag override source configuration example for .NET (server-side) SDK v9.0+.
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only
---

```csharp
using LaunchDarkly.Sdk.Server;
using LaunchDarkly.Sdk.Server.Integrations;

var config = Configuration.Builder("YOUR_SDK_KEY")
    .DataSystem(
        Components.DataSystem().Default()
            .Overrides(FileOverrides.Source().FilePaths("/etc/launchdarkly/overrides.json"))
    )
    .Build();

var client = new LdClient(config);
```
