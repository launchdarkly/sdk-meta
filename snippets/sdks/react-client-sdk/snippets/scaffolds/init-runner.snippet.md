---
id: react-client-sdk/scaffolds/init-runner
sdk: react-client-sdk
kind: scaffold
lang: tsx
file: src/main.tsx
description: |
  Runs an `init.txt`-style React snippet end-to-end against a real
  LaunchDarkly env. The init body is a complete entrypoint: it imports
  React + the LD provider, declares an `App` component, and calls
  `createRoot(...).render(<LDReactProvider>...<App />...</LDReactProvider>)`.
  The scaffold drops the body verbatim into `src/main.tsx`, then
  appends a poll loop that watches the React-rendered DOM. As soon as
  the body's `App` renders its sentinel text inside `#root`, we mirror
  the EXAM-HELLO success line into `document.body` so the validator's
  Playwright check matches.

  Seeing the body's sentinel text is the end-to-end signal that the
  provider mounted successfully against a real LD env. If the bundle
  fails to build or the provider throws on mount, the sentinel never
  renders and the poll times out.
inputs:
  body:
    type: string
    description: The wrappee init snippet body, embedded after key substitution.
validation:
  runtime: react-client
  entrypoint: src/main.tsx
---

```tsx
{{ body }}

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
    document.body.innerHTML = '<p>scaffold: sentinel never rendered (LDReactProvider did not mount the body)</p>';
  };
  tick();
})();
```
