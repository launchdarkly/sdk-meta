---
id: vue-client-sdk/scaffolds/init-runner-observability
sdk: vue-client-sdk
kind: scaffold
lang: javascript
file: src/main.js
description: |
  End-to-end runner for the `observability/initialize` snippet body.

  The wrappee body is a complete entrypoint: it imports
  `createApp`, `App`, `LDPlugin`, `Observability`, `SessionReplay`,
  etc. (the symbols come from the matching `observability/import`
  snippet — but Vue's main.js style puts the imports at module scope
  alongside the runtime calls), calls `createApp(App)`, registers
  `LDPlugin` with the client-side ID + plugins list, and mounts to
  `#app`. The companion App.vue renders a known sentinel that the
  poll loop watches for.

  We don't assert observability data flows back to LaunchDarkly —
  just that the SDK starts cleanly with the o11y plugin attached.

  The wrappee's `'SDK_KEY'` literal is substituted with the live
  `LAUNCHDARKLY_CLIENT_SIDE_ID` env var via the snippet's
  `validation.placeholders` map.
inputs:
  body:
    type: string
    description: The wrappee init snippet body, embedded after key substitution.
validation:
  runtime: vue-client
  entrypoint: src/main.js
  companions:
    - vue-client-sdk/scaffolds/init-runner-app
---

```javascript
import { createApp } from 'vue';
import App from './App.vue';
import { LDPlugin } from 'launchdarkly-vue-client-sdk';
import Observability, { LDObserve } from '@launchdarkly/observability';
import SessionReplay, { LDRecord } from '@launchdarkly/session-replay';

{{ body }}

// Leading `;` so ASI doesn't try to call the result of the body's
// final `app.mount('#app')` (the body has no trailing semicolon).
;(function pollForBodySentinel() {
  const sentinel = 'vue-init-runner-ok';
  const deadline = Date.now() + 25_000;
  const tick = () => {
    const text = document.body.textContent || '';
    if (text.includes(sentinel)) {
      document.body.innerHTML = '<p>feature flag evaluates to true</p>';
      return;
    }
    if (Date.now() < deadline) {
      setTimeout(tick, 200);
      return;
    }
    document.body.innerHTML = '<p>scaffold: sentinel never rendered (LDPlugin did not finish initializing)</p>';
  };
  tick();
})();
```
