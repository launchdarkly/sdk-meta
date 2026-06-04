---
id: react-native-client-sdk/sdk-docs/features/config/service-endpoint-configuration-v10-federal
sdk: react-native-client-sdk
kind: reference
lang: typescript
description: Service endpoint configuration example for React Native.
validation:
  scaffold: react-native-client-sdk/scaffolds/react-native-syntax-only

---

```ts
import { LDOptions } from '@launchdarkly/react-native-client-sdk'

let options: LDOptions = {
  streamUri: 'https://clientstream.launchdarkly.us',
  baseUri: 'https://clientsdk.launchdarkly.us',
  eventsUri: 'https://events.launchdarkly.us',
};
```
