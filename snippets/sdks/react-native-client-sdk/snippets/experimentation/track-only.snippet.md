---
id: react-native-client-sdk/experimentation/track-only
sdk: react-native-client-sdk
kind: reference
lang: tsx
description: Experimentation onboarding (track only) for react-native-client-sdk — initialize, identify, and add a trackMetric helper for conversion events.
validation:
  scaffold: react-native-client-sdk/scaffolds/react-native-syntax-only
---

```tsx
import React, { useEffect, useState, useCallback } from 'react';
import {
  AutoEnvAttributes,
  LDProvider,
  ReactNativeLDClient,
} from '@launchdarkly/react-native-client-sdk';

// This is your mobile key.
const ldClient = new ReactNativeLDClient('YOUR_MOBILE_KEY', AutoEnvAttributes.Enabled, {
  debug: true,
  applicationInfo: {
    id: 'ld-rn-test-app',
    version: '0.0.1',
  },
});

// A "context" is a data object representing users, devices, organizations, and other entities.
const context = { kind: 'user', key: 'EXAMPLE_CONTEXT_KEY', email: 'EXAMPLE_EMAIL' };

export default function App() {
  const [ready, setReady] = useState<boolean>(false);
  // Wait for identify to complete before rendering, so no flag evaluates
  // against defaults during startup.
  useEffect(() => {
    (async () => {
      await ldClient.identify(context);
      setReady(true);
    })();
  }, []);
  // Call trackMetric when a metric action occurs in your app —
  // a tap, a form submit, a screen view, a custom event, whatever your metric measures.
  const trackMetric = useCallback(
    (metricKey: string, data?: unknown) => {
      ldClient.track(metricKey, data);
    },
    [],
  );
  if (!ready) return null;
  return (
    <LDProvider client={ldClient}>
      <YourComponent />
    </LDProvider>
  );
}
```
