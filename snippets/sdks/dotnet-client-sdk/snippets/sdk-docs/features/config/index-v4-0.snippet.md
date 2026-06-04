---
id: dotnet-client-sdk/sdk-docs/features/config/index-v4-0
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: SDK configuration example for .NET (client-side).
validation:
  scaffold: dotnet-client-sdk/scaffolds/csharp-client-syntax-only

---

```csharp
var config = Configuration
    .Builder("example-mobile-key", ConfigurationBuilder.AutoEnvAttributes.Enabled)
    .Events(Components.SendEvents().FlushInterval(TimeSpan.FromSeconds(2)))
    .Build();
LdClient client = LdClient.Init(config, context, TimeSpan.FromSeconds(10));
```
