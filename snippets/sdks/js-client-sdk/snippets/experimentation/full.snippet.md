---
id: js-client-sdk/experimentation/full
sdk: js-client-sdk
kind: reference
lang: javascript
description: Full experimentation onboarding for js-client-sdk — initialize, identify on login/eligibility, evaluate, and track conversions.
# Bucket C: newly proposed experimentation onboarding snippet, not
# standalone-runnable (exports helpers consumed by your app). No validation
# block yet. See _experimentation-port-notes.md.
---

```javascript
import { createClient } from '@launchdarkly/js-client-sdk';

// Initialize once — extra clients can cause inconsistent experiment results.
// If you already know the user's key at startup, initialize with it directly and
// skip identifyUser() below. Use an anonymous context only when you don't yet
// know who the user is, and reuse that key when the user becomes known.
export const ldClient = createClient('YOUR_CLIENT_SIDE_ID', {
  kind: 'user',
  key: 'EXAMPLE_CONTEXT_KEY', // use a consistent key so the same user gets the same experience
  anonymous: true,
});

// await start — you only need waitForInitialization() if you wait somewhere
// other than where you start the client.
await ldClient.start();

// Call this only when the user becomes known after startup (consent/login/mid-funnel).
export async function identifyUser({ userKey, attributes }) {
  await ldClient.identify({
    ...attributes, // any attributes that affect targeting or eligibility (spread first so it can't override the fields below)
    kind: 'user',
    key: userKey, // use the logged-in user's ID so experiment assignment stays consistent
    anonymous: false,
  });
}

// Call this ONLY where the user encounters the experience.
export function evalExperimentFlag(flagKey, defaultValue) {
  return ldClient.variation(flagKey, defaultValue);
}

// Call this when the user completes a metric action.
// Use the same user key you used when evaluating the flag — mismatched keys break conversion tracking.
// The data argument is optional and accepts any shape your metric needs.
export function trackMetric(metricKey, data) {
  ldClient.track(metricKey, data);
}

// The SDK batches and flushes events automatically. Don't add manual flush()
// calls — they're unnecessary and actively harmful to performance.
// Don't skip or cache variation() calls to reduce exposure counts — LaunchDarkly deduplicates them automatically.
```
