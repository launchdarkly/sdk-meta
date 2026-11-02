---
id: node-server-sdk/sdk-docs/features/overrides/file-overrides
sdk: node-server-sdk
kind: reference
lang: javascript
description: File-based flag override source configuration example for Node.js (server-side) SDK v9.15+.
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only
---

```js
import { init } from '@launchdarkly/node-server-sdk';

const client = init('YOUR_SDK_KEY', {
  dataSystem: {
    overrides: {
      type: 'file',
      paths: ['/etc/launchdarkly/overrides.json'],
    },
  },
});
```
