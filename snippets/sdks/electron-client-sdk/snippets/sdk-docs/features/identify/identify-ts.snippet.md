---
id: electron-client-sdk/sdk-docs/features/identify/identify-ts
sdk: electron-client-sdk
kind: reference
lang: typescript
description: Identify example for Electron (TypeScript).
validation:
  scaffold: electron-client-sdk/scaffolds/electron-syntax-only
---

```typescript
import * as LDElectron from 'launchdarkly-electron-client-sdk';

const newUser: LDElectron.LDUser = {
  key: 'someone-else', name: 'John',
};

client.identify(newUser, null, (err, newFlags) => {
  console.log('value of flag for this user is: ' + newFlags['example-flag-key']);
  console.log('this should be the same: ' + client.variation('example-flag-key'));
});

// or:
client.identify(newUser).then((newFlags) => {
  // as above
});
```
