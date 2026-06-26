---
id: dotnet-server-sdk/sdk-docs/features/datasaving/file-bootstrap
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Data saving mode with file-based bootstrap and live updates for .NET (server-side).
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only
---

```csharp
using LaunchDarkly.Sdk.Server;
using LaunchDarkly.Sdk.Server.Integrations;

var config = Configuration.Builder("YOUR_SDK_KEY")
    .DataSystem(
        Components.DataSystem().Custom()
            .Initializers(
                FileData.DataSource().FilePaths("flags.json"),
                DataSystemComponents.Polling()
            )
            .Synchronizers(
                DataSystemComponents.Streaming(),
                DataSystemComponents.Polling()
            )
    )
    .Build();

var client = new LdClient(config);
```
