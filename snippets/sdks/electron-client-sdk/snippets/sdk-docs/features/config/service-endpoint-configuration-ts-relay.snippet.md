---
id: electron-client-sdk/sdk-docs/features/config/service-endpoint-configuration-ts-relay
sdk: electron-client-sdk
kind: reference
lang: typescript
description: Service endpoint configuration example for Electron.
---

```ts
import * as LDElectron from 'launchdarkly-electron-client-sdk';

const options: LDElectron.LDOptions = {
  streamUrl: 'https://your-relay-proxy.com:8030',
  baseUrl: 'https://your-relay-proxy.com:8030',
  eventsUrl: 'https://your-relay-proxy.com:8030',
};
```
