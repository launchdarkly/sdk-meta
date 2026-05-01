---
id: node-server-sdk/sdk-docs/migration-8-to-9-understanding-what-was-removed-node-js-sdk-v7-context-with-key-typescript
sdk: node-server-sdk
kind: reference
lang: typescript
description: "Node.js SDK v7+, context with key (TypeScript) in section \"Understanding what was removed\""
---

```ts
import * as ld from 'launchdarkly-node-server-sdk';

const context: ld.LDContext = {
  kind: 'user',
  key: 'example-user-key',
};
```
