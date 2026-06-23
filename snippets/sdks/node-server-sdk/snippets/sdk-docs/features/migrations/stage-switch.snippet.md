---
id: node-server-sdk/sdk-docs/features/migrations/stage-switch
sdk: node-server-sdk
kind: reference
lang: typescript
description: Per-stage migration structure for Node.js (server-side) SDK v9.
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only
---

```typescript
// define the combination of reads and writes from the new and old systems
// that should occur at each migration stage

switch (stage) {
  case LDMigrationStage.Off: { }
  case LDMigrationStage.DualWrite: { }
  case LDMigrationStage.Shadow: { }
  case LDMigrationStage.Live: { }
  case LDMigrationStage.RampDown: { }
  case LDMigrationStage.Complete: { }
  default: {
    // throw an error
  }
}
```
