---
id: dotnet-client-sdk/sdk-docs/features/config/index-v3-0
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: SDK configuration example for .NET (client-side).
---

```csharp
var config = Configuration
    .Builder("example-mobile-key")
    .Events(Components.SendEvents().FlushInterval(TimeSpan.FromSeconds(2)))
    .Build();
LdClient client = LdClient.Init(config, context);
```
