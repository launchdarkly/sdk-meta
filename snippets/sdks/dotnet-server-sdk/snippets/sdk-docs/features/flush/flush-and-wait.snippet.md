---
id: dotnet-server-sdk/sdk-docs/features/flush/flush-and-wait
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Synchronous event flush example for .NET (server-side).
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only

---

```csharp
client.FlushAndWait(TimeSpan.FromSeconds(2));
```
