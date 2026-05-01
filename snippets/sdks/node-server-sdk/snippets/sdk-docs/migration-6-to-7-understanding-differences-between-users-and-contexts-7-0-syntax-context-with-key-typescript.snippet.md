---
id: node-server-sdk/sdk-docs/migration-6-to-7-understanding-differences-between-users-and-contexts-7-0-syntax-context-with-key-typescript
sdk: node-server-sdk
kind: reference
lang: typescript
description: "7.0 syntax, context with key, TypeScript in section \"Understanding differences between users and contexts\""
---

```ts
import * as ld from 'launchdarkly-node-server-sdk';

const context: ld.LDContext = {
  kind: 'user',
  key: 'example-user-key',
};
```
