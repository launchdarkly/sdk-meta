---
id: node-client-sdk/sdk-docs/initialize-the-client-node-js-sdk-v3-typescript
sdk: node-client-sdk
kind: reference
lang: typescript
description: "Node.js SDK v3 (TypeScript) in section \"Initialize the client\""
validation:
  scaffold: node-client-sdk/scaffolds/node-client-syntax-only
---

```ts
// You'll need this context later, but you can ignore it for now.
const context: LaunchDarkly.LDContext = {
  kind: 'user',
  key: 'example-user-key'
};
const client = LaunchDarkly.initialize('example-client-side-id', context);
try {
  await client.waitForInitialization(5);
  // initialization succeeded, flag values are now available
} catch (err) {
  // initialization failed or did not complete before timeout
}
```
