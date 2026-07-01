---
id: fastly-server-sdk/scaffolds/edge-fastly-toplevel
sdk: fastly-server-sdk
kind: scaffold
lang: typescript
file: snippet.ts
description: |
  Type-checks a Fastly edge SDK import fragment against the real
  `@launchdarkly/fastly-server-sdk` package and the Fastly Compute
  `fastly:kv-store` module (from @fastly/js-compute) via the edge-tsc
  validator.
inputs:
  body:
    type: string
    description: The wrappee's import statements, type-checked as a module.
validation:
  runtime: edge-tsc
  entrypoint: snippet.ts
---

```typescript
/// <reference types="@fastly/js-compute" />
{{ body }}
```
