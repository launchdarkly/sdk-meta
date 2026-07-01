---
id: cloudflare-server-sdk/sdk-docs/initialize-the-client
sdk: cloudflare-server-sdk
kind: reference
lang: javascript
file: cloudflare-server-sdk/sdk-docs/initialize-the-client.ts
description: "Cloudflare edge SDK in section \"Initialize the client\""
validation:
  scaffold: cloudflare-server-sdk/scaffolds/edge-cloudflare-init-runner
---

```js
const client = init('example-client-side-id', env.LD_KV);
await client.waitForInitialization();
```
