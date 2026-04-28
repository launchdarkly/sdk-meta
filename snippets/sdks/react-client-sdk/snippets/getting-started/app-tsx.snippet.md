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
validation:
  runtime: react-client
  entrypoint: src/App.tsx
  companions: [react-client-sdk/getting-started/main-tsx]
---

Use the `useFlags` hook to evaluate flags. For example, in `App.tsx`:

```tsx
import { useFlags } from 'launchdarkly-react-client-sdk';

function App() {
  const { {{ featureKey | camelCase }} } = useFlags();

  return (
    <div style={{ backgroundColor: {{ featureKey | camelCase }} ? 'green' : 'red' }}>
      The {{ featureKey | camelCase }} feature flag evaluates to <b>{ {{ featureKey | camelCase }} ? 'true' : 'false'}</b>
    </div>
  );
}

export default App;
```
