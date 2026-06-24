---
id: dotnet-server-sdk/sdk-docs/features/offlinemode/offline-mode
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Offline mode example for .NET (server-side).
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only

---

```csharp
var config = Configuration.Builder("YOUR_SDK_KEY")
    .Offline(true)
    .Build();
LdClient client = new LdClient(config);
```
