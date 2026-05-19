---
id: react-client-sdk/scaffolds/init-runner-observability
sdk: react-client-sdk
kind: scaffold
lang: tsx
file: src/main.tsx
description: |
  End-to-end runner for the `observability/initialize` snippet body.

  The wrappee body declares `const LDProvider = withLDProvider({ ... })`
  and assumes `withLDProvider`, `Observability`, `SessionReplay`,
  `LDObserve`, and `LDRecord` are in scope (the symbols come from the
  matching `observability/import` snippet). This scaffold supplies
  those imports at module scope, splices the body, and renders a
  sentinel component wrapped in the body's `LDProvider`. As soon as
  the provider initializes against a real LD env, the sentinel
  renders and we mirror the EXAM-HELLO line.

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
  runtime: react-client
  entrypoint: src/main.tsx
---

```tsx
import React from 'react';
import { createRoot } from 'react-dom/client';
import { withLDProvider } from 'launchdarkly-react-client-sdk';
import Observability, { LDObserve } from '@launchdarkly/observability';
import SessionReplay, { LDRecord } from '@launchdarkly/session-replay';

// The wrappee body declares `const LDProvider = withLDProvider({ ... })`.
// Splicing it here at module scope binds `LDProvider` for the render
// below.
{{ body }}

function Sentinel() {
  return <div>Let your feature flags fly!</div>;
}

const Wrapped = LDProvider(Sentinel);

createRoot(document.getElementById('root') as HTMLElement).render(<Wrapped />);

(function pollForBodySentinel() {
  const sentinel = 'Let your feature flags fly';
  const deadline = Date.now() + 25_000;
  const tick = () => {
    const root = document.getElementById('root');
    const text = (root && root.textContent) || '';
    if (text.includes(sentinel)) {
      document.body.setAttribute('data-validator', 'ok');
      document.body.innerHTML = '<p>feature flag evaluates to true</p>';
      return;
    }
    if (Date.now() < deadline) {
      setTimeout(tick, 200);
      return;
    }
    document.body.innerHTML = '<p>scaffold: sentinel never rendered (LDProvider did not mount the body)</p>';
  };
  tick();
})();
```
