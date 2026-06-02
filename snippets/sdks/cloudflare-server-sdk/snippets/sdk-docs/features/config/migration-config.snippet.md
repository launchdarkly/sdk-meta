---
id: cloudflare-server-sdk/sdk-docs/features/config/migration-config
sdk: cloudflare-server-sdk
kind: reference
lang: typescript
description: Migration configuration example for the Cloudflare SDK v2.2.2+ — read/write functions, execution mode, latency/error tracking.
---

```ts
import {
  createMigration,
  init,
  LDConcurrentExecution,
  LDMigrationError,
  LDMigrationOptions,
  LDMigrationSuccess,
} from '@launchdarkly/cloudflare-server-sdk';

const options: LDMigrationOptions = {
  readNew: async(key?: string) => {
    console.log("Reading from new: ", key);
    return LDMigrationSuccess(true);
  },
  readOld: async(key?: string) => {
    console.log("Reading from new: ", key);
    return LDMigrationSuccess(true);
  },
  writeNew: async(params?: {key: string, value: string}) => {
    console.log("Writing to new: ", params);
    // if failure - can throw an exception
    throw new Error('example exception')
  },
  writeOld: async(params?: {key: string, value: string}) => {
    console.log("Writing to old: ", params);
    // if failure - can return the failure
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

const client = init('YOUR_SDK_KEY', env.LD_KV, { sendEvents: true });
const migration = createMigration(client, options);

```
