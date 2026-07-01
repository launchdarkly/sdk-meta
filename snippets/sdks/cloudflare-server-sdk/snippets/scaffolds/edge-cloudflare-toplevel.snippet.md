---
id: cloudflare-server-sdk/scaffolds/edge-cloudflare-toplevel
sdk: cloudflare-server-sdk
kind: scaffold
lang: typescript
file: snippet.ts
description: |
  Type-checks a Cloudflare edge SDK import fragment against the real
  `@launchdarkly/cloudflare-server-sdk` package via the edge-tsc validator.
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
