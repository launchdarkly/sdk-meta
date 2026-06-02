---
id: electron-client-sdk/sdk-docs/features/config/service-endpoint-configuration-ts-eu
sdk: electron-client-sdk
kind: reference
lang: typescript
description: Service endpoint configuration example for Electron.
---

```ts
import * as LDElectron from 'launchdarkly-electron-client-sdk';

const options: LDElectron.LDOptions = {
  streamUrl: 'https://clientstream.eu.launchdarkly.com',
  baseUrl: 'https://clientsdk.eu.launchdarkly.com',
  eventsUrl: 'https://events.eu.launchdarkly.com'
};
```
