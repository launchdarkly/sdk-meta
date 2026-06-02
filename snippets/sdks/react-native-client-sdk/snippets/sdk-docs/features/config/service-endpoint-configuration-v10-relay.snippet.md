---
id: react-native-client-sdk/sdk-docs/features/config/service-endpoint-configuration-v10-relay
sdk: react-native-client-sdk
kind: reference
lang: typescript
description: Service endpoint configuration example for React Native.
---

```ts
import { LDOptions } from '@launchdarkly/react-native-client-sdk'

let options: LDOptions = {
  streamUri: 'https://your-relay-proxy.com:8030',
  baseUri: 'https://your-relay-proxy.com:8030',
  eventsUri: 'https://your-relay-proxy.com:8030',
};
```
