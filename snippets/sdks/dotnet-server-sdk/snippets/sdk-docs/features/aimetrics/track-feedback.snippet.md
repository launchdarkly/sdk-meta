---
id: dotnet-server-sdk/sdk-docs/features/aimetrics/track-feedback
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Track output satisfaction rate for the .NET AI SDK.
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only

---

```csharp
/// Track your own output satisfaction rate.

/// Pass in Feedback.Positive or Feedback.Negative.
tracker.TrackFeedback(Feedback.Positive);
```
