---
id: electron-client-sdk/sdk-docs/features/localstorage/localstorage-ts
sdk: electron-client-sdk
kind: reference
lang: typescript
description: Local storage caching example for Electron (TypeScript).
validation:
  scaffold: electron-client-sdk/scaffolds/electron-syntax-only
---

```ts
import * as LaunchDarkly from 'launchdarkly-electron-client-sdk';

const options: LaunchDarkly.LDOptions = { bootstrap: 'localStorage' };

const client = LaunchDarkly.initializeInMain('example-client-side-id', user, options);

```
