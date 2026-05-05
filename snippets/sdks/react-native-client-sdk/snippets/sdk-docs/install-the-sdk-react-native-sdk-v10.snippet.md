---
id: react-native-client-sdk/sdk-docs/install-the-sdk-react-native-sdk-v10
sdk: react-native-client-sdk
kind: reference
lang: typescript
description: "React Native SDK v10 in section \"Install the SDK\""
# Bucket C: react-native validator harness expects a 2-file shape
# (App.tsx + src/welcome.tsx) baked into the prebuilt jest project,
# whereas the react-native-syntax-only scaffold writes a single App.js.
# The harness needs a separate parse-only mode (or the scaffold needs
# to output the App.tsx + src/welcome.tsx pair) before sdk-docs
# fragments can validate. See _sdk-docs-port-notes.md.
---

```ts
  import { LDProvider, ReactNativeLDClient } from '@launchdarkly/react-native-client-sdk';

  // optional observability plugin, requires React Native SDK v10.10+
  import { Observability, LDObserve } from '@launchdarkly/observability-react-native';
```
