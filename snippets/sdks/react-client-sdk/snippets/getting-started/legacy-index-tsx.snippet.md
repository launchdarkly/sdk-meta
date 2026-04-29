---
id: react-client-sdk/getting-started/legacy-index-tsx
sdk: react-client-sdk
kind: hello-world
lang: tsx
file: src/index.tsx
description: CRA index.tsx (legacy variant) wrapping the app with asyncWithLDProvider.
inputs:
  environmentId:
    type: client-side-id
    description: Client-side ID baked into the rendered source.
ld-application:
  slot: legacy-index-tsx
---

In `index.tsx`:

```tsx
import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import { asyncWithLDProvider } from 'launchdarkly-react-client-sdk';

(async () => {
  const LDProvider = await asyncWithLDProvider({
    clientSideID: '{{ environmentId }}',
    context: {
      kind: 'user',
      key: 'example-user-key',
      name: 'Sandy',
    },
  });

  const root = ReactDOM.createRoot(document.getElementById('root') as HTMLElement);
  root.render(
    <React.StrictMode>
      <LDProvider>
        <App />
      </LDProvider>
    </React.StrictMode>,
  );
})();
```
