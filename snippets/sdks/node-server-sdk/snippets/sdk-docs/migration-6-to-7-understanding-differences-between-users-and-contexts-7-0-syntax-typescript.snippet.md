---
id: node-server-sdk/sdk-docs/migration-6-to-7-understanding-differences-between-users-and-contexts-7-0-syntax-typescript
sdk: node-server-sdk
kind: reference
lang: typescript
description: "7.0 syntax, TypeScript in section \"Understanding differences between users and contexts\""
---

```ts
const context: ld.LDContext = {
  kind: 'multi',
  user: { key: 'example-user-key' },
  device: { key: 'example-device-key' }
}
```
