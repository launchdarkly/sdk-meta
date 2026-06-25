---
id: node-server-sdk/sdk-docs/features/testdata/configure-v7x
sdk: node-server-sdk
kind: reference
lang: js
description: Test data source configuration for Node.js (server-side) SDK v7.x and earlier.
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only
---

```js
const ld = require('launchdarkly-node-server-sdk');
const { TestData } = require('launchdarkly-node-server-sdk/integrations');

const td = new TestData();
td.update(td.flag('example-flag-key').booleanFlag().variationForAllUsers(true));
const client = ld.init('YOUR_SDK_KEY', { updateProcessor: td });

// flags can be updated at any time:
td.update(td.flag('flag-key-456def')
    .variationForUser('example-user-key', true)
    .fallthroughVariation(false));

```
