---
id: fastly-server-sdk/sdk-docs/initialize-the-client
sdk: fastly-server-sdk
kind: reference
lang: typescript
file: fastly-server-sdk/sdk-docs/initialize-the-client.ts
description: "Fastly edge SDK in section \"Initialize the client\""
validation:
  scaffold: fastly-server-sdk/scaffolds/edge-fastly-init-runner
---

```ts
const KV_STORE_NAME = 'launchdarkly';
const EVENTS_BACKEND_NAME = 'launchdarkly';
const store = new KVStore(KV_STORE_NAME);

async function handleRequest(event: FetchEvent) {
  const ldClient = init('example-client-side-id', store, {
    eventsBackendName: EVENTS_BACKEND_NAME,
  });

  await ldClient.waitForInitialization();

  // The rest of your handler code goes here
}
```
