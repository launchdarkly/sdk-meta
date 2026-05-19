---
id: node-server-sdk/scaffolds/init-runner-observability
sdk: node-server-sdk
kind: scaffold
lang: javascript
file: index.mjs
description: |
  End-to-end runner for the `observability/initialize` snippet body.

  The wrappee body assumes `init`, `Observability`, and `LDObserve`
  are in scope (the symbols come from the matching
  `observability/import` snippet). This scaffold supplies those
  imports at module scope, splices the body, and awaits the
  resulting `ldClient.waitForInitialization()`. We don't assert that
  observability data actually flows back to LaunchDarkly — just that
  the SDK starts cleanly with the o11y plugin attached. A clean start
  emits the EXAM-HELLO line.

  The wrappee's `'SDK_KEY'` literal is substituted with the live
  `LAUNCHDARKLY_SDK_KEY` env var via the snippet's
  `validation.placeholders` map (handled by the dispatcher upstream).
inputs:
  body:
    type: string
    description: The wrappee init snippet body, embedded after key substitution.
validation:
  runtime: node
  entrypoint: index.mjs
  requirements: |
    @launchdarkly/node-server-sdk
    @launchdarkly/observability-node
---

```javascript
import { init } from '@launchdarkly/node-server-sdk';
import { Observability, LDObserve } from '@launchdarkly/observability-node';

// The wrappee body declares `const ldClient = init('SDK_KEY', { plugins: [...] });`.
// Splicing it here at module scope binds `ldClient` for the
// initialization wait below.
{{ body }}

try {
  await ldClient.waitForInitialization({ timeout: 10 });
  console.log('SDK successfully initialized!');
} catch (e) {
  console.error('scaffold: waitForInitialization rejected:', e.message);
  process.exit(1);
}

// Best-effort flush so events the body might have queued get sent.
try { await ldClient.flush(); } catch (_) {}
try { await ldClient.close(); } catch (_) {}

console.log('feature flag evaluates to true');
```
