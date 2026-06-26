---
id: node-server-sdk/sdk-docs/features/migrations/migration-variation
sdk: node-server-sdk
kind: reference
lang: typescript
description: Migration stage evaluation (migrationVariation) for Node.js (server-side) SDK v9.
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only
---

```typescript
  import { LDContext, LDMigrationStage } from '@launchdarkly/node-server-sdk';

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
