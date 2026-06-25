---
id: go-server-sdk/sdk-docs/features/migrations/track-migration
sdk: go-server-sdk
kind: reference
lang: go
description: Migration metrics tracking (TrackMigrationOp) for Go SDK v7+ (LDClient).
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only
---

```go
event, _ := tracker.Build();

err := client.TrackMigrationOp(*event);
```
