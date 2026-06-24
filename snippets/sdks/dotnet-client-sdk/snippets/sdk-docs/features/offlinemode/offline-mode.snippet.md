---
id: dotnet-client-sdk/sdk-docs/features/offlinemode/offline-mode
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: Offline mode example for .NET (client-side).
validation:
  scaffold: dotnet-client-sdk/scaffolds/csharp-client-syntax-only

---

```csharp
var config = Configuration
    .Builder("example-mobile-key", ConfigurationBuilder.AutoEnvAttributes.Enabled)
    .Offline(true)
    .Build();
LdClient client = LdClient.Init(config, context, TimeSpan.FromSeconds(10));
```
