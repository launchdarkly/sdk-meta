---
id: go-server-sdk/sdk-docs/features/aimetrics/track-time-to-first-token
sdk: go-server-sdk
kind: reference
lang: go
description: Track time to first token for the Go AI SDK.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only

---

```go
// Track the time it takes to generate the first token

// Pass in the time (in ms) until your first token is generated
// This may include network latency, depending on how you calculate it

tracker.TrackTimeToFirstToken(10 * time.Millisecond);
```
