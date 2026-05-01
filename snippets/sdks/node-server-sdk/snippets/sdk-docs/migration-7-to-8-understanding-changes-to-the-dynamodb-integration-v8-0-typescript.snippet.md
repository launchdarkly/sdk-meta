---
id: node-server-sdk/sdk-docs/migration-7-to-8-understanding-changes-to-the-dynamodb-integration-v8-0-typescript
sdk: node-server-sdk
kind: reference
lang: typescript
description: "v8.0+ (TypeScript) in section \"Understanding changes to the DynamoDB integration\""
---

```ts
import * as ld from '@launchdarkly/node-server-sdk';
import { DynamoDBFeatureStore } from '@launchdarkly/node-server-sdk-dynamodb';

const store = DynamoDBFeatureStore(
  'your-table',
  { cacheTTL: 30 }
);

const options: ld.LDOptions = {
  featureStore: store,
};
const client = ld.init('YOUR_SDK_KEY', options);
```
