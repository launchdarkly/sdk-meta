---
id: node-server-sdk/sdk-docs/migration-8-to-9-configuring-the-migration-node-js-sdk-v9-0
sdk: node-server-sdk
kind: reference
lang: typescript
description: "Node.js SDK v9.0 in section \"Configuring the migration\""
---

```ts
import * as ld from '@launchdarkly/node-server-sdk';
const options: ld.LDMigrationOptions = {
  readNew: async(key?: string) => {
    console.log("Reading from new: ", key);
    return LDMigrationSuccess(true);
  },
  readOld: async(key?: string) => {
    console.log("Reading from old: ", key);
    return LDMigrationSuccess(true);
  },
  writeNew: async(params?: {key: string, value: string}) => {
    console.log("Writing to new: ", params);
    // if failure
    return LDMigrationError(new Error('example error'));
  },
  writeOld: async(params?: {key: string, value: string}) => {
    console.log("Writing to old: ", params);
    // if failure
    return LDMigrationError(new Error('example error'));
  },

  check: (old, new) => {
    // Define your consistency check for read operations
    // and return a boolean. Depending on your migration,
    // this may be as simple as 'return a === b;'
  },

  execution: new LDConcurrentExecution(),
    // or new LDSerialExecution(LDExecutionOrdering.Random),
    // or new LDSerialExecution(LDExecutionOrdering.Fixed),

  latencyTracking: true, // defaults to true
  errorTracking: true, // defaults to true
}

const client = ld.init('YOUR_SDK_KEY');
const migration = ld.createMigration(client, options);
```
