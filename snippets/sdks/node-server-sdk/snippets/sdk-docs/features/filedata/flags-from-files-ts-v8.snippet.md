---
id: node-server-sdk/sdk-docs/features/filedata/flags-from-files-ts-v8
sdk: node-server-sdk
kind: reference
lang: ts
description: File data source configuration example for Node.js (server-side) SDK v8.x (TypeScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```ts
import { LDOptions, init } from '@launchdarkly/node-server-sdk';
import { FileDataSourceFactory } from '@launchdarkly/node-server-sdk/integrations';

const fileData = new FileDataSourceFactory({
  paths: [ 'file1.json', 'file2.json' ],
});

const options: LDOptions = {
  updateProcessor: fileData.getFactory()
};

const client = init('YOUR_SDK_KEY', options);
```
