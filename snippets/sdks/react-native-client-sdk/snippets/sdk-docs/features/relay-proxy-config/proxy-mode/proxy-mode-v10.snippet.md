---
id: react-native-client-sdk/sdk-docs/features/relay-proxy-config/proxy-mode/proxy-mode-v10
sdk: react-native-client-sdk
kind: reference
lang: typescript
description: Proxy mode configuration example for React Native.
validation:
  scaffold: react-native-client-sdk/scaffolds/react-native-syntax-only
---

```ts
import { LDOptions } from '@launchdarkly/react-native-client-sdk'

let options: LDOptions = {
  streamUri: 'https://your-relay-proxy.com:8030',
  baseUri: 'https://your-relay-proxy.com:8030',
  eventsUri: 'https://your-relay-proxy.com:8030',
};
```
