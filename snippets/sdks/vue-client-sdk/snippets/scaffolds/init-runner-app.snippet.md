---
id: vue-client-sdk/scaffolds/init-runner-app
sdk: vue-client-sdk
kind: scaffold
lang: vue
file: src/App.vue
description: |
  Companion App.vue for the vue-client init scaffold. The init body's
  `import App from './App.vue';` resolves to this file. Displays a
  known sentinel inside `#app` so the scaffold's poll loop can detect
  that LDPlugin's app.mount finished without throwing.
---

```vue
<template>
  <div>vue-init-runner-ok</div>
</template>
```
