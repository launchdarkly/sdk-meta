---
id: vue-client-sdk/scaffolds/vue-sfc-syntax-only
sdk: vue-client-sdk
kind: scaffold
lang: vue
file: src/App.vue
description: |
  Parse-only validator for Vue client SDK doc fragments whose body is
  a Vue Single-File Component (SFC) -- typically a `<script setup>`
  block followed by a `<template>` block. The body is staged as the
  Vite project's `src/App.vue`, where `@vitejs/plugin-vue` compiles
  the SFC and tsdown / vite catches syntax errors. The companion
  `main.js` mounts an empty replacement App so the Vue runtime
  doesn't try to evaluate the body's `useLDClient` etc. against a
  missing LDPlugin -- success is signalled by the companion writing
  the EXAM-HELLO line directly to the page.

  For pure-JavaScript bodies (e.g. a `main.js` example wiring
  `createApp(App).use(LDPlugin)`), use `vue-syntax-only` instead.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, written to src/App.vue.
validation:
  runtime: vue-client
  entrypoint: src/App.vue
  companions:
    - vue-client-sdk/scaffolds/vue-sfc-syntax-only-main
---

```vue
{{ body }}
```
