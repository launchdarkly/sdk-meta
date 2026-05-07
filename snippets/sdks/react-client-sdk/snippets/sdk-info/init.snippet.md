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
import { createRoot } from 'react-dom/client';
import { createLDReactProvider, LDContext } from '@launchdarkly/react-sdk';

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
const LDReactProvider = createLDReactProvider('YOUR_CLIENT_SIDE_ID', context);

createRoot(document.getElementById('root') as HTMLElement).render(
  <LDReactProvider>
    <App />
  </LDReactProvider>,
);
```
