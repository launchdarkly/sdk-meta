---
id: node-server-sdk/sdk-docs/features/migrations/read-write
sdk: node-server-sdk
kind: reference
lang: typescript
description: Migration read and write example for Node.js (server-side) SDK v9.
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only
---

```typescript
import { LDContext, LDMigrationStage, createMigration } from '@launchdarkly/node-server-sdk';

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
