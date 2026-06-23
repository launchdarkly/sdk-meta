---
id: electron-client-sdk/sdk-docs/features/privateattrs/private-attribute-names-ts
sdk: electron-client-sdk
kind: reference
lang: ts
description: Marking specific attributes private for Electron (TypeScript).
validation:
  scaffold: electron-client-sdk/scaffolds/electron-syntax-only

---

```ts
import * as LDElectron from 'launchdarkly-electron-client-sdk';

const user: LDElectron.LDUser = {
  key: 'example-user-key',
  name: 'Sandy Smith',
  email: 'sandy@example.com',
  privateAttributeNames: ['email'],
};

const client = LDElectron.initializeInMain('example-client-side-id', user, {
  privateAttributeNames: ['email'],
});
```
