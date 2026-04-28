---
id: react-client-sdk/getting-started/app-tsx
sdk: react-client-sdk
kind: hello-world
lang: tsx
file: src/App.tsx
description: App component that uses useFlags to render the flag value.
inputs:
  featureKey:
    type: flag-key
    description: Default flag key (camelCased) baked into the rendered source.
ld-application:
  slot: app-tsx
---

Use the `useFlags` hook to evaluate flags. For example, in `App.tsx`:

```tsx
import { useFlags } from 'launchdarkly-react-client-sdk';

function App() {
  const { {{ featureKey }} } = useFlags();

  return (
    <div style={{ backgroundColor: {{ featureKey }} ? 'green' : 'red' }}>
      The {{ featureKey }} feature flag evaluates to <b>{ {{ featureKey }} ? 'true' : 'false'}</b>
    </div>
  );
}

export default App;
```
