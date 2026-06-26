---
id: cloudflare-server-sdk/sdk-docs/features/migrations/migration-variation
sdk: cloudflare-server-sdk
kind: reference
lang: typescript
description: Migration stage evaluation (migrationVariation) for Cloudflare SDK v2.2.2+.
validation:
  scaffold: cloudflare-server-sdk/scaffolds/cloudflare-syntax-only
---

```typescript
  import {
    LDContext,
    LDMigrationStage,
  } from '@launchdarkly/cloudflare-server-sdk'

  const context: LDContext = {
     kind: 'user',
     key: 'example-user-key',
     name: 'Sandy',
  };

  const { value, tracker } = await client.migrationVariation(
    'example-migration-flag-key',
    context,
    LDMigrationStage.Off
  );
```
