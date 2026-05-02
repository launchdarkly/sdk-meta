---
id: js-client-sdk/scaffolds/init-runner
sdk: js-client-sdk
kind: scaffold
lang: javascript
file: src/app.ts
description: |
  Runs an `init.txt`-style snippet end-to-end against a real LaunchDarkly
  env. The init body assumes `createClient` is in scope (gonfalon's docs
  put the import on the install snippet) — the scaffold prepends the
  v4 module import. The js-client validator builds this with tsdown and
  loads it in headless Chromium; on a clean wrappee run we set the
  page's body text to the EXAM-HELLO line so the harness's polling loop
  matches.

  The body's failure branch logs to `console.error`; we mirror that to
  `document.body.textContent` only on the success path so a failed init
  never produces a false positive.
inputs:
  body:
    type: string
    description: The wrappee init snippet body, embedded after key substitution.
validation:
  runtime: js-client
  entrypoint: src/app.ts
---

```javascript
import { createClient } from '@launchdarkly/js-client-sdk';

(async () => {
  let _scaffoldSucceeded = false;
  const _origLog = console.log.bind(console);
  console.log = (...args) => {
    if (args.join(' ').includes('SDK successfully initialized')) {
      _scaffoldSucceeded = true;
    }
    _origLog(...args);
  };

  {{ body }}

  if (_scaffoldSucceeded) {
    document.body.textContent = 'feature flag evaluates to true';
  } else {
    document.body.textContent = 'scaffold: wrappee did not print success';
  }
})();
```
