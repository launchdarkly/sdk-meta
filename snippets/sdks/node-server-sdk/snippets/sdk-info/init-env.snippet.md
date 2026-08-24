---
id: node-server-sdk/sdk-info/init-env
sdk: node-server-sdk
kind: init
lang: javascript
file: node-server-sdk/init-env.txt
description: Client initialization snippet for node-server-sdk reading the SDK key from an environment variable.
validation:
  scaffold: node-server-sdk/scaffolds/init-runner
---

```javascript
import * as LaunchDarkly from '@launchdarkly/node-server-sdk';

// Reads your LaunchDarkly SDK key from the LAUNCHDARKLY_SDK_KEY
// environment variable, so one build can run in every environment.
const client = LaunchDarkly.init(process.env.LAUNCHDARKLY_SDK_KEY);

client.once('ready', function () {
  // For onboarding purposes only we flush events as soon as
  // possible so we quickly detect your connection.
  // You don't have to do this in practice because events are automatically flushed.
  client.flush();
  console.log('SDK successfully initialized!');
});
```
