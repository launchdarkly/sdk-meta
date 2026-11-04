---
id: react-client-sdk/sdk-info/init-observability
sdk: react-client-sdk
kind: init
lang: tsx
file: react-client-sdk/init-observability.txt
description: Client initialization snippet for react-client-sdk with the observability and session replay plugins.
validation:
  scaffold: react-client-sdk/scaffolds/init-runner
  placeholders:
    YOUR_CLIENT_SIDE_ID: LAUNCHDARKLY_CLIENT_SIDE_ID
---

```tsx
// Add the code below to the root of your React app.
import { createRoot } from 'react-dom/client';
import { createLDReactProvider, LDContext } from '@launchdarkly/react-sdk';
import Observability from '@launchdarkly/observability';
import SessionReplay from '@launchdarkly/session-replay';

function App() {
  return <div>Let your feature flags fly!</div>
}

// A "context" is a data object representing users, devices, organizations, and other entities.
const context: LDContext = {
  kind: 'user',
  key: 'EXAMPLE_CONTEXT_KEY',
  email: 'biz@face.dev',
};

// This is your client-side ID.
// privacySetting 'strict' masks all text and inputs in session replays.
const LDReactProvider = createLDReactProvider('YOUR_CLIENT_SIDE_ID', context, {
  ldOptions: {
    plugins: [
      new Observability(),
      new SessionReplay({ privacySetting: 'strict' }),
    ],
  },
});

createRoot(document.getElementById('root') as HTMLElement).render(
  <LDReactProvider>
    <App />
  </LDReactProvider>,
);
```
