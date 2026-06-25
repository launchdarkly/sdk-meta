---
id: electron-client-sdk/sdk-docs/features/bootstrapping/bootstrapping-ts
sdk: electron-client-sdk
kind: reference
lang: typescript
description: Bootstrapping example for Electron (TypeScript).
validation:
  scaffold: electron-client-sdk/scaffolds/electron-syntax-only
---

```ts
import * as LaunchDarkly from 'launchdarkly-electron-client-sdk';

const options: LaunchDarkly.LDOptions = {
  bootstrap: {
    flagKey1: flagValue1,
    flagKey2: flagValue2,
  },
};

const client = LaunchDarkly.initializeInMain('example-client-side-id', user, options);
```
