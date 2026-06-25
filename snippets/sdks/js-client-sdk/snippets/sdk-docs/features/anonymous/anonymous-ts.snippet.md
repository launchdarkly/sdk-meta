---
id: js-client-sdk/sdk-docs/features/anonymous/anonymous-ts
sdk: js-client-sdk
kind: reference
lang: typescript
description: Anonymous context example for JavaScript, SDK v3.0 (TypeScript).
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only
---

```typescript
import * as ld from 'launchdarkly-js-client-sdk';

const anonymousContext: ld.LDContext = {
  kind: 'user',
  anonymous: true
};

// A multi-context can contain both anonymous and non-anonymous contexts.
// Here, the organization is not anonymous.
const multiContext: ld.LDContext = {
  kind: 'multi',
  user: anonymousContext,
  org: {
    key: 'example-organization-key',
    name: 'Example organization name'
  }
}
```
