---
id: node-server-sdk/sdk-docs/features/anonymous/anonymous-ai-ts
sdk: node-server-sdk
kind: reference
lang: typescript
description: Anonymous context example for the Node.js (server-side) AI SDK (TypeScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only
---

```typescript
const context: LDContext = {
  kind: 'user',
  key: 'example-user-key',
  anonymous: true,
}
```
