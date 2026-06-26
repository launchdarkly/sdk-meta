---
id: vercel-server-sdk/sdk-docs/features/migrations/migration-variation
sdk: vercel-server-sdk
kind: reference
lang: typescript
description: Migration stage evaluation (migrationVariation) for Vercel SDK v1.1.6+.
validation:
  scaffold: vercel-server-sdk/scaffolds/vercel-syntax-only
---

```typescript
  import {
    LDContext,
    LDMigrationStage,
  } from '@launchdarkly/vercel-server-sdk';

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
