---
id: node-server-sdk/sdk-docs/features/storing-data/dynamodb/dynamodb-ts-v7
sdk: node-server-sdk
kind: reference
lang: typescript
description: DynamoDB feature store configuration example for Node.js (server-side) SDK v7.x and earlier (TypeScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```ts
import * as ld from 'launchdarkly-node-server-sdk';
import { DynamoDBFeatureStore } from 'launchdarkly-node-server-sdk-dynamodb';

const store = DynamoDBFeatureStore(
  'my-table',
  { cacheTTL: 30 }
);

const options: ld.LDOptions = {
  featureStore: store,
};
const client = ld.init('YOUR_SDK_KEY', options);
```
