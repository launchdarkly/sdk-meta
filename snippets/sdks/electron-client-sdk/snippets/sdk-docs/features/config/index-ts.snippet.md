---
id: electron-client-sdk/sdk-docs/features/config/index-ts
sdk: electron-client-sdk
kind: reference
lang: typescript
description: SDK configuration example for Electron.
validation:
  scaffold: electron-client-sdk/scaffolds/electron-syntax-only

---

```ts
import * as LDElectron from 'launchdarkly-electron-client-sdk'

const options: LDElectron.LDOptions = { allAttributesPrivate: true };
const client = LDElectron.initializeInMain('example-client-side-id', user, options);
```
