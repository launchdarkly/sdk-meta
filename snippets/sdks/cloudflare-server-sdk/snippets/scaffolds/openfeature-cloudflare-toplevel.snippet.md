---
id: cloudflare-server-sdk/scaffolds/openfeature-cloudflare-toplevel
sdk: cloudflare-server-sdk
kind: scaffold
lang: typescript
file: snippet.ts
description: |
  Type-checks the Cloudflare OpenFeature import fragment against the real
  `@openfeature/server-sdk` and `@launchdarkly/openfeature-cloudflare-server`
  packages via the edge-tsc validator. A typo in a package or export name
  fails module resolution.
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
