---
id: go-server-sdk/sdk-docs/features/migrations/track-migration-scopedclient
sdk: go-server-sdk
kind: reference
lang: go
description: Migration metrics tracking (TrackMigrationOp) for Go SDK v7.13.4+ (LDScopedClient).
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only
---

```go
event, _ := tracker.Build();

err := scopedClient.TrackMigrationOp(*event);
// LDScopedClient is in beta and may change without notice.
```
