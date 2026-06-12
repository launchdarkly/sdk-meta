---
id: node-server-sdk/sdk-docs/features/datasaving/file-bootstrap
sdk: node-server-sdk
kind: reference
lang: javascript
description: Data saving mode with file-based bootstrap and live updates for Node.js (server-side).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only
---

```js
const ldClient = LaunchDarkly.init('YOUR_SDK_KEY', {
  dataSystem: {
    dataSource: {
      dataSourceOptionsType: 'custom',
      initializers: [
        { type: 'file', paths: ['flags.json'] },
        { type: 'polling' },
      ],
      synchronizers: [
        { type: 'streaming' },
        { type: 'polling' },
      ],
    }
  }
});
```
