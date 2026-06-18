---
id: node-server-sdk/sdk-docs/features/contextconfig/multi-context-ai
sdk: node-server-sdk
kind: reference
lang: typescript
description: Multi-context example for Node.js (server-side) AI SDK (TypeScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```typescript
const context: LDContext = {
  kind: 'multi',
  user: { key: 'example-user-key' },
  device: { key: 'example-device-key' }
}
```
