---
id: electron-client-sdk/sdk-docs/features/relay-proxy-config/proxy-mode/proxy-mode-ts
sdk: electron-client-sdk
kind: reference
lang: typescript
description: Proxy mode configuration example for Electron.
validation:
  scaffold: electron-client-sdk/scaffolds/electron-syntax-only
---

```ts
import * as LDElectron from 'launchdarkly-electron-client-sdk';

const options: LDElectron.LDOptions = {
  streamUrl: 'https://your-relay-proxy.com:8030',
  baseUrl: 'https://your-relay-proxy.com:8030',
  eventsUrl: 'https://your-relay-proxy.com:8030',
};
```
