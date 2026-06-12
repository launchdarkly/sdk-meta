---
id: dotnet-server-sdk/sdk-docs/features/aimetrics/track-tokens
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Track token usage manually for the .NET AI SDK.
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only

---

```csharp
/// Track your own token usage.

tracker.TrackTokens(response.Usage);
```
