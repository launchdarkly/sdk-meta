---
id: js-client-sdk/scaffolds/js-syntax-only
sdk: js-client-sdk
kind: scaffold
lang: javascript
file: src/app.ts
description: |
  Parse-only validator for JavaScript client SDK doc fragments. Stages
  itself as `src/app.ts` so the js-client validator's pre-baked tsdown
  project picks it up; the bundle is loaded into headless Chromium and
  the page is asserted to print the EXAM-HELLO success line.

  The wrappee body is wrapped inside a never-invoked function — its
  references to `ldclient`, hooks like `useFlags`, etc. don't have to
  resolve, but the file must be syntactically valid TypeScript-flavored
  JS.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: js-client
  entrypoint: src/app.ts
---

```javascript
// Wrap in an async IIFE so the wrappee body can use top-level `await`
// (e.g. `await client.waitForInitialization(...)`). tsdown's parser
// rejects bare top-level `await` outside a module-scope `async` IIFE.
// The body never executes — the IIFE's `if (false)` guard means the
// EXAM-HELLO line is the only side effect.
(async function _wrappee() {
  if (false) {
{{ body }}
  }
})();

document.body.textContent = 'feature flag evaluates to true';
```
