---
id: akamai-server-edgekv-sdk/sdk-docs/features/migrations/stage-switch
sdk: akamai-server-edgekv-sdk
kind: reference
lang: typescript
description: Per-stage migration structure for Akamai SDK v1.0.9+.
validation:
  scaffold: akamai-server-edgekv-sdk/scaffolds/akamai-syntax-only
---

```typescript
import { LDMigrationStage } from '@launchdarkly/akamai-server-edgekv-sdk';

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
