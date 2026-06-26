---
id: go-server-sdk/sdk-docs/features/aimetrics/track-feedback
sdk: go-server-sdk
kind: reference
lang: go
description: Track output satisfaction rate for the Go AI SDK.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only

---

```go
// Track your own output satisfaction rate.

// Pass in feedbackPositive or feedbackNegative.
tracker.TrackFeedback(ldai.FeedbackPositive);
```
