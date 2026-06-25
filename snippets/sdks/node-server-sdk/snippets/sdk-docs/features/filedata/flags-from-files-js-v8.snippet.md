---
id: node-server-sdk/sdk-docs/features/filedata/flags-from-files-js-v8
sdk: node-server-sdk
kind: reference
lang: js
description: File data source configuration example for Node.js (server-side) SDK v8.x (JavaScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```js
import { init } from '@launchdarkly/node-server-sdk';
import { FileDataSourceFactory } from '@launchdarkly/node-server-sdk/integrations';

const fileData = new FileDataSourceFactory({
  paths: [ 'file1.json', 'file2.json' ]
});

const options = {
  updateProcessor: fileData.getFactory()
};

const client = init('YOUR_SDK_KEY', options);
```
