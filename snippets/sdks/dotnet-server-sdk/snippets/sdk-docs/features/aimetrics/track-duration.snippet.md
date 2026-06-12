---
id: dotnet-server-sdk/sdk-docs/features/aimetrics/track-duration
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Track duration manually for the .NET AI SDK.
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only

---

```csharp
/// Track your own start and stop time.
/// Set duration to the time (in ms) that your AI model generation takes.
/// The duration may include network latency, depending on how you calculate it.

tracker.TrackDuration(response.Metrics.LatencyMs);
```
