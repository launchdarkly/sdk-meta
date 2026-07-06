---
id: cloudflare-server-sdk/sdk-docs/initialize-the-client-with-events
sdk: cloudflare-server-sdk
kind: reference
lang: javascript
file: cloudflare-server-sdk/sdk-docs/initialize-the-client-with-events.ts
description: "Cloudflare edge SDK in section \"Initialize the client\" (with events)"
validation:
  scaffold: cloudflare-server-sdk/scaffolds/edge-cloudflare-init-runner
---

```js
const client = init('example-client-side-id', env.LD_KV, { sendEvents: true });
await client.waitForInitialization();
```
