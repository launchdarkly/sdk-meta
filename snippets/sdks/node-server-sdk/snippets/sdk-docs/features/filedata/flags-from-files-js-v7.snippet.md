---
id: node-server-sdk/sdk-docs/features/filedata/flags-from-files-js-v7
sdk: node-server-sdk
kind: reference
lang: js
description: File data source configuration example for Node.js (server-side) SDK v7.x (JavaScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```js
const ld = require('launchdarkly-node-server-sdk');
const { FileDataSource } = require('launchdarkly-node-server-sdk/integrations');

const dataSource = FileDataSource({
  paths: [ 'file1.json', 'file2.json' ]
});

const options = {
  updateProcessor: dataSource
};

const client = ld.init('YOUR_SDK_KEY', options);
```
