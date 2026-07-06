---
id: vercel-server-sdk/sdk-docs/initialize-the-client-with-events
sdk: vercel-server-sdk
kind: reference
lang: javascript
file: vercel-server-sdk/sdk-docs/initialize-the-client-with-events.ts
description: "Vercel edge SDK in section \"Initialize the client\" (with events)"
validation:
  scaffold: vercel-server-sdk/scaffolds/edge-vercel-init-events
---

```js
const ldClient = init('example-client-side-id', edgeConfigClient, { sendEvents: true });
await ldClient.waitForInitialization();
```
