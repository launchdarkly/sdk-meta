---
id: node-server-sdk/sdk-docs/migration-7-to-8-understanding-changes-to-the-dynamodb-integration-v8-0-javascript
sdk: node-server-sdk
kind: reference
lang: javascript
description: "v8.0+ (JavaScript) in section \"Understanding changes to the DynamoDB integration\""
---

```js
const ld = require('@launchdarkly/node-server-sdk');
const { DynamoDBFeatureStore } = require('@launchdarkly/node-server-sdk-dynamodb');

const store = DynamoDBFeatureStore(
  'your-table',
  { cacheTTL: 30 }
);

const options = {
  featureStore: store
};
const client = ld.init('YOUR_SDK_KEY', options);
```
