---
id: vercel-server-sdk/sdk-docs/initialize-the-client
sdk: vercel-server-sdk
kind: reference
lang: javascript
file: vercel-server-sdk/sdk-docs/initialize-the-client.ts
description: "Vercel edge SDK in section \"Initialize the client\""
validation:
  scaffold: vercel-server-sdk/scaffolds/edge-vercel-init-runner
---

```js
const edgeConfigClient = createClient(process.env.EDGE_CONFIG)
const ldClient = init('example-client-side-id', edgeConfigClient)

await ldClient.waitForInitialization()
```
