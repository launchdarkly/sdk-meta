---
id: node-server-sdk/scaffolds/openfeature-node-runner
sdk: node-server-sdk
kind: scaffold
lang: javascript
file: index.mjs
description: |
  Runs an OpenFeature provider doc fragment that assumes a registered
  provider and a bound `client`, `provider`, and `context` already
  exist — the "construct a context", "evaluate a context", "track
  metrics", and "access the LaunchDarkly client" fragments. The
  scaffold registers a real LaunchDarkly provider, binds those names,
  then runs the fragment inside its own block so a fragment that
  re-declares `context` (the construct-a-context examples) shadows the
  outer binding rather than colliding with it. After the fragment runs,
  the scaffold evaluates a flag with the outer context and prints the
  success line the harness greps for. Requires LaunchDarkly credentials
  because the provider connects.
inputs:
  body:
    type: string
    description: The wrappee fragment, run with `provider`, `client`, and `context` in scope.
validation:
  runtime: node
  entrypoint: index.mjs
  requirements: |
    @openfeature/server-sdk
    @launchdarkly/node-server-sdk
    @launchdarkly/openfeature-node-server
---

```javascript
import { OpenFeature } from '@openfeature/server-sdk';
import { LaunchDarklyProvider } from '@launchdarkly/openfeature-node-server';

const provider = new LaunchDarklyProvider(process.env.LAUNCHDARKLY_SDK_KEY);
await OpenFeature.setProviderAndWait(provider);
const client = OpenFeature.getClient();
const context = { targetingKey: 'example-user-key' };

{
{{ body }}
}

const _value = await client.getBooleanValue(process.env.LAUNCHDARKLY_FLAG_KEY, false, context);

console.log('feature flag evaluates to true');
```
