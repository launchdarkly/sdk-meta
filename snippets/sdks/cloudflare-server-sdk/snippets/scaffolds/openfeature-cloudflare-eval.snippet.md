---
id: cloudflare-server-sdk/scaffolds/openfeature-cloudflare-eval
sdk: cloudflare-server-sdk
kind: scaffold
lang: typescript
file: snippet.ts
description: |
  Type-checks a Cloudflare OpenFeature doc fragment that assumes a
  registered provider and bound `client`, `context`, and `ctx` -- the
  "construct a context", "evaluate a context", "track metrics", and
  "access the LaunchDarkly client" fragments. The scaffold constructs a
  real `LaunchDarklyProvider`, binds those names, then splices the
  fragment inside its own block so a fragment that re-declares `context`
  (the construct-a-context examples) shadows the outer binding rather
  than colliding with it. Calls like `client.getBooleanValue(...)` and
  `provider.getClient()` are checked against the real published types.
inputs:
  body:
    type: string
    description: The wrappee fragment; type-checked with `provider`, `client`, `context`, and `ctx` in scope.
validation:
  runtime: edge-tsc
  entrypoint: snippet.ts
---

```typescript
import { OpenFeature } from '@openfeature/server-sdk';
import { LaunchDarklyProvider } from '@launchdarkly/openfeature-cloudflare-server';

declare const env: { LD_KV: ConstructorParameters<typeof LaunchDarklyProvider>[1] };
declare const ctx: import('@cloudflare/workers-types').ExecutionContext;

const provider = new LaunchDarklyProvider('example-client-side-id', env.LD_KV);
await OpenFeature.setProviderAndWait(provider);
const client = OpenFeature.getClient();
const context = { targetingKey: 'example-user-key' };

{
{{ body }}
}
```
