---
id: go-server-sdk/sdk-docs/features/migrations/stage-switch
sdk: go-server-sdk
kind: reference
lang: go
description: Per-stage migration structure for Go SDK v7+.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only
---

```go
// define the combination of reads and writes from the new and old systems
// that should occur at each migration stage

switch stage {
  case ldmigration.Off:
  case ldmigration.DualWrite:
  case ldmigration.Shadow:
  case ldmigration.Live:
  case ldmigration.RampDown:
  case ldmigration.Complete:
  default: {
    // throw an error
  }
}
```
