---
id: react-client-sdk/scaffolds/flag-eval-runner
sdk: react-client-sdk
kind: scaffold
lang: tsx
file: src/snippet-body.tsx
description: |
  Runs a `flagEval.txt`-style React snippet end-to-end against a real
  LaunchDarkly env. Unlike the init scaffold (whose wrappee body is a
  complete module-scope program), the flag-eval body is *not*
  standalone:

    - Its top-level `import { useBoolVariation } from '...'` is
      module-scope syntax that we want to keep at module scope.
    - Its `const flagValue = useBoolVariation(...)` and
      `if (...) { ... }` form a block that's only legal *inside a
      render context* (a React component called inside
      `<LDReactProvider>`).
    - The wrappee's flag-key string (e.g. `'feature-key'`) is a
      human-readable placeholder; the live flag key comes from
      `LAUNCHDARKLY_FLAG_KEY` and is substituted upstream via the
      wrappee's `validation.placeholders` block before the body
      reaches this scaffold.

  Splicing the body verbatim into a function body would yield a parse
  error on the `import`. Splicing at module scope would crash on the
  hooks call. So this scaffold stages the wrappee body verbatim at
  `src/snippet-body.tsx`, and the `react-client` validator harness
  (in flag-eval mode, selected via `validation.env: SNIPPET_MODE`)
  rewrites it: it lifts top-level `import` lines to module scope,
  wraps the remainder in a `WrappedFlagEvalBody` function component,
  and emits `src/main.tsx` + `src/App.tsx` boilerplate that mounts
  the component inside `<LDReactProvider>`. The `WrappedFlagEvalBody`
  unconditionally renders the EXAM-HELLO success line so the
  Playwright check passes once the body has mounted -- the assertion
  surface is "the body parsed, imports resolved, the hook ran inside
  a render context", not the flag's truth value.
inputs:
  body:
    type: string
    description: The wrappee flag-eval snippet body, staged verbatim before harness rewrite.
validation:
  runtime: react-client
  entrypoint: src/snippet-body.tsx
  env:
    SNIPPET_MODE: flag-eval
---

```tsx
{{ body }}
```
