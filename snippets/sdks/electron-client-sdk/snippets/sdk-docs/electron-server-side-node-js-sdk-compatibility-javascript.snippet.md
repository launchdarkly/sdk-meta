---
id: electron-client-sdk/sdk-docs/electron-server-side-node-js-sdk-compatibility-javascript
sdk: electron-client-sdk
kind: reference
lang: javascript
description: "JavaScript in section \"Server-side Node.js SDK compatibility\""
---

```js
const realClient = LDElectron.initializeInMain('example-client-side-id', user, options)

const wrappedClient = LDElectron.createNodeSdkAdapter(realClient)

wrappedClient.waitForInitialization().then(function () {
  wrappedClient.variation(flagKey, user, defaultValue, function (err, result) {
    console.log('flag value is ' + result)
  })
})
```
