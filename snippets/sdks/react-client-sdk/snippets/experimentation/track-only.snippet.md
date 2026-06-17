---
id: react-client-sdk/experimentation/track-only
sdk: react-client-sdk
kind: reference
lang: tsx
description: Experimentation onboarding (track only) for react-client-sdk — initialize and add a trackMetric helper for conversion events.
validation:
  scaffold: react-client-sdk/scaffolds/react-syntax-only
---

```tsx
// Add the code below to the root of your React app.
import { StrictMode, useCallback } from 'react';
import { createRoot } from 'react-dom/client';
import { createLDReactProvider, useLDClient } from '@launchdarkly/react-sdk';

function App() {
  const ldClient = useLDClient();
  // Call trackMetric when a metric action occurs in your app —
  // a click, a form submit, a page view, a custom event, whatever your metric measures.
  const trackMetric = useCallback(
    (metricKey: string, data?: unknown) => {
      ldClient?.track(metricKey, data);
    },
    [ldClient],
  );
  return <div>Let your feature flags fly!</div>;
}

// A "context" is a data object representing users, devices, organizations, and other entities.
const context = {
  kind: 'user',
  key: 'EXAMPLE_CONTEXT_KEY',
  email: 'EXAMPLE_EMAIL',
};

// Initialize the SDK so flag values are ready before your app renders.
// This is your client-side ID.
const LDProvider = createLDReactProvider('YOUR_CLIENT_SIDE_ID', context);

createRoot(document.getElementById('root') as HTMLElement).render(
  <StrictMode>
    <LDProvider>
      <App />
    </LDProvider>
  </StrictMode>,
);
```
