---
id: node-server-sdk/sdk-docs/features/datasaving/standard-setup
sdk: node-server-sdk
kind: reference
lang: javascript
description: Data saving mode standard setup for Node.js (server-side).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only
---

```js
const ldClient = LaunchDarkly.init('YOUR_SDK_KEY', {
  dataSystem: {
    dataSource: {
        dataSourceOptionsType: 'standard',

        // if you use the stream, streamInitialReconnectDelay, or pollInterval options,
        // these options are now part of the dataSystem options,
        // and are set within the dataSource option
    }
    // if you use the persistentStore or useLDD option,
    // these options are now part of the dataSystem option
  }
});
```
