---
id: node-server-sdk/sdk-docs/migration-6-to-7-understanding-changes-to-private-attributes-7-0-syntax-typescript
sdk: node-server-sdk
kind: reference
lang: typescript
description: "7.0 syntax, TypeScript in section \"Understanding changes to private attributes\""
---

```ts
import * as ld from 'launchdarkly-node-server-sdk';
const options: ld.LDOptions {
  privateAttributes: ['email', 'address']
}

const client = ld.init('YOUR_SDK_KEY', options);
```
