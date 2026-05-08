---
id: react-native-client-sdk/scaffolds/init-runner-observability
sdk: react-native-client-sdk
kind: scaffold
lang: tsx
file: App.tsx
description: |
  End-to-end runner for the `observability/initialize` snippet body.

  The wrappee body declares
    `const client = new ReactNativeLDClient('SDK_KEY', AutoEnvAttributes.Enabled, { plugins: [...] });`
  and assumes `ReactNativeLDClient`, `AutoEnvAttributes`, and
  `Observability` are in scope (the symbols come from the matching
  `observability/import` snippet). This scaffold supplies those
  imports at module scope, splices the body, and renders a sentinel
  component wrapped in `<LDProvider client={client}>`. The sentinel
  awaits `client.waitForInitialization` and renders the EXAM-HELLO
  line on success.

  We don't assert observability data flows back to LaunchDarkly —
  just that the SDK starts cleanly with the o11y plugin attached.

  The wrappee's `'SDK_KEY'` literal is substituted with the live
  `LAUNCHDARKLY_MOBILE_KEY` env var via the snippet's
  `validation.placeholders` map.
inputs:
  body:
    type: string
    description: The wrappee init snippet body, embedded after key substitution.
validation:
  runtime: react-native-client
  entrypoint: App.tsx
  companions:
    - react-native-client-sdk/scaffolds/init-runner-welcome
---

```tsx
import React, { useEffect, useState } from 'react';
import { Text, View } from 'react-native';
import {
  AutoEnvAttributes,
  LDProvider,
  ReactNativeLDClient,
  useLDClient,
} from '@launchdarkly/react-native-client-sdk';
import { Observability } from '@launchdarkly/observability-react-native';

// The wrappee body declares
//   const client = new ReactNativeLDClient('SDK_KEY', AutoEnvAttributes.Enabled, { plugins: [...] });
// Splicing it here at module scope binds `client` for the LDProvider
// below.
{{ body }}

const YourComponent = () => {
  const ldClient = useLDClient();
  const [ready, setReady] = useState(false);
  useEffect(() => {
    if (!ldClient) return;
    let cancelled = false;
    (async () => {
      try {
        await ldClient.waitForInitialization(5_000);
        if (!cancelled) setReady(true);
      } catch (e) {
        // leave ready=false; the test will time out and report the real error
      }
    })();
    return () => { cancelled = true; };
  }, [ldClient]);
  if (!ready) {
    return <View><Text>waiting for LD client</Text></View>;
  }
  return <View><Text>feature flag evaluates to true</Text></View>;
};

const App = () => (
  <LDProvider client={client}>
    <YourComponent />
  </LDProvider>
);

export default App;
```
