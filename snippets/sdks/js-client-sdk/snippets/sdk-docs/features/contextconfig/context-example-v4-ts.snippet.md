---
id: js-client-sdk/sdk-docs/features/contextconfig/context-example-v4-ts
sdk: js-client-sdk
kind: reference
lang: typescript
description: Context example for JavaScript SDK v4.x (TypeScript).
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only

---

```typescript
import * as ld from '@launchdarkly/js-client-sdk';

const context: ld.LDContext = {
  kind: 'user',
  key: 'example-user-key',
  firstName: 'Sandy',
  lastName: 'Smith',
  email: 'sandy@example.com',
  groups: ['Acme', 'Global Health Services']
}
```
