---
id: go-server-sdk/sdk-docs/features/aimetrics/track-duration
sdk: go-server-sdk
kind: reference
lang: go
description: Track duration manually for the Go AI SDK.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only

---

```go
// Track your own start and stop time.

// Set duration to the time that your AI model generation takes.
// The duration may include network latency, depending on how you calculate it.

tracker.TrackDuration(10 * time.Millisecond);
```
