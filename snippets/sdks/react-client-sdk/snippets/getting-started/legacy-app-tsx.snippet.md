---
id: react-client-sdk/getting-started/legacy-app-tsx
sdk: react-client-sdk
kind: hello-world
lang: tsx
file: src/App.tsx
description: CRA App.tsx (legacy variant) using useFlags to render the flag value.
inputs:
  featureKey:
    type: flag-key
    description: Default flag key (camelCased) baked into the rendered source. Note that gonfalon camel-cases the supplied flag key before substituting; for validation we use the env-var value as-is.
ld-application:
  slot: legacy-app-tsx
---

In `App.tsx`:

```tsx
import './App.css';
import { useFlags } from 'launchdarkly-react-client-sdk';

function App() {
  const { {{ featureKey }} } = useFlags();

  return (
      <div className="App">
        <header className="App-header" style={{backgroundColor: {{ featureKey }} ? '#00844B' : '#373841'}}>
            <p>The {{ featureKey }} feature flag evaluates to <b>{ {{ featureKey }} ? 'True' : 'False'}</b></p>
        </header>
      </div>
  );
}

export default App;
```
