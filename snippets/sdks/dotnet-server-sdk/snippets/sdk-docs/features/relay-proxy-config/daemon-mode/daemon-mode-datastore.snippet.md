---
id: dotnet-server-sdk/sdk-docs/features/relay-proxy-config/daemon-mode/daemon-mode-datastore
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Daemon mode configuration example for .NET (server-side).
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only
---

```csharp
var config = Configuration.Builder("YOUR_SDK_KEY")
    .DataStore(
        Components.PersistentDataStore(
            SomeDatabaseName.DataStore()
        )
    )
    .DataSource(Components.ExternalUpdatesOnly)
    .Build();
```
