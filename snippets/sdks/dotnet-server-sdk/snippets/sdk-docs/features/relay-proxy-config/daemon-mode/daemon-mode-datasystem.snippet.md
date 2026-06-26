---
id: dotnet-server-sdk/sdk-docs/features/relay-proxy-config/daemon-mode/daemon-mode-datasystem
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Daemon mode DataSystem configuration example for .NET (server-side).
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only
---

```csharp
var config = Configuration.Builder("YOUR_SDK_KEY")
    .DataSystem(
        Components.DataSystem().Daemon(
            Components.PersistentDataStore(Redis.DataStore())
        )
    )
    .Build();
```
