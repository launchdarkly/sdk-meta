---
id: node-client-sdk/scaffolds/node-client-syntax-only
sdk: node-client-sdk
kind: scaffold
lang: javascript
file: index.js
description: |
  Parse-only validator for Node client SDK doc fragments.
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
  function _wrappee() {
{{ body }}
  }
  console.log('feature flag evaluates to true');
})();
```
