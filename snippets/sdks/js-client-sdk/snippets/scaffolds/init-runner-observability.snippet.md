---
id: js-client-sdk/scaffolds/init-runner-observability
sdk: js-client-sdk
kind: scaffold
lang: javascript
file: src/app.ts
description: |
  End-to-end runner for the `observability/initialize` snippet body.

  The wrappee body assumes `LDClient`, `Observability`, `SessionReplay`,
  `LDObserve`, and `LDRecord` are in scope (the symbols come from the
  matching `observability/import` snippet). This scaffold supplies
  those imports at module scope, splices the body inside an async
  IIFE, and awaits the resulting `client.waitForInitialization()`.
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
  runtime: js-client
  entrypoint: src/app.ts
---

```javascript
import LDClient from 'launchdarkly-js-client-sdk';
import Observability, { LDObserve } from '@launchdarkly/observability';
import SessionReplay, { LDRecord } from '@launchdarkly/session-replay';

(async () => {
  // The wrappee body declares
  //   const context = { kind: 'user', key: '...' };
  //   const client = LDClient.initialize('SDK_KEY', context, { plugins: [...] });
  // Splicing it here at function scope binds `client` for the
  // initialization await below.
  {{ body }}

  try {
    await client.waitForInitialization(10);
    document.body.textContent = 'feature flag evaluates to true';
  } catch (e) {
    document.body.textContent = 'scaffold: waitForInitialization rejected: ' + (e && e.message);
  }
})();
```
