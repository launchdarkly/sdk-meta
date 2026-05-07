---
id: vue-client-sdk/scaffolds/vue-syntax-only-app
sdk: vue-client-sdk
kind: scaffold
lang: vue
file: src/App.vue
description: |
  Companion App.vue for `vue-syntax-only`. Provides a stable Vue
  component the JS-body main.js can import (via the canonical
  `import App from './App.vue'`) without supplying its own. The
  component renders the EXAM-HELLO success line, but the syntax-only
  scaffold's main.js sets `document.body.textContent` directly, so
  this companion's render path is never the one the harness asserts
  against — it exists purely to satisfy module resolution.
---

```vue
<template>
  <div>feature flag evaluates to true</div>
</template>
```
