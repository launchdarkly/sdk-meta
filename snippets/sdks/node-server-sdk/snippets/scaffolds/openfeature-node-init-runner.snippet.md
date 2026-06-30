---
id: node-server-sdk/scaffolds/openfeature-node-init-runner
sdk: node-server-sdk
kind: scaffold
lang: javascript
file: index.mjs
description: |
  Runs an OpenFeature "initialize the provider" fragment end-to-end
  against a real LaunchDarkly environment. The fragment is expected to
  register a LaunchDarkly provider with OpenFeature and bind a client
  (`const client = OpenFeature.getClient()`); the scaffold supplies the
  `import` statements the fragment assumes are already present, then
  uses the bound `client` to evaluate a flag and print the success line
  the harness greps for. The fragment's `YOUR_SDK_KEY` literal is
  replaced with the real key via the snippet's `validation.placeholders`
  before this runs.
inputs:
  body:
    type: string
    description: The wrappee init fragment; registers the provider and binds `client`.
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

{{ body }}

const _value = await client.getBooleanValue(process.env.LAUNCHDARKLY_FLAG_KEY, false, {
  targetingKey: 'example-user-key',
});

console.log('feature flag evaluates to true');
```
