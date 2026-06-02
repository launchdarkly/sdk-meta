---
id: electron-client-sdk/sdk-docs/features/config/service-endpoint-configuration-ts-federal
sdk: electron-client-sdk
kind: reference
lang: typescript
description: Service endpoint configuration example for Electron.
---

```ts
import * as LDElectron from 'launchdarkly-electron-client-sdk';

const options: LDElectron.LDOptions = {
  streamUrl: 'https://clientstream.launchdarkly.us',
  baseUrl: 'https://clientsdk.launchdarkly.us',
  eventsUrl: 'https://events.launchdarkly.us'
};
```
