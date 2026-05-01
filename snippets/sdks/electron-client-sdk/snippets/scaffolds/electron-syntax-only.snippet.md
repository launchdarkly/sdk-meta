---
id: electron-client-sdk/scaffolds/electron-syntax-only
sdk: electron-client-sdk
kind: scaffold
lang: javascript
file: index.js
description: |
  Parse-only validator for Electron client SDK doc fragments. Uses the JS-client Docker validator since Electron is a JS variant.
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
