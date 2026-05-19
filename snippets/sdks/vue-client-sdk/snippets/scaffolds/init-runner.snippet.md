---
id: vue-client-sdk/scaffolds/init-runner
sdk: vue-client-sdk
kind: scaffold
lang: javascript
file: src/main.js
description: |
  Runs an `init.txt`-style Vue snippet end-to-end against a real
  LaunchDarkly env. The init body is a complete entrypoint: it imports
  Vue + the LD plugin, calls `createApp(App)`, registers `LDPlugin`
  with a clientSideID, and mounts to `#app`. The scaffold drops the
  body verbatim into `src/main.js`, supplies the App.vue companion
  with a known sentinel, and appends a poll loop that promotes the
  sentinel into the page-level EXAM-HELLO success line once the
  LDPlugin has finished initializing and the App has mounted.

  The Vue plugin keeps its initialization transparent — by the time
  `app.mount('#app')` returns, `<App />` has already rendered. So
  seeing the sentinel at all confirms the createApp + use(LDPlugin)
  call didn't throw on the supplied clientSideID/context.
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
{{ body }}

// Leading `;` so ASI doesn't try to call the result of the body's
// final `app.mount('#app')` if the body lacks a trailing semicolon.
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
