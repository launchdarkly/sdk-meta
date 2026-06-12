---
id: node-server-sdk/sdk-docs/features/storing-data/dynamodb/dynamodb-js-v7
sdk: node-server-sdk
kind: reference
lang: javascript
description: DynamoDB feature store configuration example for Node.js (server-side) SDK v7.x and earlier (JavaScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```js
const ld = require('launchdarkly-node-server-sdk');
const { DynamoDBFeatureStore } = require('launchdarkly-node-server-sdk-dynamodb');

const store = DynamoDBFeatureStore('my-table', { cacheTTL: 30 });

const options = {
  featureStore: store
};
const client = ld.init('YOUR_SDK_KEY', options);
```
