---
id: node-server-sdk/sdk-docs/migration-7-to-8-understanding-changes-to-the-filedata-integration-v8-0-javascript
sdk: node-server-sdk
kind: reference
lang: javascript
description: "v8.0+ (JavaScript) in section \"Understanding changes to the FileData integration\""
---

```js
const ld = require('@launchdarkly/node-server-sdk');
const { FileDataSourceFactory } = require('@launchdarkly/node-server-sdk/integrations');

const fileData = new FileDataSourceFactory({
  paths: [ 'file1.json', 'file2.json' ]
});

const options = {
  updateProcessor: fileData.getFactory()
};

const client = ld.init('YOUR_SDK_KEY', options);
```
