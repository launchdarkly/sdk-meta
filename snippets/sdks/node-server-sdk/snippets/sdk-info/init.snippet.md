---
id: node-server-sdk/sdk-info/init
sdk: node-server-sdk
kind: init
lang: javascript
file: node-server-sdk/init.txt
description: Client initialization snippet for node-server-sdk.
---

```javascript
import * as LaunchDarkly from '@launchdarkly/node-server-sdk';

// This is your LaunchDarkly SDK key.
// Never hardcode your SDK key in production.
const client = LaunchDarkly.init('YOUR_SDK_KEY');

client.once('ready', function () {
  // For onboarding purposes only we flush events as soon as
  // possible so we quickly detect your connection.
  // You don't have to do this in practice because events are automatically flushed.
  client.flush();
  console.log('SDK successfully initialized!');
});
```
