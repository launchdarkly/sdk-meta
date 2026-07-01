---
id: akamai-server-edgekv-sdk/scaffolds/edge-akamai-toplevel
sdk: akamai-server-edgekv-sdk
kind: scaffold
lang: typescript
file: snippet.ts
description: |
  Type-checks an Akamai edge SDK import fragment against the real
  `@launchdarkly/akamai-server-edgekv-sdk` package. Routes through the
  `edge-tsc` validator (tsc --noEmit with module resolution), so a named
  import that the package doesn't export fails the check.
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
