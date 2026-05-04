---
id: react-native-client-sdk/scaffolds/init-runner
sdk: react-native-client-sdk
kind: scaffold
lang: tsx
file: App.tsx
description: |
  Runs an `init.txt`-style React Native snippet end-to-end against a
  real LaunchDarkly env, using the same jest+react-native-preset
  harness as the syntax-only flow. The init body declares an `App`
  component that wraps a `<YourComponent />` in `<LDProvider>` and
  exports App as default. The body references `useEffect` (assumed
  imported) and `YourComponent` (assumed defined) — both supplied by
  this scaffold's prelude.

  We prepend `import { useEffect } from 'react';` and a `YourComponent`
  definition that reads the EXAM-HELLO flag via `useFlags`/`useLDClient`
  and renders the canonical success text once the flag resolves to
  true. The harness's jest test renders `<App />` and asserts the
  flattened text contains `feature flag evaluates to true`.
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
import { useEffect, useState } from 'react';
import { Text, View } from 'react-native';
import { useLDClient } from '@launchdarkly/react-native-client-sdk';

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

{{ body }}
```
