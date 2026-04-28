---
id: react-native-client-sdk/getting-started/app-tsx
sdk: react-native-client-sdk
kind: hello-world
lang: tsx
file: App.tsx
description: Root component that wires the LDProvider with the React Native client.
inputs:
  mobileKey:
    type: mobile-key
    description: Mobile key baked into the rendered source.
ld-application:
  slot: app-tsx
validation:
  runtime: react-native-client
  entrypoint: App.tsx
  companions: [react-native-client-sdk/getting-started/welcome-tsx]
---

In `App.tsx`:

```tsx
import {
  AutoEnvAttributes,
  LDProvider,
  ReactNativeLDClient,
} from '@launchdarkly/react-native-client-sdk';

import Welcome from './src/welcome';

const featureClient = new ReactNativeLDClient(
  '{{ mobileKey }}',
  AutoEnvAttributes.Enabled,
  {
    debug: true,
    applicationInfo: {
      id: 'ld-rn-test-app',
      version: '0.0.1',
    },
  },
);

const App = () => {
  return (
    <LDProvider client={featureClient}>
      <Welcome />
    </LDProvider>
  );
};

export default App;
```
