---
id: node-server-sdk/sdk-docs/features/privateattrs/config-ts
sdk: node-server-sdk
kind: reference
lang: ts
description: Private attribute configuration for Node.js SDK v7.0+ (TypeScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```ts
import { LDOptions } from '@launchdarkly/node-server-sdk';

// All attributes marked private
const options: LDOptions = {
  allAttributesPrivate: true
};

const client = ld.init('YOUR_SDK_KEY', options);

// Two attributes marked private
const optionsSomePrivate: LDOptions = {
  privateAttributes: ['email', 'address']
};

const clientSomePrivate = ld.init('YOUR_SDK_KEY', optionsSomePrivate);
```
