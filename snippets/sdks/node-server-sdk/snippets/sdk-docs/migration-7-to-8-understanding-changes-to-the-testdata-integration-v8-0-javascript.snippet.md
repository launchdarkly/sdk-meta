---
id: node-server-sdk/sdk-docs/migration-7-to-8-understanding-changes-to-the-testdata-integration-v8-0-javascript
sdk: node-server-sdk
kind: reference
lang: javascript
description: "v8.0+ (JavaScript) in section \"Understanding changes to the TestData integration\""
---

```js
import { init } from '@launchdarkly/node-server-sdk';
import { TestData } from '@launchdarkly/node-server-sdk/integrations';

const td = new TestData();
testData.update(td.flag('example-flag-key').booleanFlag().variationForAll(true));
const client = init('YOUR_SDK_KEY', { updateProcessor: td.getFactory() });

// flags can be updated at any time:
td.update(td.flag('flag-key-456def')
    .variationForUser('example-user-key', true)
    .fallthroughVariation(false));

```
