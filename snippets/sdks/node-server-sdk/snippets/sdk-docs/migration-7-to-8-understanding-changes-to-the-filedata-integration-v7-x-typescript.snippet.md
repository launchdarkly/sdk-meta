---
id: node-server-sdk/sdk-docs/migration-7-to-8-understanding-changes-to-the-filedata-integration-v7-x-typescript
sdk: node-server-sdk
kind: reference
lang: typescript
description: "v7.x (TypeScript) in section \"Understanding changes to the FileData integration\""
---

```ts
import * as ld from 'launchdarkly-node-server-sdk';
const { FileDataSource } = require('launchdarkly-node-server-sdk/integrations');

const dataSource = FileDataSource({
  paths: [ 'file1.json', 'file2.json' ],
});

const options: ld.LDOptions = {
  updateProcessor: dataSource
};

const client = ld.init('YOUR_SDK_KEY', options);
```
