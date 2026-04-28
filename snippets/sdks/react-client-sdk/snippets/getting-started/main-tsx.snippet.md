---
id: react-client-sdk/getting-started/main-tsx
sdk: react-client-sdk
kind: hello-world
lang: tsx
file: src/main.tsx
description: Vite-app entrypoint that wraps the React app with LDProvider.
inputs:
  environmentId:
    type: client-side-id
    description: Client-side ID baked into the rendered source.
ld-application:
  slot: main-tsx
---

In `main.tsx`, wrap your application with `LDProvider`:

```tsx
import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { LDProvider } from 'launchdarkly-react-client-sdk';
import App from './App';

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <LDProvider
      clientSideID="{{ environmentId }}"
      context={{ kind: 'user', key: 'example-user-key', name: 'Sandy' }}
    >
      <App />
    </LDProvider>
  </StrictMode>,
);
```
