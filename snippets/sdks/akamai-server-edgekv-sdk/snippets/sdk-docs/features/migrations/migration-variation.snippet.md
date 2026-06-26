---
id: akamai-server-edgekv-sdk/sdk-docs/features/migrations/migration-variation
sdk: akamai-server-edgekv-sdk
kind: reference
lang: typescript
description: Migration stage evaluation (migrationVariation) for Akamai SDK v1.0.9+.
validation:
  scaffold: akamai-server-edgekv-sdk/scaffolds/akamai-syntax-only
---

```typescript
  import {
    LDContext,
    LDMigrationStage,
  } from '@launchdarkly/akamai-server-edgekv-sdk';

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
