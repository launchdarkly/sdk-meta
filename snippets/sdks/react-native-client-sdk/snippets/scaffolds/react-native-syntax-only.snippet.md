---
id: react-native-client-sdk/scaffolds/react-native-syntax-only
sdk: react-native-client-sdk
kind: scaffold
lang: tsx
file: App.tsx
description: |
  Parse-only validator for React Native client SDK doc fragments.

  The harness reads `SNIPPET_MODE=syntax-only` from validation.env to
  dispatch into a Babel-only parse pass (no jest run, no LD
  initialization). Bodies with top-level `import …;` directives are
  handled via the `//IMPORT_LIFT_TARGET` / `//BODY_BEGIN` /
  `//BODY_END` marker pair: the harness's awk pre-step lifts any
  `import` lines from inside the body block up to module scope (ESM
  forbids imports inside a function body). Mirrors the react-client
  harness's IMPORT_LIFT pattern.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: react-native-client
  entrypoint: App.tsx
  env:
    SNIPPET_MODE: syntax-only
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

export default function App() {
  return null;
}
```
