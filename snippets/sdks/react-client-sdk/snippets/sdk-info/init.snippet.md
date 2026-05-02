---
id: react-client-sdk/sdk-info/init
sdk: react-client-sdk
kind: init
lang: tsx
file: react-client-sdk/init.txt
description: Client initialization snippet for react-client-sdk.
validation:
  scaffold: react-client-sdk/scaffolds/init-runner
  placeholders:
    YOUR_CLIENT_SIDE_ID: LAUNCHDARKLY_CLIENT_SIDE_ID
---

```tsx
// Add the code below to the root of your React app.
import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { LDProvider } from 'launchdarkly-react-client-sdk';

function App() {
  return <div>Let your feature flags fly!</div>
}

// A "context" is a data object representing users, devices, organizations, and other entities.
const context = {
  kind: 'user',
  key: 'EXAMPLE_CONTEXT_KEY',
  email: 'biz@face.dev',
};

// This is your client-side ID.
createRoot(document.getElementById('root') as HTMLElement).render(
  <StrictMode>
    <LDProvider clientSideID="YOUR_CLIENT_SIDE_ID" context={context}>
      <App />
    </LDProvider>
  </StrictMode>,
);
```
