---
id: dotnet-server-sdk/sdk-docs/features/storing-data/index/persistent-store
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Persistent feature store configuration example for .NET (server-side).
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only

---

```csharp
using LaunchDarkly.Sdk.Server;
using LaunchDarkly.Sdk.Server.Integrations;

var config = Configuration.Builder(sdkKey)
    .DataStore(
        Components.PersistentDataStore(
            SomeDatabaseName.DataStore()
        )
    )
    .Build();
```
