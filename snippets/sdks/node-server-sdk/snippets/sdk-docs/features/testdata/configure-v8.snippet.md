---
id: node-server-sdk/sdk-docs/features/testdata/configure-v8
sdk: node-server-sdk
kind: reference
lang: js
description: Test data source configuration for Node.js (server-side) SDK v8.x.
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only
---

```js
const ld = require('@launchdarkly/node-server-sdk');
const { TestData } = require('@launchdarkly/node-server-sdk/integrations');

const td = new TestData();
td.update(td.flag('example-flag-key').booleanFlag().variationForAll(true));
const client = ld.init('YOUR_SDK_KEY', { updateProcessor: td.getFactory() });

// flags can be updated at any time:
td.update(td.flag('flag-key-456def')
    .variationForUser('example-user-key', true)
    .fallthroughVariation(false));

```
