---
id: react-client-sdk/sdk-docs/initialize-the-client-react-web-sdk-v4-0
sdk: react-client-sdk
kind: reference
lang: javascript
description: "React Web SDK v4.0 in section \"Initialize the client\""
validation:
  scaffold: react-client-sdk/scaffolds/react-syntax-only
---

```js
import { useInitializationStatus, createLDReactProvider } from '@launchdarkly/react-sdk';

const LDProvider = createLDReactProvider('your-client-side-id', { kind: 'user', key: 'user-key' });

function App() {
  const { status, error } = useInitializationStatus();
  if (status === 'initializing') return <div> initializing </div>;
  if (status === 'failed') return <div> Error: {error?.message} </div>;
  return <div> Your application code </div>;
}

root.render(<LDProvider><App /></LDProvider>);
```
