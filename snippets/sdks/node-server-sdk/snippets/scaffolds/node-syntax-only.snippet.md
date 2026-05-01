---
id: node-server-sdk/scaffolds/node-syntax-only
sdk: node-server-sdk
kind: scaffold
lang: javascript
file: index.js
description: |
  Parse-only validator for Node server SDK doc fragments.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: node
  entrypoint: index.js
---

```javascript
(async () => {
  // Wrap the wrappee in a function that's never invoked. The presence of
  // unresolved symbols won't trip a syntax error; only malformed JS will.
  function _wrappee() {
{{ body }}
  }
  console.log('feature flag evaluates to true');
})();
```
