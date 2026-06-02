---
id: fastly-server-sdk/sdk-docs/features/config/index
sdk: fastly-server-sdk
kind: reference
lang: typescript
description: SDK configuration example for Fastly.
---

```ts
import { KVStore } from 'fastly:kv-store';
import { init, LDOptions } from '@launchdarkly/fastly-server-sdk';

const KV_STORE_NAME = 'launchdarkly';
const EVENTS_BACKEND_NAME = 'launchdarkly';
const store = new KVStore(KV_STORE_NAME);

async function handleRequest(event: FetchEvent) {
  const ldClient = init('example-client-side-id', store, {
    eventsBackendName: EVENTS_BACKEND_NAME,
  });

  await ldClient.waitForInitialization();

  ...
}
```
