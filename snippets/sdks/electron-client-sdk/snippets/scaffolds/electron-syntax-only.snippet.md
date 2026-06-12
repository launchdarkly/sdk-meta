---
id: electron-client-sdk/scaffolds/electron-syntax-only
sdk: electron-client-sdk
kind: scaffold
lang: javascript
file: src/app.ts
description: |
  Parse-only validator for Electron client SDK doc fragments. Uses the
  JS-client Docker validator since Electron is a JavaScript variant —
  the staged file path matches the js-client harness's hard-coded
  `src/app.ts` entrypoint so the wrappee body actually reaches tsdown
  + Chromium. The body sits in a never-invoked async IIFE so its
  references to Electron-only globals (`require`, `electron.app`, etc.)
  don't have to resolve at runtime; tsdown still parses the syntax.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, written to src/app.ts.
validation:
  runtime: js-client
  entrypoint: src/app.ts
---

```javascript
//IMPORT_LIFT_TARGET

// Wrap in an async IIFE so the wrappee body can use top-level `await`
// (e.g. `await client.waitForInitialization(...)`); the IIFE's
// `if (false)` guard means the body is never executed at runtime.
//
// Bodies with top-level `import ...;` directives are handled via the
// `//IMPORT_LIFT_TARGET` / `//BODY_BEGIN` / `//BODY_END` marker pair:
// the js-client harness's awk pre-step lifts any `import` lines from
// inside the body block up to module scope (ESM forbids imports
// inside a function body).
(async function _wrappee() {
  if (false) {
//BODY_BEGIN
{{ body }}
//BODY_END
  }
})();

document.body.textContent = 'feature flag evaluates to true';
```
