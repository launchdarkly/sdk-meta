---
id: js-client-sdk/sdk-docs/install-the-sdk-javascript-and-typescript-js-sdk-v4-x
sdk: js-client-sdk
kind: reference
lang: javascript
description: "JavaScript and TypeScript, JS SDK v4.x in section \"Install the SDK\""
# TODO(snippet-bug): same shape as the v3.7 install snippet — three
# parallel `LDClient` import styles in one fence (CommonJS / ESM / TS)
# that redeclare the binding when parsed as one file. Needs splitting
# into three snippets in a follow-up snippet-bugs PR.
---

```js
// Using CommonJS (for Node.js environments without ES modules)
const LDClient = require('@launchdarkly/js-client-sdk');

// Using ES2015 modules (modern JavaScript environments and most bundlers)
import LDClient from '@launchdarkly/js-client-sdk';

// Using TypeScript (same as ES modules, with type support)
import LDClient from '@launchdarkly/js-client-sdk';
```
