---
id: js-client-sdk/sdk-docs/features/privateattrs/context-ts
sdk: js-client-sdk
kind: reference
lang: ts
description: Marking specific attributes private in the context object for JavaScript SDK v3.x+ (TypeScript).
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only

---

```ts
import * as ld from 'launchdarkly-js-client-sdk';

const context: ld.LDContext = {
  key: 'example-context-key',
  name: 'Sandy Smith',
  email: 'sandy@example.com',
  _meta: {
    privateAttributes: ['email'],
  }
};

```
