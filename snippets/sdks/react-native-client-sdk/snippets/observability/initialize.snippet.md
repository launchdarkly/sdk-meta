---
id: react-native-client-sdk/observability/initialize
sdk: react-native-client-sdk
kind: initialize
lang: javascript
file: react-native-client-sdk/observability/initialize.txt
description: Initialize react-native-client-sdk with observability plugin.
validation:
  scaffold: react-native-client-sdk/scaffolds/init-runner-observability
  placeholders:
    SDK_KEY: LAUNCHDARKLY_MOBILE_KEY
---

```javascript
const client = new ReactNativeLDClient(
    'SDK_KEY',
    // … your existing config, if relevant
    AutoEnvAttributes.Enabled,
    {
      plugins: [
        new Observability({
          serviceName: 'example-service',
          serviceVersion: 'example-sha'
        })
      ],
    }
);
```
