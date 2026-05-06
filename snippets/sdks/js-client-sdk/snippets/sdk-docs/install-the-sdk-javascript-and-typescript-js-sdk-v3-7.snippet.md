---
id: js-client-sdk/sdk-docs/install-the-sdk-javascript-and-typescript-js-sdk-v3-7
sdk: js-client-sdk
kind: reference
lang: javascript
description: "JavaScript and TypeScript, JS SDK v3.7+ in section \"Install the SDK\""
# Bucket C: body shows three alternative `LDClient` import styles
# (CommonJS require, ES module default import, TS import) as parallel
# examples — but co-existing in one file they redeclare LDClient.
# The docs intend "use one of these"; the validator can only see the
# whole body. See _sdk-docs-port-notes.md.
---

```js
// Using CommonJS (for Node.js environments without ES modules)
const LDClient = require('launchdarkly-js-client-sdk');

// Using ES2015 modules (modern JavaScript environments and most bundlers)
import LDClient from 'launchdarkly-js-client-sdk';

// Using TypeScript (same as ES modules, with type support)
import LDClient from 'launchdarkly-js-client-sdk';
```
