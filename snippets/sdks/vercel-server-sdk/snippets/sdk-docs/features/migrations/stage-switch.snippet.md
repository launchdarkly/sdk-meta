---
id: vercel-server-sdk/sdk-docs/features/migrations/stage-switch
sdk: vercel-server-sdk
kind: reference
lang: typescript
description: Per-stage migration structure for Vercel SDK v1.1.6+.
validation:
  scaffold: vercel-server-sdk/scaffolds/vercel-syntax-only
---

```typescript
import { LDMigrationStage } from '@launchdarkly/vercel-server-sdk';

// define the combination of reads and writes from the new and old systems
// that should occur at each migration stage

switch (value) {
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
