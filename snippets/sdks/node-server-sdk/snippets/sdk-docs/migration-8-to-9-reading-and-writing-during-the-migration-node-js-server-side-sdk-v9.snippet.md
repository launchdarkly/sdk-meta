---
id: node-server-sdk/sdk-docs/migration-8-to-9-reading-and-writing-during-the-migration-node-js-server-side-sdk-v9
sdk: node-server-sdk
kind: reference
lang: typescript
description: "Node.js (server-side) SDK v9 in section \"Reading and writing during the migration\""
---

```ts
const ld = require('@launchdarkly/node-server-sdk');

const context: ld.LDContext = {
   kind: 'user',
   key: 'example-user-key',
   name: 'Sandy',
};

// this is the migration stage to use if the flag's migration stage
// is not available from LaunchDarkly
let defaultStage: ld.LDMigrationStage = LDMigrationStage.Off;

const migration = ld.createMigration(client, options);

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
