---
id: js-client-sdk/scaffolds/js-syntax-only
sdk: js-client-sdk
kind: scaffold
lang: javascript
file: index.js
description: |
  Parse-only validator for JavaScript client SDK doc fragments.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: js-client
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
