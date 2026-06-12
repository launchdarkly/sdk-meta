---
id: vercel-server-sdk/sdk-docs/features/migrations/read-write
sdk: vercel-server-sdk
kind: reference
lang: typescript
description: Migration read and write example for Vercel SDK v1.1.6+.
validation:
  scaffold: vercel-server-sdk/scaffolds/vercel-syntax-only
---

```typescript
import {
  createMigration,
  LDContext,
  LDMigrationStage,
} from '@launchdarkly/vercel-server-sdk';

const context: LDContext = {
   kind: 'user',
   key: 'example-user-key',
   name: 'Sandy',
};

// this is the migration stage to use if the flag's migration stage
// is not available from LaunchDarkly
let defaultStage: LDMigrationStage = LDMigrationStage.Off;

const migration = createMigration(client, options);

// when you need to perform a read in your application
migration.read(
  'example-migration-flag-key',
  context,
  defaultStage
);

// when you need to perform a write in your application
migration.write(
  'example-migration-flag-key',
  context,
  defaultStage
);
```
