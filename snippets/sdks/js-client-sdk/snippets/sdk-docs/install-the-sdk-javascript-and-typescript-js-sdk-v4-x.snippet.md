---
id: js-client-sdk/sdk-docs/install-the-sdk-javascript-and-typescript-js-sdk-v4-x
sdk: js-client-sdk
kind: reference
lang: javascript
description: "JavaScript and TypeScript, JS SDK v4.x in section \"Install the SDK\""
# Bucket C: same shape as the v3.7 install snippet — three parallel
# import styles in one body that redeclare LDClient. See
# _sdk-docs-port-notes.md.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only
---

```js
// Using CommonJS (for Node.js environments without ES modules)
const LDClient = require('@launchdarkly/js-client-sdk');

// Using ES2015 modules (modern JavaScript environments and most bundlers)
import LDClient from '@launchdarkly/js-client-sdk';

// Using TypeScript (same as ES modules, with type support)
import LDClient from '@launchdarkly/js-client-sdk';
```
