---
id: dotnet-server-sdk/sdk-docs/features/aimetrics/track-time-to-first-token
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Track time to first token for the .NET AI SDK.
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only

---

```csharp
/// Track the time it takes to generate the first token
/// Pass in the time (in ms) until your first token is generated
/// This may include network latency, depending on how you calculate it
tracker.TrackTimeToFirstToken(1000);
```
