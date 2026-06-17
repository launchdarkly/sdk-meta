---
id: dotnet-server-sdk/sdk-docs/features/flush/flush-interval
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Flush interval configuration example for .NET (server-side).
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only

---

```csharp
var config = Configuration.Builder("YOUR_SDK_KEY")
    .Events(
        Components.SendEvents().FlushInterval(TimeSpan.FromSeconds(10))
    )
    .Build();
var client = new LdClient(config);

```
