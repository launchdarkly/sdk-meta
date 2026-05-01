---
id: node-server-sdk/sdk-docs/migration-6-to-7-understanding-changes-to-alias-events-7-0-syntax-typescript
sdk: node-server-sdk
kind: reference
lang: typescript
description: "7.0 syntax, TypeScript in section \"Understanding changes to alias events\""
---

```ts
const context: ld.LDContext = {
  kind: 'multi',
  user: { key: "example-user-key" },
  device: { key: "example-device-key" }
}

client.identify(context)
```
