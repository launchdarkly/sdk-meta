---
id: cloudflare-server-sdk/sdk-docs/features/migrations/stage-switch
sdk: cloudflare-server-sdk
kind: reference
lang: typescript
description: Per-stage migration structure for Cloudflare SDK v2.2.2+.
validation:
  scaffold: cloudflare-server-sdk/scaffolds/cloudflare-syntax-only
---

```typescript
import { LDMigrationStage } from '@launchdarkly/cloudflare-server-sdk';

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
