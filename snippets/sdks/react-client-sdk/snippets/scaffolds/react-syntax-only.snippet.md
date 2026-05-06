---
id: react-client-sdk/scaffolds/react-syntax-only
sdk: react-client-sdk
kind: scaffold
lang: tsx
file: src/Snippet.tsx
description: |
  Parse-only validator for React client SDK doc fragments.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: react-client
  entrypoint: src/Snippet.tsx
---

```tsx
//IMPORT_LIFT_TARGET

// @ts-expect-error -- never invoked; references to undefined names are fine
async function _wrappee() {
  if (false) {
//BODY_BEGIN
{{ body }}
//BODY_END
  }
}

document.body.textContent = 'feature flag evaluates to true';
```
