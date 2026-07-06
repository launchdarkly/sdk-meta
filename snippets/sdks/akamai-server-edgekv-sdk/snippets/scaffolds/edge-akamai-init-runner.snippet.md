---
id: akamai-server-edgekv-sdk/scaffolds/edge-akamai-init-runner
sdk: akamai-server-edgekv-sdk
kind: scaffold
lang: typescript
file: snippet.ts
description: |
  Type-checks an Akamai "initialize the client" fragment against the real
  package. The scaffold supplies the `import { init }` the fragment
  assumes from the page's import block, then splices the fragment (which
  calls `init(...)` and binds `ldClient`). tsc checks the `init` config
  object against the real parameter type, so a misnamed or mistyped
  option fails.
inputs:
  body:
    type: string
    description: The wrappee init fragment; calls `init(...)` and binds `ldClient`.
validation:
  runtime: edge-tsc
  entrypoint: snippet.ts
---

```typescript
import { init } from '@launchdarkly/akamai-server-edgekv-sdk';

{{ body }}
```
