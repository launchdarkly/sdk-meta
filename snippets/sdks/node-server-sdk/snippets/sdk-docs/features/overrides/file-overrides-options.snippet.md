---
id: node-server-sdk/sdk-docs/features/overrides/file-overrides-options
sdk: node-server-sdk
kind: reference
lang: javascript
description: File-based flag override source with multiple files, duplicate key handling, and change detection for Node.js (server-side) SDK v9.15+.
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only
---

```js
const client = init('YOUR_SDK_KEY', {
  dataSystem: {
    overrides: {
      type: 'file',
      paths: ['/etc/launchdarkly/overrides.json', '/etc/launchdarkly/team-overrides.json'],
      duplicateKeysHandling: 'ignore',
      changeDetection: 'watching',
    },
  },
});
```
