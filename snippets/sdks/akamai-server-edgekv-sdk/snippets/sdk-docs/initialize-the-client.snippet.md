---
id: akamai-server-edgekv-sdk/sdk-docs/initialize-the-client
sdk: akamai-server-edgekv-sdk
kind: reference
lang: typescript
file: akamai-server-edgekv-sdk/sdk-docs/initialize-the-client.ts
description: "Akamai edge SDK in section \"Initialize the client\""
validation:
  scaffold: akamai-server-edgekv-sdk/scaffolds/edge-akamai-init-runner
---

```ts
const ldClient = init({
  sdkKey: 'example-client-side-id',
  namespace: 'your-edgekv-namespace',
  group: 'your-edgekv-group-id'
});
```
