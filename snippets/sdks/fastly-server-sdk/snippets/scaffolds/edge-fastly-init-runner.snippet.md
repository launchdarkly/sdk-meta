---
id: fastly-server-sdk/scaffolds/edge-fastly-init-runner
sdk: fastly-server-sdk
kind: scaffold
lang: typescript
file: snippet.ts
description: |
  Type-checks a Fastly "initialize the client" fragment. Supplies the
  `KVStore` and `init` imports the page's import block documents, then
  splices the fragment, which builds the KV store and calls init(...)
  with the events-backend option. tsc checks the init call against the
  real signature.
inputs:
  body:
    type: string
    description: The wrappee init fragment; builds store and binds ldClient.
validation:
  runtime: edge-tsc
  entrypoint: snippet.ts
---

```typescript
/// <reference types="@fastly/js-compute" />
import { KVStore } from 'fastly:kv-store';
import { init } from '@launchdarkly/fastly-server-sdk';

{{ body }}
```
