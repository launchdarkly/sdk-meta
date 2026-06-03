---
id: dotnet-server-sdk/sdk-docs/features/config/index
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: SDK configuration example for .NET (server-side).
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only

---

```csharp
var config = Configuration.Builder("YOUR_SDK_KEY")
    .Events(
        Components.SendEvents().FlushInterval(TimeSpan.FromSeconds(2))
    )
    .StartWaitTime(TimeSpan.FromSeconds(5))
    .Build();
var client = new LdClient(config);
```
