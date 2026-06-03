---
id: js-client-sdk/sdk-docs/install-the-sdk-javascript-and-typescript-js-sdk-v3-7
sdk: js-client-sdk
kind: reference
lang: javascript
description: "JavaScript and TypeScript, JS SDK v3.7+ in section \"Install the SDK\""
# TODO(snippet-bug): body shows three alternative `LDClient` import
# styles (CommonJS require, ES module default import, TS import) in
# one fence as parallel examples — they redeclare `LDClient` when the
# validator parses the whole body. Fix in a follow-up snippet-bugs
# PR by splitting this snippet into three: install-the-sdk-commonjs,
# install-the-sdk-esm, install-the-sdk-typescript — each with one
# import style. Markers in the docs MDX get re-pointed at the split.
---

```js
// Using CommonJS (for Node.js environments without ES modules)
const LDClient = require('launchdarkly-js-client-sdk');

// Using ES2015 modules (modern JavaScript environments and most bundlers)
import LDClient from 'launchdarkly-js-client-sdk';

// Using TypeScript (same as ES modules, with type support)
import LDClient from 'launchdarkly-js-client-sdk';
```
