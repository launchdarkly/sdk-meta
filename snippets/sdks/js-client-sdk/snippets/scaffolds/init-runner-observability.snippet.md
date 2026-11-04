---
id: js-client-sdk/scaffolds/init-runner-observability
sdk: js-client-sdk
kind: scaffold
lang: javascript
file: src/app.ts
description: |
  End-to-end runner for the `observability/initialize` snippet body.

  The wrappee body assumes `createClient`, `Observability`, `SessionReplay`,
  `LDObserve`, and `LDRecord` are in scope (the symbols come from the
  matching `observability/import` snippet). This scaffold supplies
  those imports at module scope, splices the body inside an async
  IIFE, and awaits the resulting `client.waitForInitialization()`,
  failing unless it resolves with `status: 'complete'`.
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
import { createClient } from '@launchdarkly/js-client-sdk';
import Observability, { LDObserve } from '@launchdarkly/observability';
import SessionReplay, { LDRecord } from '@launchdarkly/session-replay';

(async () => {
  // The wrappee body declares
  //   const context = { kind: 'user', key: '...' };
  //   const client = createClient('SDK_KEY', context, { plugins: [...] });
  //   client.start();
  // Splicing it here at function scope binds `client` for the
  // initialization await below.
  {{ body }}

  const result = await client.waitForInitialization({ timeout: 10 });
  if (result.status === 'complete') {
    document.body.textContent = 'feature flag evaluates to true';
  } else {
    document.body.textContent = 'scaffold: waitForInitialization status: ' + result.status;
  }
})();
```
