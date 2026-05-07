---
id: node-server-sdk/sdk-docs/initialize-the-client-node-js-sdk
sdk: node-server-sdk
kind: reference
lang: javascript
description: "Node.js SDK in section \"Initialize the client\""
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only
---

```js
const client = init(
  'YOUR_SDK_KEY',
  {
    // optional observability plugin, requires Node.js (server-side) SDK v9.10+
    plugins: [ new Observability(), ],
    // other options
  },
);
```
