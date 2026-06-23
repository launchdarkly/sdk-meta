---
id: node-server-sdk/sdk-docs/features/contextconfig/multi-context-ts
sdk: node-server-sdk
kind: reference
lang: typescript
description: Multi-context example for Node.js (server-side) SDK v7.x and later (TypeScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```typescript
const context: ld.LDContext = {
  kind: 'multi',
  user: { key: 'example-user-key' },
  device: { key: 'example-device-key' }
}
```
