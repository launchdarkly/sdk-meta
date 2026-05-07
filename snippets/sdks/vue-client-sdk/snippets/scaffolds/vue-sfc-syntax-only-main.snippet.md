---
id: vue-client-sdk/scaffolds/vue-sfc-syntax-only-main
sdk: vue-client-sdk
kind: scaffold
lang: javascript
file: src/main.js
description: |
  Companion main.js for the `vue-sfc-syntax-only` scaffold. Vite's
  default entry is `src/main.js`; if we leave the Dockerfile's
  placeholder `main.js` (which mounts the staged `App.vue`), the
  body's `useLDClient` / `useLDFlag` calls fire at mount time with no
  LDPlugin installed and the page renders an empty body instead of
  the EXAM-HELLO line. Replacing main.js with this stub bypasses the
  body-as-App render entirely and writes the success line directly to
  the page, so the harness's DOM check matches as long as the staged
  App.vue compiled cleanly.
---

```javascript
document.body.textContent = 'feature flag evaluates to true';
```
