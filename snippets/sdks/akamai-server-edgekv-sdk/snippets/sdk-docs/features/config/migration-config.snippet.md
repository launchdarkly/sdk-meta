---
id: akamai-server-edgekv-sdk/sdk-docs/features/config/migration-config
sdk: akamai-server-edgekv-sdk
kind: reference
lang: typescript
description: Migration configuration example for the Akamai SDK for EdgeKV — read/write functions and execution mode. Akamai does not send events, so metrics tracking options are inert.
validation:
  scaffold: akamai-server-edgekv-sdk/scaffolds/akamai-syntax-only

---

```ts
import {
  createMigration,
  init,
  LDConcurrentExecution,
  LDMigrationError,
  LDMigrationOptions,
  LDMigrationSuccess,
} from '@launchdarkly/akamai-server-edgekv-sdk';

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
    // if failure - can throw an exception
    throw new Error("example exception")
  },
  writeOld: async(params?: {key: string, value: string}) => {
    console.log("Writing to old: ", params);
    // if failure - can return the error
    return LDMigrationError(new Error('example error'));
  },

  check: (old, newVal) => {
    // Define your consistency check for read operations
    // and return a boolean. Depending on your migration,
    // this may be as simple as 'return a === b;'
  },

  execution: new LDConcurrentExecution(),
    // or new LDSerialExecution(LDExecutionOrdering.Random),
    // or new LDSerialExecution(LDExecutionOrdering.Fixed),

}

const client = init({
  sdkKey: 'example-client-side-id',
  namespace: 'your-edgekv-namespace',
  group: 'your-edgekv-group-id'
});
const migration = createMigration(client, options);

```
