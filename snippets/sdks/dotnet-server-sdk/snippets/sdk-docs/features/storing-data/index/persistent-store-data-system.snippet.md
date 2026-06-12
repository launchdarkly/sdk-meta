---
id: dotnet-server-sdk/sdk-docs/features/storing-data/index/persistent-store-data-system
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Persistent store configuration via the DataSystem builder for .NET (server-side) SDK 8.11+.
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only

---

```csharp
var config = Configuration.Builder("YOUR_SDK_KEY")
    .DataSystem(
        Components.DataSystem().PersistentStore(
            Components.PersistentDataStore(Redis.DataStore())
        )
    )
    .Build();
```
