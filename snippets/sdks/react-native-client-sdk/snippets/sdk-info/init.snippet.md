---
id: react-native-client-sdk/sdk-info/init
sdk: react-native-client-sdk
kind: init
lang: tsx
file: react-native-client-sdk/init.txt
description: Client initialization snippet for react-native-client-sdk.
---

```tsx
import {
  AutoEnvAttributes,
  LDProvider,
  ReactNativeLDClient,
} from '@launchdarkly/react-native-client-sdk';

// This is your mobile key.
const ldClient = new ReactNativeLDClient('YOUR_MOBILE_KEY', AutoEnvAttributes.Enabled, {
  debug: true,
  applicationInfo: {
    id: 'ld-rn-test-app',
    version: '0.0.1',
  },
});

// A "context" is a data object representing users, devices, organizations, and other entities.
const context = { kind: 'user', key: 'EXAMPLE_CONTEXT_KEY' };

const App = () => {
  useEffect(() => {
    ldClient.identify(context);
  }, []);

  return (
    <LDProvider client={ldClient}>
      <YourComponent />
    </LDProvider>
  );
};

export default App;
```
