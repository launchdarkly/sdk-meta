---
id: node-server-sdk/sdk-docs/migration-7-to-8-understanding-changes-to-the-testdata-integration-v7-x-javascript
sdk: node-server-sdk
kind: reference
lang: javascript
description: "v7.x (JavaScript) in section \"Understanding changes to the TestData integration\""
---

```js
const { TestData } = require('launchdarkly-node-server-sdk/integrations');

const td = TestData();
testData.update(td.flag('example-flag-key').booleanFlag().variationForAllUsers(true));
const client = new LDClient('YOUR_SDK_KEY', { updateProcessor: td });

// flags can be updated at any time:
td.update(td.flag('flag-key-456def')
    .variationForUser('example-user-key', true)
    .fallthroughVariation(false));

```
