---
id: electron-client-sdk/sdk-docs/features/privateattrs/all-attributes-private-ts
sdk: electron-client-sdk
kind: reference
lang: ts
description: Marking all attributes private for Electron (TypeScript).
validation:
  scaffold: electron-client-sdk/scaffolds/electron-syntax-only

---

```ts
import * as LDElectron from 'launchdarkly-electron-client-sdk';

const user: LDElectron.LDUser = {
  key: 'example-user-key',
  name: 'Sandy Smith',
  email: 'sandy@example.com',
};

const client = LDElectron.initializeInMain('example-client-side-id', user, {
  allAttributesPrivate: true,
});
```
