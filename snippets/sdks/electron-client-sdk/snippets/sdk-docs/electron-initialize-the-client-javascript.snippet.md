---
id: electron-client-sdk/sdk-docs/electron-initialize-the-client-javascript
sdk: electron-client-sdk
kind: reference
lang: javascript
description: "JavaScript in section \"Initialize the client\""
---

```js
const LDElectron = require('launchdarkly-electron-client-sdk')

// You'll need this user later, but you can ignore it for now.
const user = { key: 'example' }
const options = {}
const client = LDElectron.initializeInMain('example-client-side-id', user, options)
```
