---
id: go-server-sdk/sdk-docs/features/aimetrics/get-summary
sdk: go-server-sdk
kind: reference
lang: go
description: Retrieve automatically recorded metrics with GetSummary for the Go AI SDK.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only

---

```go
summary := tracker.GetSummary();

// recorded metrics available in summary.Duration, summary.Feedback,
// summary.Tokens, summary.Success, summary.TimeToFirstToken

```
