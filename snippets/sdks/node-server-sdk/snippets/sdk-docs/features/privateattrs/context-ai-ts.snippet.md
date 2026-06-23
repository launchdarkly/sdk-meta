---
id: node-server-sdk/sdk-docs/features/privateattrs/context-ai-ts
sdk: node-server-sdk
kind: reference
lang: ts
description: Marking context attributes private for the Node.js (server-side) AI SDK (TypeScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```ts
const context: LDContext = {
  kind: 'user',
  key: 'example-user-key',
  email: 'sandy@example.com',
  _meta: {
    privateAttributes: ['email'],
  }
};
```
