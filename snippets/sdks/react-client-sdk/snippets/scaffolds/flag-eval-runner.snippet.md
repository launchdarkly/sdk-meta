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

    - Its top-level `import { useFlags } from '...'` is module-scope
      syntax that we want to keep at module scope.
    - Its `const { ... } = useFlags();` and `if (...) { ... }` form a
      block that's only legal *inside a render context* (a React
      component called inside `<LDProvider>`).
    - Its destructured identifier (`featureKey`) is the human-readable
      placeholder; the React SDK exposes the live flag under the
      camelCased form of `LAUNCHDARKLY_FLAG_KEY` (e.g. `hello-boolean`
      becomes `helloBoolean`).

  Splicing the body verbatim into a function body would yield a parse
  error on the `import`. Splicing at module scope would crash on the
  hooks call. So this scaffold stages the wrappee body verbatim at
  `src/snippet-body.tsx`, and the `react-client` validator harness
  (in flag-eval mode, selected via `validation.env: SNIPPET_MODE`)
  rewrites it: it lifts top-level `import` lines to module scope,
  substitutes the body's `featureKey` token for the camelCased flag
  identifier, wraps the remainder in a `WrappedFlagEvalBody` function
  component, and emits `src/main.tsx` + `src/App.tsx` boilerplate that
  mounts the component inside `<LDProvider>`. The `WrappedFlagEvalBody`
  returns a sentinel string that the harness's Playwright check
  promotes into the EXAM-HELLO success line on a true evaluation.
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
