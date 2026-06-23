---
id: electron-client-sdk/sdk-docs/features/contextconfig/context-example-ts
sdk: electron-client-sdk
kind: reference
lang: typescript
description: User example for Electron SDK (TypeScript).
validation:
  scaffold: electron-client-sdk/scaffolds/electron-syntax-only

---

```typescript
import * as LDElectron from 'launchdarkly-electron-client-sdk';

const user: LDElectron.LDUser = {
  key: 'example-user-key',
  firstName: 'Sandy',
  lastName: 'Smith',
  email: 'sandy@example.com',
  custom: {
    groups: ['Acme', 'Global Health Services'],
  },
};
```
