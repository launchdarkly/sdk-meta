---
id: vue-client-sdk/scaffolds/vue-syntax-only
sdk: vue-client-sdk
kind: scaffold
lang: javascript
file: src/main.js
description: |
  Parse-only validator for Vue client SDK doc fragments whose body is
  plain JavaScript (e.g. a `main.js` showing how to wire createApp +
  LDPlugin). The body is staged as the Vite project's `src/main.js`
  entrypoint, replacing the Dockerfile's placeholder; if it builds and
  runs without throwing during mount, the companion `App.vue` renders
  the EXAM-HELLO line.

  For Vue SFC bodies (`<script setup>...</script><template>...</template>`),
  use `vue-sfc-syntax-only` instead — those need to live at `src/App.vue`
  with the Vite vue plugin handling the SFC compilation.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, written to src/main.js.
validation:
  runtime: vue-client
  entrypoint: src/main.js
  companions:
    - vue-client-sdk/scaffolds/vue-syntax-only-app
---

```javascript
// Body is staged at module scope so its top-level
// `import { ... } from '...';` directives resolve via Vite's
// bundler. The body's runtime side-effects (createApp + mount)
// may throw against an unconfigured LDPlugin; that's fine — the
// EXAM-HELLO success line lives in the App.vue companion's
// `<template>` and is rendered from a separate `index.html` body
// element that the body never overwrites.

{{ body }}
```
