---
id: vue-client-sdk/scaffolds/vue-syntax-only
sdk: vue-client-sdk
kind: scaffold
lang: javascript
file: src/snippet.js
description: |
  Parse-only validator for Vue client SDK doc fragments.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: vue-client
  entrypoint: src/snippet.js
---

```javascript
(function _wrappee() {
{{ body }}
})();

console.log('feature flag evaluates to true');
```
