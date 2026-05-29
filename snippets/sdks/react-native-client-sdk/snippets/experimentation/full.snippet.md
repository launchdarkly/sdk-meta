---
id: react-native-client-sdk/experimentation/full
sdk: react-native-client-sdk
kind: reference
lang: tsx
description: Full experimentation onboarding for react-native-client-sdk — initialize, identify on login/eligibility, and track conversions.
# Bucket C: newly proposed experimentation onboarding snippet, not
# standalone-runnable (renders your own YourComponent). No validation block
# yet. See _experimentation-port-notes.md.
---

```tsx
import React, { useEffect, useState, useCallback } from 'react';
import {
  AutoEnvAttributes,
  LDProvider,
  ReactNativeLDClient,
} from '@launchdarkly/react-native-client-sdk';

// Initialize once — extra clients can cause inconsistent experiment results.
const ldClient = new ReactNativeLDClient('YOUR_MOBILE_KEY', AutoEnvAttributes.Enabled, {
  debug: true,
  applicationInfo: {
    id: 'ld-rn-test-app',
    version: '0.0.1',
  },
});

// A "context" is a data object representing users, devices, organizations, and other entities.
// If you already know the user's key at startup, initialize with it directly.
// Use an anonymous context only when you don't yet know who the user is.
const context = {
  kind: 'user',
  key: 'EXAMPLE_CONTEXT_KEY', // use a consistent key so the same user gets the same experience
  anonymous: true,
};

export default function App() {
  const [ready, setReady] = useState<boolean>(false);

  // Identify with the initial context on startup so flags are usable right away.
  useEffect(() => {
    (async () => {
      await ldClient.identify(context);
      setReady(true);
    })();
  }, []);

  // Call this only when the user becomes known after startup (login/eligibility).
  // Reuse the anonymous key where possible so experiment assignment stays stable,
  // and wait for identify to finish before evaluating experiment flags.
  const identifyUser = useCallback(
    async ({ userKey, attributes }: { userKey: string; attributes: Record<string, unknown> }): Promise<void> => {
      await ldClient.identify({
        kind: 'user',
        key: userKey, // use the logged-in user's ID so experiment assignment stays consistent
        anonymous: false,
        ...attributes, // any attributes that affect targeting or eligibility
      });
    },
    []
  );

  // Call this when the user completes a metric action.
  // Use the same user key you used when evaluating the flag — mismatched keys break conversion tracking.
  // The data argument is optional and accepts any shape your metric needs.
  const trackMetric = useCallback(
    (metricKey: string, data?: unknown): void => {
      ldClient.track(metricKey, data);
    },
    []
  );

  // The SDK batches and flushes events automatically, including when the app is
  // backgrounded. Don't add manual flush() calls or AppState listeners — they're
  // unnecessary and make real problems harder to spot.
  // Don't skip or cache flag evaluations to reduce exposure counts — LaunchDarkly deduplicates them automatically.

  if (!ready) return null;

  return (
    <LDProvider client={ldClient}>
      <YourComponent
        identifyUser={identifyUser}
        onConversion={() => trackMetric('YOUR_METRIC_KEY', /* optional data */)}
      />
    </LDProvider>
  );
}
```
