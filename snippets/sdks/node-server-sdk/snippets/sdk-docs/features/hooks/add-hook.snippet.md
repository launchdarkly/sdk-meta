---
id: node-server-sdk/sdk-docs/features/hooks/add-hook
sdk: node-server-sdk
kind: reference
lang: ts
description: Adding a hook to an existing client for Node.js (server-side).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only
---

```ts

const client = ld.init('YOUR_SDK_KEY', options);
client.addHook(new ExampleHook());
```
