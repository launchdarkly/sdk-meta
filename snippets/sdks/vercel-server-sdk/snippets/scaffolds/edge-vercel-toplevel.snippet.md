---
id: vercel-server-sdk/scaffolds/edge-vercel-toplevel
sdk: vercel-server-sdk
kind: scaffold
lang: typescript
file: snippet.ts
description: |
  Type-checks a Vercel edge SDK import fragment against the real
  `@launchdarkly/vercel-server-sdk` and `@vercel/edge-config` packages
  via the edge-tsc validator.
inputs:
  body:
    type: string
    description: The wrappee's import statements, type-checked as a module.
validation:
  runtime: edge-tsc
  entrypoint: snippet.ts
---

```typescript
{{ body }}
```
