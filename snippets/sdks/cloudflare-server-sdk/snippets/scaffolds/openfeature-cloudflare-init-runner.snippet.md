---
id: cloudflare-server-sdk/scaffolds/openfeature-cloudflare-init-runner
sdk: cloudflare-server-sdk
kind: scaffold
lang: typescript
file: snippet.ts
description: |
  Type-checks the Cloudflare OpenFeature "initialize the provider" fragment.
  Supplies the imports the docs page establishes earlier (the
  import-the-namespaces fragment) and a Workers `env` binding whose `LD_KV`
  is a real `KVNamespace`, so `new LaunchDarklyProvider(clientSideID,
  env.LD_KV, options)` and `OpenFeature.setProviderAndWait(provider)` are
  checked against the real signatures.
inputs:
  body:
    type: string
    description: The wrappee init fragment; constructs the provider with env.LD_KV and binds client.
validation:
  runtime: edge-tsc
  entrypoint: snippet.ts
---

```typescript
import { OpenFeature } from '@openfeature/server-sdk';
import { LaunchDarklyProvider } from '@launchdarkly/openfeature-cloudflare-server';

declare const env: { LD_KV: ConstructorParameters<typeof LaunchDarklyProvider>[1] };

{{ body }}
```
