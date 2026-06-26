---
id: dotnet-server-sdk/sdk-docs/features/datasaving/standard-setup
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Data saving mode standard setup for .NET (server-side).
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only
---

```csharp
using LaunchDarkly.Sdk.Server;

var config = Configuration.Builder("YOUR_SDK_KEY")
    .DataSystem(Components.DataSystem().Default())
    .Build();

var client = new LdClient(config);
```
