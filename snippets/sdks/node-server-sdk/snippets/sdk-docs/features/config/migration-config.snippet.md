---
id: node-server-sdk/sdk-docs/features/config/migration-config
sdk: node-server-sdk
kind: reference
lang: typescript
description: Migration configuration example for the Node.js (server-side) SDK v9 — read/write functions, execution mode, latency/error tracking.
---

```ts
import { LDMigrationOptions, init, createMigration } from '@launchdarkly/node-server-sdk';
const options: LDMigrationOptions = {
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

const client = init('YOUR_SDK_KEY');
const migration = createMigration(client, options);

```
